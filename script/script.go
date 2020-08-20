package script

import (
	c "github.com/genome-engine/genome/engine/collection"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type (
	Script struct {
		result
		Delimiter `yaml:"delimiters"`
		Generate  `yaml:"generate"`
		Logs      bool   `yaml:"logs"`
		Parse     string `yaml:"parse"`
		Template  string `yaml:"template"`
		count     int
	}

	result struct {
		parse c.Collection
		temp  string
	}

	Generate struct {
		Path  string `yaml:"path"`
		Mode  string `yaml:"mode"`
		Label string `yaml:"label"`
	}
)

func NewScript(path string) (s *Script, err error) {
	var prefix = func(prefix string) bool { return strings.HasPrefix(path, prefix) }
	var notExist = func(path string) bool {
		_, err := os.Stat(path)
		return os.IsNotExist(err)
	}

	if !prefix(".yaml") {
		if !notExist(path + ".yaml") {
			path += ".yaml"
		}
	}

	if !prefix(".yml") {
		if !notExist(path + ".yml") {
			path += ".yml"
		}
	}

	if notExist(path) {
		return nil, os.ErrNotExist
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
