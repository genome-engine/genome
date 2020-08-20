package formatter

import (
	c "github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/units"
	"strconv"
	"strings"
)

const FileExtension string = ".collection"

type Formatter struct {
	c.Collection
}

func NewFormatter(collector c.Collection, logs bool) *Formatter {
	if collector == nil {
		collector = c.New("Formatting", logs)
	}
	return &Formatter{Collection: collector}
}

func (f *Formatter) ToText() string {
	objMap := f.UnitsMap()

	var text strings.Builder

	for root, children := range objMap {
		text.WriteString("id:" + strconv.FormatInt(int64(root.GetId()), 10))
		text.WriteString(" name:" + root.GetName())
		text.WriteString(" selector:" + root.GetSelector().Name() + "\n")

		for _, child := range children {
			text.WriteString("\tid:" + strconv.FormatInt(int64(child.GetId()), 10))
			text.WriteString(" name:" + child.GetName())
			text.WriteString(" selector:" + child.GetSelector().Name() + "\n")
		}
	}

	return text.String()
}

func (f *Formatter) ToCollection(collectionString string) c.Collection {
	if collectionString == "" {
		return f.Collection
	}

	var root units.Unit

	lines := strings.Split(collectionString, "\n")

	for _, line := range lines {
		prefix := strings.HasPrefix(line, "\t")
		u := convertToUnit(line)

		if !prefix {
			root = u
			_ = f.Add(root)
			continue
		} else {
			_ = f.Add(root, u)
		}
	}

	return f.Collection
}

func (Formatter) ObjMapToText(objectMap map[units.Unit][]units.Unit) string {
	var text strings.Builder

	for root, children := range objectMap {
		text.WriteString("id:" + strconv.FormatInt(int64(root.GetId()), 10))
		text.WriteString(" name:" + root.GetName())
		text.WriteString(" selector:" + root.GetSelector().Name() + "\n")

		for _, child := range children {
			text.WriteString("\tid:" + strconv.FormatInt(int64(child.GetId()), 10))
			text.WriteString(" name:" + child.GetName())
			text.WriteString(" selector:" + child.GetSelector().Name() + "\n")
		}
	}

	return text.String()
}

func convertToUnit(line string) units.Unit {
	var (
		id       int
		err      error
		name     string
		selector units.Selector
	)

	line = strings.TrimLeft(line, "\t")
	elements := strings.Split(line, " ")

	for _, element := range elements {
		parts := strings.Split(element, ":")

		if len(parts) == 2 {
			switch parts[0] {
			case "id":
				id, err = strconv.Atoi(parts[1])

				if err != nil {
					continue
				}

			case "name":
				name = parts[1]

			case "selector":
				selector = units.ToSelector(parts[1])
			}
		}
	}

	return units.Init(id, name, selector)
}
