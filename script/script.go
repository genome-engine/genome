package script

import (
	"fmt"
	c "github.com/genome-engine/genome/engine/collection"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type (
	Script struct {
		result
		Generate `yaml:"generate"`
		Parse    string `yaml:"parse"`
		Template string `yaml:"template"`
	}

	result struct {
		parse c.Collector
		temp  string
	}

	Generate struct {
		Path string `yaml:"path"`
		Mode string `yaml:"mode"`
		Tag  string `yaml:"tag"`
	}
)

func NewScript(path string) (s *Script, err error) {
	if !strings.HasSuffix(path, ".yaml") {
		return nil, fmt.Errorf("The transmitted file has no yaml extension!\n File:%v ", path)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(src, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
