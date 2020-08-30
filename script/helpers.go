package script

import (
	"fmt"
	"github.com/genome-engine/genome/funcs"
	"github.com/genome-engine/genome/temp_env"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	t "text/template"
)

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
	var delim = s.getDelimiters(s.Generate.Path)

	if strings.Contains(s.Generate.Path, delim) {
		dirs := strings.Split(s.Generate.Path, delim)
		lastIndex = len(dirs) - 1

		return strings.Join(dirs[:lastIndex], delim)
	}

	return s.Generate.Path
}

func (s *Script) getDelimiters(fullPath string) string {
	if strings.Contains(fullPath, "/") {
		return "/"
	}
	return "\\"
}

func (s *Script) createDirs(filename string) error {
	var delim = s.getDelimiters(filename)
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

func walkFunc(source *strings.Builder, d Delimiter) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		src, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if _, err := t.New("").Delims(d.Delimiters()).Funcs(funcs.Funcs()).Parse(string(src)); err == nil {
			source.Write(src)
		}

		return nil
	}
}
