package parser

import (
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/parser/visitors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type (
	Parser struct {
		Config
	}
	Config struct {
		Interfaces collection.Collector
		Collector  collection.Collector
		Modes      []visitors.VisitMode
		Path       string
	}
)

func New(config Config) *Parser {
	if config.Modes == nil {
		config.Modes = visitors.AllModes
	}
	return &Parser{Config: config}
}

func (p *Parser) Parse() error {
	if err := p.Collector.Merge(p.Interfaces); err != nil {
		return err
	}
	return filepath.Walk(p.Path, p.walkFunc())
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

		genVisitor := visitors.NewGeneralVisitor(path, packMainDir, string(src), p.Collector, p.Modes...)

		ast.Walk(genVisitor, file)

		return nil
	}
}
