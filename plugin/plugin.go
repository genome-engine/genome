package plugin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

const PlugEnv = "GENOME_PLUG"
const Name = "Name"
const Func = "Wrap"

func Load() map[string]interface{} {
	var loading = map[string]interface{}{}

	if os.Getenv(PlugEnv) != "" {
		loading = load(os.Getenv(PlugEnv))
	}

	_, err := os.Stat("./plugins")
	if os.IsNotExist(err) {
		return nil
	}

	for name, fun := range load("./plugins") {
		if _, ok := loading[name]; !ok {
			loading[name] = fun
		} else {
			fmt.Printf("\"\\t\\tFunction with name[%v] exist in builtin function list of genome.\", name")
		}
	}

	return loading
}

func load(path string) map[string]interface{} {
	var loading = map[string]interface{}{}

	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".so") {
			p, err := plugin.Open(path)
			if err != nil {
				return fmt.Errorf("Plugin Loading Error: %s\n ", err)
			}

			name, err := p.Lookup(Name)
			if err != nil {
				return fmt.Errorf("Plugin out Loading Error: %s ", err)
			}

			wrapper, err := p.Lookup(Func)
			if err != nil {
				return fmt.Errorf("Plugin Func Loading Error: %v ", err.Error())
			}

			loading[*name.(*string)] = wrapper.(func() interface{})()
		}
		return nil
	}); err != nil {
		log.Print(err.Error())
		return loading
	}

	return loading
}
