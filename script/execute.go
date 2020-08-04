package script

import (
	"fmt"
	c "github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/generator"
	p "github.com/genome-engine/genome/engine/parser"
	"github.com/genome-engine/genome/temp_env"
	"io/ioutil"
	"os/exec"
	"strings"
	t "text/template"
)

func (s *Script) Execute() error {
	if err := s.execParse(); err != nil {
		return fmt.Errorf("Parsing Error: %v \n", err.Error())
	}

	if err := s.execTemp(); err != nil {
		return fmt.Errorf("Templating Error: %v \n", err.Error())
	}

	if err := s.execGen(); err != nil {
		return fmt.Errorf("Generating Error: %v \n", err.Error())
	}

	if err := exec.Command("go", "fmt", s.Generate.Path).Run(); err != nil {
		return fmt.Errorf("Formating Error: %v \n", err.Error())
	}

	return nil
}

func (s *Script) execParse() error {
	collection := c.New()

	config := p.Config{
		Collector: collection,
		Path:      s.Parse,
	}

	err := p.New(config).Parse()
	if err != nil {
		return err
	}

	s.result.parse = collection

	return nil
}

func (s *Script) execTemp() error {
	var w = &strings.Builder{}
	src, err := ioutil.ReadFile(s.Template)
	if err != nil {
		return err
	}

	funcMap := t.FuncMap{
		"lower":  strings.ToLower,
		"trim_r": strings.TrimRight,
		"trim_l": strings.TrimLeft,
		"trim":   strings.Trim,
		"f":      temp_env.NewFilterer,
	}

	temp, err := t.New("").Funcs(funcMap).Parse(string(src))

	if err != nil {
		return err
	}
	env := temp_env.New(s.parse)

	err = temp.Execute(w, env)
	if err != nil {
		return err
	}

	s.result.temp = w.String()

	return nil
}

func (s *Script) execGen() error {
	var info generator.GenerationInfo
	info.Path = s.Generate.Path
	info.Mode = generator.CreateFile
	info.Source = s.result.temp

	return generator.NewGenerator(info).Generate()
}
