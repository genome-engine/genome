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
	"regexp"
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

	config := p.Config{
		Collection: *collection,
		Path:       s.Parse,
	}

	err := p.New(config, s.Logs).Parse()
	if err != nil {
		s.log("Error from parsing.Exit with error.")
		return err
	}

	s.log("Parsing complete.Collection received.")
	s.result.parse = *collection

	return nil
}

func (s *Script) execTemp() error {
	s.log("Templating run.")

	var w = &strings.Builder{}

	src, err := ioutil.ReadFile(s.Template)
	if err != nil {
		s.log("Error from file reading.Exit with error.")
		return err
	}

	s.log("Initializing template for execution.")
	temp, err := t.New("").Delims(s.Delimiter.Delimiters()).Funcs(funcs.Funcs).Parse(string(src))
	if err != nil {
		s.log("Error from template parsing.Exit with error.")
		return err
	}
	s.log("Executing template.")
	env := temp_env.New(s.parse, s.Logs)
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

func (s *Script) isMultiFileSource() bool {
	return regexp.MustCompile(fmt.Sprintf("(%v|%v)[\\w+/\\.]+", temp_env.FileStart, temp_env.FileEnd)).MatchString(s.temp)
}

func (s *Script) splitSource() map[string]string {
	var (
		exp = func(s string) *regexp.Regexp { return regexp.MustCompile(fmt.Sprintf("%v[\\w+/\\.]+", s)) }

		start = exp(temp_env.FileStart) //#file-start: filename.extension | if not extension -> extension = .go
		end   = exp(temp_env.FileEnd)   //#file-end: filename.extension

		startIdx = start.FindAllStringIndex(s.temp, -1)
		endIdx   = end.FindAllStringIndex(s.temp, -1)

		startBuffer = map[string]int{}    //filename -> end index
		dataBuffer  = map[string]string{} //filename -> source

		sourceResidue strings.Builder
		min, max      int
		zeroWas       bool
	)

	for _, idx := range startIdx {
		if len(idx) == 0 {
			continue
		}

		if min == 0 && !zeroWas {
			min = idx[0]
		}

		if min > idx[0] {
			min = idx[0]
		}

		filename := strings.Split(s.temp[idx[0]-1:idx[1]], ":")[1]
		startBuffer[filename] = idx[1]
	}

	for _, idx := range endIdx {
		if len(idx) == 0 {
			continue
		}

		if idx[1] > max {
			max = idx[1]
		}

		filename := strings.Split(s.temp[idx[0]-1:idx[1]], ":")[1]
		if index, ok := startBuffer[filename]; ok {
			dataBuffer[filename] = s.temp[index:idx[0]]
		}
	}

	sourceResidue.WriteString(s.temp[:min])
	sourceResidue.WriteString(s.temp[max:])
	s.temp = sourceResidue.String()

	return dataBuffer
}

func (s *Script) getDirsPath() string {
	var lastIndex int
	var delim = s.getDelim(s.Generate.Path)

	if strings.Contains(s.Generate.Path, delim) {
		dirs := strings.Split(s.Generate.Path, delim)
		lastIndex = len(dirs) - 1

		return strings.Join(dirs[:lastIndex], delim)
	}

	return s.Generate.Path
}

func (s *Script) getDelim(fullPath string) string {
	if strings.Contains(fullPath, "/") {
		return "/"
	}
	return "\\"
}

func (s *Script) createDirs(filename string) error {
	var delim = s.getDelim(filename)
	var dirsPath = s.getDirsPath()

	if strings.Contains(filename, delim) {
		elements := strings.Split(filename, delim)

		for _, dir := range elements[:len(elements)-1] {
			if err := os.Mkdir(path.Join(dirsPath, dir), os.ModePerm); err != nil && !os.IsExist(err) {
				return err
			}
		}
	}
	return nil
}
