package interface_finder

import (
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/collection/formatter"
	"github.com/genome-engine/genome/engine/parser/visitors"
	"github.com/genome-engine/genome/engine/units"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const Filename = "interfaces"

type DefaultInterfaceFinder struct {
	collection.Collector
	locations       map[string]string
	preservingPaths map[string]collection.UnitsMap
	savePath        string
}

func New(savePath string, locations map[string]string, collector collection.Collector) *DefaultInterfaceFinder {
	return &DefaultInterfaceFinder{
		savePath:        savePath,
		locations:       locations,
		Collector:       collector,
		preservingPaths: map[string]collection.UnitsMap{},
	}
}

func (finder *DefaultInterfaceFinder) Start() error {
	// iteration in transmitted locations
	for lockName, location := range finder.locations {
		if err := filepath.Walk(location, finder.walkFunc()); err != nil {
			return err
		}

		lockName = path.Join(finder.savePath, lockName, Filename+formatter.FileExtension)
		finder.preservingPaths[lockName] = finder.UnitsMap()
		finder.Clear()
	}

	return nil
}
func (finder *DefaultInterfaceFinder) SaveResult() error {
	newFormatter := formatter.NewFormatter(finder.Collector)
	for lock := range finder.locations {
		if err := os.Mkdir(path.Join(finder.savePath, lock), os.ModePerm); os.IsExist(err) {
			continue
		}
	}
	for sPath, objMap := range finder.preservingPaths {

		content := []byte(newFormatter.ObjMapToText(objMap))

		if err := ioutil.WriteFile(sPath, content, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
func (finder *DefaultInterfaceFinder) walkFunc() filepath.WalkFunc {
	var (
		isGoFile     = func(path string) bool { return strings.HasSuffix(path, ".go") }
		isGoTestFile = func(path string) bool { return strings.HasSuffix(path, "_test.go") }
	)

	return func(path string, info os.FileInfo, err error) error {
		if ContainsTestDirs(path) || info.IsDir() || !isGoFile(path) || isGoTestFile(path) {
			return nil
		}

		pathParts := strings.Split(path, string(filepath.Separator))
		packMainDir := pathParts[len(pathParts)-2]

		// get go file.txt content
		src, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// file.txt set and visitor initialization on new iteration filepath.Walk
		var (
			set     = token.NewFileSet()
			visitor = visitors.NewGeneralVisitor(
				path, packMainDir, string(src),
				collection.New(),
				visitors.Interfaces,
			)
		)

		// parsing
		f, err := parser.ParseFile(set, path, src, 0)
		if err != nil {
			return err
		}

		// collection of interfaces, methods and packages
		ast.Walk(visitor, f)

		// if no selectors of interest were found in the file.txt, proceed to the next file.txt.
		if len(visitor.Collection.Search(units.GoInterface)) == 0 {
			return nil
		}

		// transfer of units collected by visitor to interface_finder
		if err := finder.Collector.Merge(visitor.Collection); err != nil {
			return err
		}

		return nil
	}
}

//if tests was found on the way or testdata returns true.
func ContainsTestDirs(path string) bool {
	testDirs := []string{"testdata", "tests"}

	for _, dir := range testDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}

	return false
}
