package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const InsertPrefix = "#genome-insert:"
const EndInsertPrefix = "#genome-insert-end"

type (
	Mode int

	Generator struct {
		infos []GenerationInfo
	}

	GenerationInfo struct {
		Mode
		Source string
		Path   string
		InsertionLabel
	}

	InsertionLabel struct {
		LabelName string
		insertionBuffer
	}

	insertionBuffer struct {
		begin, center, src, end string
	}
)

func NewGenerator(infos ...GenerationInfo) *Generator {
	return &Generator{infos: infos}
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
		switch info.Mode {
		case CreateFile:
			file, err := os.Create(info.Path)
			if err != nil {
				return err
			}

			_, err = file.WriteString(info.Source)
			return err
		case InsertToFile:
			originalSource, err := ioutil.ReadFile(info.Path)
			if err != nil {
				return err
			}
			err = info.InsertionLabel.fillBuffer(string(originalSource), info.Source)
			if err != nil {
				return err
			}

			buffer := info.insertionBuffer

			newSource := buffer.begin + buffer.center + buffer.src + buffer.end

			err = os.Remove(info.Path)
			if err != nil {
				return err
			}

			file, err := os.Create(info.Path)
			if err != nil {
				return err
			}

			_, err = file.WriteString(newSource)
			if err != nil {
				return err
			}

			_ = file.Close()

			return exec.Command("go", []string{"fmt", info.Path}...).Run()
		}
	}

	return nil
}

func (l *InsertionLabel) fillBuffer(originalSource, addonSource string) error {
	var beginIndex int
	l.src = "\n" + addonSource + "\n"

	lines := strings.Split(originalSource, "\n")

	for i, line := range lines {
		switch {
		case strings.Contains(line, InsertPrefix+l.LabelName):
			beginIndex = i + 1
			l.insertionBuffer.begin = strings.Join(lines[:i+1], "\n")
			continue
		case strings.Contains(line, EndInsertPrefix):
			if beginIndex == 0 && l.insertionBuffer.begin == "" {
				return fmt.Errorf("The label for the end of insertion is set," +
					" but it was not possible to find a named label for the beginning of insertion in this file.txt. ")
			}

			l.center = strings.Join(lines[beginIndex:i], "\n")
			l.end = strings.Join(lines[i:], "\n")
			continue
		default:
			continue
		}

	}
	return nil
}
