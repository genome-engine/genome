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
		logs  bool
		count int
		Config
	}
	Config struct {
		Interfaces collection.Collection
		Collection collection.Collection
		Modes      []visitors.VisitMode
		Path       string
	}
)

func New(config Config, logs bool) *Parser {
	config.Collection.ChangeQualifier("Parsing")
	if config.Modes == nil {
		config.Modes = visitors.AllModes
	}
	return &Parser{Config: config, logs: logs}
}

func (p *Parser) Parse() error {
	if err := p.Collection.Merge(p.Interfaces); err != nil {
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
		file, err := parser.ParseFile(set, "", src, parser.ParseComments)

		if err != nil {
			return err
		}

		genVisitor := visitors.NewGeneralVisitor(path, packMainDir, string(src), p.Collection, p.Modes...)

		ast.Walk(genVisitor, file)
		return nil
	}
}
