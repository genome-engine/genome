package units

type (
	Package struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		Main       bool
		Path       string
	}
	Import struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		Value      string
	}
	Structure struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		Fields     []StructField
	}
	Interface struct {
		Selector      Selector
		Comment       string
		IsExported    bool
		Name          string
		ID            int
		IsStructField bool
	}
	Custom struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		Type       string
	}
	Method struct {
		Selector            Selector
		Comment             string
		IsExported          bool
		Name                string
		ID                  int
		InInterfaceDecl     bool
		Signature           string
		Parameters, Returns []Parameter
		Body                string
	}
	Function struct {
		Selector            Selector
		Comment             string
		IsExported          bool
		Name                string
		ID                  int
		Signature           string
		Parameters, Returns []Parameter
		FuncBody            string
	}
	Constant struct {
		Selector   Selector
		Comment    string
		Name       string
		Type       string
		IsExported bool
		Enum       bool
		ID         int
	}
	Variable struct {
		Selector   Selector
		IsExported bool
		Comment    string
		Name       string
		Type       string
		ID         int
	}
	Unknown struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		Type       string
	}

	StructField struct {
		Comment    string
		Tag        string
		IsExported bool
		Name       string
		Type       string
	}
	Parameter struct {
		Name string
		Type string
	}
)
