package visitors

import (
	"github.com/genome-engine/genome/engine/identifier"
	"github.com/genome-engine/genome/engine/units"
	"go/ast"
	"go/token"
	"strings"
)

func (vis *StructsVisitor) getFields(fields *ast.FieldList, structure units.Unit) []units.StructField {
	var structFields []units.StructField

	for _, field := range fields.List {
		structField := units.StructField{}

		var comment string

		if doc := field.Doc; doc != nil {
			comment = vis.src[doc.Pos()-1 : doc.End()-1]
		}

		ts, te := field.Type.Pos()-1, field.Type.End()-1
		fieldType := vis.src[ts:te]

		if field.Type != nil {
			switch t := field.Type.(type) {
			case *ast.Ident:
				if t.Obj != nil {
					switch typeDecl := t.Obj.Decl.(type) {
					case *ast.TypeSpec:
						switch structType := typeDecl.Type.(type) {
						case *ast.StructType:
							fieldType = typeDecl.Name.Name

							structVis := &StructsVisitor{
								parent:     structure,
								structName: typeDecl.Name.Name,
								comment:    comment,
								src:        vis.src,
								pack:       vis.pack,
								Collector:  vis.Collector,
							}

							ast.Walk(structVis, structType)

							break
						case *ast.InterfaceType:
							fieldType = typeDecl.Name.Name

							ifaceVis := &InterfacesVisitor{
								src:       vis.src,
								ifaceName: typeDecl.Name.Name,
								pack:      vis.pack,
								comment:   comment,
								isField:   true,
								parent:    structure,
								Collector: vis.Collector,
							}
							ast.Walk(ifaceVis, structType)
							break
						default:
							fieldType = typeDecl.Name.Name

							customVis := &CustomsVisitor{
								src:       vis.src,
								pack:      vis.pack,
								comment:   comment,
								parent:    structure,
								Collector: vis.Collector,
							}

							ast.Walk(customVis, typeDecl)
							break
						}
					}
				}
			case *ast.SelectorExpr:
				s, e := t.Pos()-1, t.End()-1
				fieldType = vis.src[s:e]
				break
			}
		}

		structField.Type = fieldType
		structField.Comment = comment
		if field.Tag != nil {
			structField.Tag = field.Tag.Value
		}

		switch {
		case len(field.Names) == 0:
			structField.Name = fieldType
			structField.IsExported = exported(fieldType)
			structFields = append(structFields, structField)
			break
		case len(field.Names) == 1:
			structField.Name = field.Names[0].Name
			structField.IsExported = exported(structField.Name)
			structFields = append(structFields, structField)
			break
		case len(field.Names) > 1:
			if structField.Tag != "" {
				structField.Name = field.Names[0].Name
				structFields = append(structFields, structField)
				for i := 1; i <= len(field.Names)-1; i++ {
					fldName := field.Names[i].Name
					exported := exported(fldName)
					fld := units.StructField{Name: fldName, IsExported: exported}
					fld.Type = fieldType
					structFields = append(structFields, fld)
				}

				break
			}
			for _, ident := range field.Names {
				fldName := ident.Name
				exported := exported(fldName)
				fld := units.StructField{Name: fldName, IsExported: exported}
				fld.Type = fieldType
				structFields = append(structFields, fld)
			}
		}
	}

	return structFields
}

func getParams(fields []*ast.Field, src string) []units.Param {
	var list []units.Param
	var parameter units.Param

	for _, result := range fields {
		s, e := result.Type.Pos()-1, result.Type.End()-1
		parameter.Type = src[s:e]

		switch {
		case len(result.Names) == 0:
			parameter.Name = parameter.Type
			list = append(list, parameter)
		case len(result.Names) == 1:
			parameter.Name = result.Names[0].Name
			list = append(list, parameter)
		case len(result.Names) > 1:
			for _, ident := range result.Names {
				parameter.Name = ident.Name
				list = append(list, parameter)
			}
		}
	}

	return list
}

func id(name string) int {
	return identifier.GenerateID(name)
}

func exported(name string) bool {
	var typ = name

	if strings.Contains(name, ".") {
		typ = strings.Split(name, ".")[1]
	}

	return ast.IsExported(typ)
}

