package visitors

import (
	"github.com/genome-engine/genome/engine/types"
	"github.com/genome-engine/genome/engine/units"
	"go/ast"
	"go/token"
	"strings"
)

func (vis *GeneralVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	case *ast.File:
		var isMain bool
		packName := n.Name.Name
		packId := id(packName)

		if packName == "main" {
			packName = vis.packMainDir
			isMain = true
		}

		pack := units.NewPackage(packId, packName)
		pack.IsMain = isMain

		vis.pack = pack
		_ = vis.Collection.Add(vis.pack)

		if modeExist(vis.modes, ValueDecls) {
			for _, decl := range n.Decls {
				switch genDecl := decl.(type) {
				case *ast.GenDecl:
					switch genDecl.Tok {
					case token.VAR:
						vis.varsHandle(genDecl)
					case token.CONST:
						vis.constHandle(genDecl)
					}
				}
			}
		}
	case *ast.ImportSpec:
		if modeExist(vis.modes, Imports) {
			importVis := NewImportVisitor(vis.src, vis.Collection)
			importVis.pack = vis.pack
			ast.Walk(importVis, n)

			vis.importsPaths = append(vis.importsPaths, importVis.importsPaths...)
		}
		return vis
	case *ast.TypeSpec:
		switch typeSpec := n.Type.(type) {
		case *ast.StructType:
			if modeExist(vis.modes, Structs) {
				structVis := NewStructVisitor(vis.src, vis.Collection)
				structVis.structName = n.Name.Name
				structVis.pack = vis.pack
				ast.Walk(structVis, typeSpec)
			}
			return vis
		case *ast.InterfaceType:
			if modeExist(vis.modes, Interfaces) {
				ifaceVis := NewInterfaceVisitor(vis.src, vis.Collection)
				ifaceVis.ifaceName = n.Name.Name
				ifaceVis.pack = vis.pack
				ast.Walk(ifaceVis, typeSpec)
			}
			return vis
		default:
			if modeExist(vis.modes, Customs) {
				customVis := NewCustomVisitor(vis.src, vis.Collection)
				customVis.pack = vis.pack
				ast.Walk(customVis, n)
			}
			return vis
		}
	case *ast.FuncDecl:
		if modeExist(vis.modes, Functions) {
			funcVis := NewFuncVisitor(vis.src, vis.Collection)
			funcVis.pack = vis.pack
			ast.Walk(funcVis, n)
		}
	}
	return vis
}
func (vis *ImportVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch importSpec := node.(type) {
	case *ast.ImportSpec:
		vis.importsPaths = append(vis.importsPaths, strings.Trim(importSpec.Path.Value, "\""))
		var (
			imprt *units.Import

			impName string
			impId   int
		)

		if importSpec.Name == nil {

			parts := strings.Split(importSpec.Path.Value, "/")
			impName = parts[len(parts)-1]
			impName = strings.Trim(impName, "\"")
			impId = id(impName + "import")

			imprt = units.NewImport(impId, impName)
			imprt.Value = importSpec.Path.Value

			_ = vis.Collector.Add(vis.pack, imprt)
			return vis
		}

		switch importSpec.Name.Name {
		case ".", "_":
			parts := strings.Split(importSpec.Path.Value, "/")
			impName = parts[len(parts)-1]
			impName = strings.TrimRight(impName, "\"")
			impId = id(impName)

			imprt = units.NewImport(impId, impName)
			imprt.Value = importSpec.Path.Value

			_ = vis.Collector.Add(vis.pack, imprt)

			return vis
		default:
			impName = strings.Trim(importSpec.Name.Name, "\"")
			impId = id(impName)

			imprt = units.NewImport(impId, impName)

			if vis.pack != nil {
				_ = vis.Collector.Add(vis.pack, imprt)
			}

			return vis
		}

	}
	return vis
}
func (vis *StructsVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch structType := node.(type) {
	case *ast.StructType:
		var (
			structUnit *units.Structure
			structName string
			structId   int
		)

		if vis.structName != "" {
			structName = vis.structName
			structId = id(structName)
			structUnit = units.NewStruct(structId, structName)
			structUnit.Fields = vis.getFields(structType.Fields, structUnit)
			structUnit.IsExported = exported(structName)
			_ = vis.Collector.Add(vis.pack, structUnit)
			if vis.parent != nil {
				_ = vis.Collector.Add(vis.parent, structUnit)
			}
		}
	}
	return vis
}
func (vis *InterfacesVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch interfaceType := node.(type) {
	case *ast.InterfaceType:
		ifaceUnit := units.NewIface(id(vis.ifaceName), vis.ifaceName)

		if vis.pack != nil {
			_ = vis.Collector.Add(vis.pack, ifaceUnit)
		}

		for _, method := range interfaceType.Methods.List {
			switch methodType := method.Type.(type) {
			case *ast.FuncType:
				var (
					methodName = method.Names[0].Name
					methodId   = id(methodName)
					exported   = exported(methodName)
					methodUnit = units.NewMethod(methodId, methodName)
				)

				methodUnit.IsExported = exported
				methodUnit.InInterfaceDecl = true

				s, e := methodType.Pos()-1, methodType.End()-1
				methodUnit.Signature = methodName + vis.src[s:e]

				methodUnit.Type = types.Init("func " + methodUnit.Signature)

				if methodType.Params != nil {
					methodUnit.Parameters = getParamsOrReturns(methodType.Params.List, vis.src)
				}
				if methodType.Results != nil {
					methodUnit.Returns = getParamsOrReturns(methodType.Results.List, vis.src)
				}

				_ = vis.Collector.Add(ifaceUnit, methodUnit)
				if vis.parent != nil {
					//println(vis.parent.GetName())
					_ = vis.Collector.Add(vis.parent, ifaceUnit)
				}
			case *ast.Ident:
				if methodType.Obj == nil {
					return vis
				}
				if methodType.Obj.Decl != nil {
					switch tspec := methodType.Obj.Decl.(type) {
					case *ast.TypeSpec:
						switch nIface := tspec.Type.(type) {
						case *ast.InterfaceType:
							ifaceVis := &InterfacesVisitor{
								src:       vis.src,
								ifaceName: tspec.Name.Name,
								pack:      vis.pack,
								parent:    ifaceUnit,
								Collector: vis.Collector,
							}
							ast.Walk(ifaceVis, nIface)
						}
					}
				}
			}
		}
	}

	return vis
}
func (vis *CustomsVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch custom := node.(type) {
	case *ast.TypeSpec:
		switch custom.Type.(type) {
		case *ast.StructType:
			return vis
		case *ast.InterfaceType:
			return vis
		}

		var (
			customUnit = units.NewCustom(id(custom.Name.Name), custom.Name.Name)
		)

		customUnit.IsExported = exported(customUnit.GetName())
		if vis.pack != nil {
			_ = vis.Collector.Add(vis.pack, customUnit)
		}
		if vis.parent != nil {
			_ = vis.Collector.Add(vis.parent, customUnit)
		}

		return vis
	}
	return vis
}
func (vis *FuncsVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch funcDecl := node.(type) {
	case *ast.FuncDecl:
		var (
			funcName = funcDecl.Name.Name
			funcId   = id(funcName)

			signStart, signEnd = funcDecl.Name.Pos() - 1, funcDecl.Type.End() - 1
			signature          = "func " + vis.src[signStart:signEnd]

			typ = types.Init(signature)

			params  = map[string]types.Type{}
			returns = map[string]types.Type{}

			bodyStart, bodyEnd = funcDecl.Body.Pos() - 1, funcDecl.Body.End() - 1
			body               = vis.src[bodyStart:bodyEnd]
		)

		if funcDecl.Type.Params != nil {
			params = getParamsOrReturns(funcDecl.Type.Params.List, vis.src)
		}
		if funcDecl.Type.Results != nil {
			returns = getParamsOrReturns(funcDecl.Type.Results.List, vis.src)
		}

		switch funcDecl.Recv.NumFields() {
		case 0:
			funcUnit := units.NewFunc(funcId, funcName)
			funcUnit.IsExported = exported(funcName)
			funcUnit.Signature = signature
			funcUnit.Returns = returns
			funcUnit.Parameters = params
			funcUnit.FuncBody = body
			funcUnit.Type = typ

			if vis.pack != nil {
				_ = vis.Collection.Add(vis.pack, funcUnit)
			}
		case 1:
			s, e := funcDecl.Recv.List[0].Type.Pos()-1, funcDecl.Recv.List[0].End()-1
			var ownerType = strings.TrimLeft(vis.src[s:e], "*")
			var ownerId = id(ownerType)

			var owner = units.Init(ownerId, ownerType, units.GoUnknown)

			funcUnit := units.NewMethod(funcId, funcName)
			funcUnit.IsExported = exported(funcName)
			funcUnit.Signature = signature
			funcUnit.Returns = returns
			funcUnit.Parameters = params
			funcUnit.Body = body
			funcUnit.Type = typ

			_ = vis.Collection.Add(owner, funcUnit)
			if vis.pack != nil {
				_ = vis.Collection.Add(vis.pack, funcUnit)
			}
		}
	}

	return vis
}
