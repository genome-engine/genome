package script

import (
	"fmt"
	c "github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/generator"
	p "github.com/genome-engine/genome/engine/parser"
	"github.com/genome-engine/genome/funcs"
	"github.com/genome-engine/genome/temp_env"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	t "text/template"
)

func (s *Script) log(info string, args ...interface{}) {
	if !s.Logs {
		return
	}
	s.count++
	fmt.Printf("\t%d.[Script] %v\n", s.count, fmt.Sprintf(info, args...))
}

func (s *Script) Execute() error {
	if err := s.execParse(); err != nil {
		return fmt.Errorf("Parsing Error: %v \n", err.Error())
	}

	s.log("The parsing phase was a success.")

	if err := s.execTemp(); err != nil {
		return fmt.Errorf("\tTemplating Error: %v \n", err.Error())
	}

	s.log("The template filling phase was successful.")

	if err := s.execGen(); err != nil {
		return fmt.Errorf("\t- [Error]Generating Error: %v \n", err.Error())
	}

	s.log("The generation phase was a success.")
	s.log("Script execution was a success.")

	return nil
}

func (s *Script) execParse() error {
	s.log("Parsing running")
	collection := c.New("Script", s.Logs)

	for _, prs := range s.Parses {
		config := p.Config{
			Collection: *collection,
			Path:       prs.Path,
		}

		err := p.New(config, s.Logs).Parse()
		if err != nil {
			s.log("Error from parsing.Exit with error.")
			return err
		}

		s.log("Parsing complete.Collection received.")
	}

	s.result.collection = *collection

	return nil
}

func (s *Script) execTemp() error {
	s.log("Templating run.")

	var (
		w      = &strings.Builder{}
		source = &strings.Builder{}
	)

	if s.GlobTemps && os.Getenv("GENOME_TEMPS") != "" {
		s.log("The use of global templates is activated.")
		if err := filepath.Walk(os.Getenv("GENOME_TEMPS"), walkFunc(source, s.Delimiter)); err != nil {
			return err
		}
	}

	s.log("Template bonding.")
	for _, template := range s.Templates {
		src, err := ioutil.ReadFile(template.Path)
		if err != nil {
			s.log("Error from file reading.Exit with error.")
			return err
		}
		source.WriteString("\n")
		source.Write(src)
	}

	s.log("Initializing template for execution.")
	temp, err := t.New("").Delims(s.Delimiters()).Funcs(funcs.Funcs()).Parse(source.String())
	if err != nil {
		s.log("Error from template parsing.Exit with error.")
		return err
	}

	s.log("Executing template.")
	env := temp_env.New(s.collection, s.Logs)
	if err := temp.Execute(w, env); err != nil {
		s.log("Error from template execution. Exit with error.")
		return err
	}

	s.temp = strings.TrimFunc(w.String(), func(r rune) bool {
		return r == '\n'
	})

	s.log("Templating complete.")
	return nil
}

func (s *Script) execGen() error {
	if s.isMultiFileSource() {
		var dirsPath = s.getDirsPath()

		for filename, source := range s.splitSource() {
			var info = generator.GenerationInfo{
				Mode:   generator.CreateFile,
				Source: source,
				Path:   path.Join(dirsPath, filename),
			}

			if err := s.createDirs(filename); err != nil {
				return err
			}

			if err := generator.NewGenerator(s.Logs, info).Generate(); err != nil {
				return err
			}
		}

		if s.temp == "" || !strings.Contains(s.Generate.Path, ".") {
			return nil
		}
	}

	s.log("Generating run.")
	var info = generator.GenerationInfo{
		Path: s.Generate.Path, Mode: generator.ToMode(s.Mode), Source: s.temp,
		InsertionLabel: generator.InsertionLabel{LabelName: s.Generate.Label},
	}
	err := generator.NewGenerator(s.Logs, info).Generate()
	if err != nil {
		s.log("Error from generating.Exit with error.")
		return err
	}
	return nil
}