func (vis *GeneralVisitor) varsHandle(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		switch variable := spec.(type) {
		case *ast.ValueSpec:
			var lenNames = len(variable.Names) - 1
			var comment string

			if doc := variable.Comment; doc != nil {
				comment = vis.src[doc.Pos()-1 : doc.End()-1]
			}

			for i := 0; i <= lenNames; i++ {
				var varName = variable.Names[i].Name
				var varType = variable.Type

				var varUnit = units.NewVar(id(varName), varName)
				varUnit.IsExported = exported(varName)
				varUnit.Comment = comment
				if varType != nil {
					s, e := varType.Pos()-1, varType.End()-1
					varUnit.Type = vis.src[s:e]

					_ = vis.Collection.Add(vis.pack, varUnit)
					continue
				}

				switch value := variable.Values[i].(type) {
				case *ast.BasicLit:
					typName := strings.ToLower(value.Kind.String())
					if typName == "float" {
						typName += "64"
					}

					typ := typName
					varUnit.Type = typ
					_ = vis.Collection.Add(vis.pack, varUnit)
					continue
				case *ast.CompositeLit:
					s, e := value.Type.Pos()-1, value.Type.End()-1
					typName := vis.src[s:e]

					typ := typName
					varUnit.Type = typ
					_ = vis.Collection.Add(vis.pack, varUnit)
					continue
				case *ast.SelectorExpr:
					s, e := value.Pos()-1, value.End()-1
					varUnit.Type = vis.src[s:e]

					_ = vis.Collection.Add(vis.pack, varUnit)
					continue
				}
			}
		}
	}
}

func (vis *GeneralVisitor) constHandle(n *ast.GenDecl) {
	if n.Lparen != token.NoPos && n.Rparen != token.NoPos {
		var (
			enumFound bool
			enumType  string

			enums []units.Unit
			enum  *units.Const
		)

		for _, spec := range n.Specs {
			switch constant := spec.(type) {
			case *ast.ValueSpec:
				var (
					constName = constant.Names[0].Name
					value     string
					comment   string
				)

				if doc := constant.Comment; doc != nil {
					comment = vis.src[doc.Pos()-1 : doc.End()-1]
				}

				if constant.Values != nil {
					start, end := constant.Values[0].Pos()-1, constant.Values[0].End()-1
					value = vis.src[start:end]

					switch strings.HasSuffix(value, "iota") {
					case true:
						if enumFound && len(enums) > 0 {
							_ = vis.Collection.Add(vis.pack, enums...)

							enums = []units.Unit{}
						}

						enumFound = true
						enumType = "int"

						if constant.Type != nil {
							start, end = constant.Type.Pos()-1, constant.Type.End()-1
							typeName := vis.src[start:end]
							enumType = typeName
						}

						enum = units.NewConst(id(constName), constName)
						enum.Comment = comment
						enum.Type = enumType
						enum.IsExported = exported(constName)
						enum.Enum = true
						enums = append(enums, enum)

						continue
					case false:
						var constType string
						var constId = id(constName)

						constUnit := units.NewConst(constId, constName)
						constUnit.IsExported = exported(constName)
						constUnit.Comment = comment

						if constant.Type != nil {
							start, end = constant.Type.Pos()-1, constant.Type.End()-1
							typeName := vis.src[start:end]
							constType = typeName
							constUnit.Type = constType

							_ = vis.Collection.Add(vis.pack, constUnit)
							continue
						}

						switch valType := constant.Values[0].(type) {
						case *ast.BasicLit:
							typeName := strings.ToLower(valType.Kind.String())
							if typeName == "float" {
								typeName += "64"
							}
							constType = typeName
							constUnit.Type = constType

							_ = vis.Collection.Add(vis.pack, constUnit)
							continue
						}
					}
				}

				if constant.Values == nil && enumFound && enumType != "" {
					enum = units.NewConst(id(constName), constName)
					enum.Type = enumType
					enum.IsExported = exported(constName)
					enum.Enum = true
					enums = append(enums, enum)
					continue
				}
			}
		}

		if len(enums) > 0 {
			_ = vis.Collection.Add(vis.pack, enums...)
		}
	}
}
