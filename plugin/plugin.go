package plugin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

const Name = "out"
const Func = "Wrapper"

func Load() map[string]interface{} {
	var loading map[string]interface{}

	_, err := os.Stat("./plugins")
	if os.IsNotExist(err) {
		return nil
	}

	if err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
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

			loading[fmt.Sprintf("%v", name)] = wrapper.(func() interface{})()
		}
		return nil
	}); err != nil {
		log.Print(err.Error())
		return loading
	}

	return loading
}
