package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func (i *GenerationInfo) log(info string, args ...interface{}) {
	if !i.logs {
		return
	}
	i.count++
	fmt.Printf("\t\t%d.[Generator]%v\n", i.count, fmt.Sprintf(info, args...))
}

func (l *InsertionLabel) log(info string, args ...interface{}) {
	if !l.logs {
		return
	}
	l.count++
	fmt.Printf("\t\t\t%d %v.\n", l.count, fmt.Sprintf(info, args...))
}

const InsertPrefix = "#genome-insert:"
const EndInsertPrefix = "#genome-insert-end"

type (
	Mode int

	Generator struct {
		logs  bool
		infos []GenerationInfo
	}

	GenerationInfo struct {
		Mode
		logs   bool
		Source string
		Path   string
		InsertionLabel
		count int
	}

	InsertionLabel struct {
		count     int
		logs      bool
		LabelName string
		insertionBuffer
	}

	insertionBuffer struct {
		begin, end string
	}
)

func NewGenerator(logs bool, infos ...GenerationInfo) *Generator {
	return &Generator{infos: infos, logs: logs}
}

const (
	CreateFile Mode = iota
	InsertToFile
)

func ToMode(s string) Mode {
	if m, ok := map[string]Mode{"insert": InsertToFile, "create": CreateFile}[s]; ok {
		return m
	}
	return CreateFile
}

func (m Mode) String() string { return map[Mode]string{InsertToFile: "insert", CreateFile: "write"}[m] }

func (g *Generator) Generate() error {
	for _, info := range g.infos {
		info.logs = g.logs
		info.log("Generation start with mode: %v", info.Mode.String())
		switch info.Mode {
		case CreateFile:
			file, err := os.Create(info.Path)
			if err != nil {
				info.log("Error from creating.Exit with error.")
				return err
			}

			info.log("Writing source into file")
			_, err = file.WriteString(info.Source)

			_ = file.Close()

			info.formatting()

			return err
		case InsertToFile:
			info.log("Reading original file: %v", info.Path)
			originalSource, err := ioutil.ReadFile(info.Path)
			if err != nil {
				fmt.Printf("\t\tError from reading.Exit with error.")
				return err
			}
			info.log("Filling insertion buffer")
			info.InsertionLabel.logs = info.logs
			err = info.InsertionLabel.fillBuffer(string(originalSource))
			if err != nil {
				info.log("Error from filling. Exit with error")
				return err
			}

			buffer := info.insertionBuffer

			newSource := buffer.begin + info.Source + "\n" + buffer.end
			if newSource == "" {
				info.log("New source is empty.")
			}

			info.log("Removing old file: %v", info.Path)
			err = os.Remove(info.Path)
			if err != nil {
				info.log("Error from removing. Exit with error.")
				return err
			}
			info.log("Recreating old file with new source.")
			file, err := os.Create(info.Path)
			if err != nil {
				info.log("Error from creating.Exit with error.")
				return err
			}

			info.log("Writing new source")
			_, err = file.WriteString(newSource)
			if err != nil {
				info.log("Error from writing. Exit with error.")
				return err
			}

			_ = file.Close()

			info.formatting()
			return nil
		}
	}

	return nil
}

func (l *InsertionLabel) fillBuffer(originalSource string) error {
	var b strings.Builder
	var found bool

	lines := strings.Split(originalSource, "\n")

	l.log("Iterating from source lines")
	for i, line := range lines {
		if !found {
			b.WriteString(line)
		}
		switch {
		case strings.Contains(line, InsertPrefix+l.LabelName):
			l.log("Label %v was found", InsertPrefix+l.LabelName)

			found = true
			l.insertionBuffer.begin = b.String()
			b.Reset()
			continue
		case strings.Contains(line, EndInsertPrefix):
			l.log("%v was found", EndInsertPrefix)
			if !found {
				l.log("Label %v not was found", InsertPrefix+l.LabelName)
				return fmt.Errorf("The label for the end of insertion is set," +
					" but it was not possible to find a named label for the beginning of insertion in this file.txt. ")
			}

			l.end = strings.Join(lines[i:], "\n")
			continue
		default:
			continue
		}

	}
	return nil
}

func (i *GenerationInfo) formatting() {
	i.log("Executing go fmt %v", i.Path)
	err := exec.Command("go", "fmt", i.Path).Run()
	if err != nil {
		i.log("[Warning] fmt error.")
	}
}
