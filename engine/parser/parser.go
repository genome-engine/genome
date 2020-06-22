package parser

import (
	"fmt"
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/parser/visitors"
	"github.com/genome-engine/genome/engine/units"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Parser struct {
	Config
	imports []string
}

func New(config Config) *Parser {
	return &Parser{Config: config}
}

func (p *Parser) Parse() error {
	err := filepath.Walk(p.Path, p.walkFunc())

	if err != nil {
		return err
	}

	if !p.InspectImplements {
		return nil
	}

	if p.ImplementsCollection == nil {
		return fmt.Errorf("Collection of interfecs not handed over ")
	}

	p.findImplements()

	return nil
}

func (p *Parser) walkFunc() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		separator := string(filepath.Separator)
		pathParts := strings.Split(path, separator)

		packMainDir := pathParts[len(pathParts)-2]

		src, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		set := token.NewFileSet()
		file, err := parser.ParseFile(set, "", src, 0)

		if err != nil {
			return err
		}

		genVisitor := visitors.NewGeneralVisitor(path, packMainDir, string(src), p.GeneralCollection, p.Modes...)

		ast.Walk(genVisitor, file)

		p.imports = genVisitor.FoundImports()

		return nil
	}
}
func (p *Parser) getMethodsWithOwners() map[int][]units.Unit {
	//            ownerId->[]method
	var methods = map[int][]units.Unit{}
	var potentialOwners = p.GeneralCollection.SearchBySelectors([]units.Selector{units.GoStruct, units.GoCustom}...)
	for _, owner := range potentialOwners {
		if m, e := p.GeneralCollection.SearchChildren(owner, units.GoMethod); m != nil && e == nil {
			methods[owner.GetId()] = m
		}
	}
	return methods
}
func (p *Parser) findImplements() {
	var methods = p.getMethodsWithOwners()
	var owner units.Unit
	var iface units.Unit
	var ifaceMethods []units.Unit

	for ownerId, m := range methods {
		owner = p.GeneralCollection.SearchById(ownerId)
		for _, method := range m {
			ifaces, _ := p.ImplementsCollection.SearchParents(method, units.GoInterface)
			if len(ifaces) == 0 {
				continue
			}

			iface = ifaces[0]
			ifaceMethods, _ = p.ImplementsCollection.SearchChildren(iface, units.GoMethod)

			if compareMethods(ifaceMethods, m) {
				_ = p.GeneralCollection.Add(owner, iface)
			}
		}
	}
}

func compareMethods(ifaceMethods, ownerMethods []units.Unit) bool {
	if len(ifaceMethods) > len(ownerMethods) {
		return false
	}
	var (
		necessaryMatches = len(ifaceMethods)
		actualMatches    int
	)

	for _, ifaceMethod := range ifaceMethods {
		if collection.UnitExist(ownerMethods, ifaceMethod) {
			actualMatches++
		}
	}

	return necessaryMatches == actualMatches
}
