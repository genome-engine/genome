package units

type (
	//The Unit interface is used to universalize elements of the collection.
	//This data will then have to go through type conversion to basic structures.
	Unit interface {
		GetId() int
		GetName() string
		GetSelector() Selector
	}

	//The Selector is created for identification of an accessory of default_template object
	//and also for creation of possible variants of nesting of selectors in each other.
	//More details about selectors and units are written in ./units/README.md,
	//i.e. in the basic implementation of units and selectors.
	Selector interface {
		Name() string
		CanContain(selector Selector) bool
		ChildSelectors() []Selector
		ParentSelectors() []Selector
	}
)

type (
	Pack struct {
		Selector   Selector
		Comment    string
		Name       string
		Path       string
		IsExported bool
		Main       bool
		ID         int
	}
	Import struct {
		Selector   Selector
		Comment    string
		Name       string
		Value      string
		IsExported bool
		ID         int
	}
	Struct struct {
		Selector   Selector
		Comment    string
		Name       string
		IsExported bool
		ID         int
		Fields     []StructField
	}
	Iface struct {
		Selector      Selector
		Comment       string
		Name          string
		IsExported    bool
		IsStructField bool
		ID            int
	}
	Custom struct {
		Selector   Selector
		Comment    string
		Name       string
		Type       string
		IsExported bool
		ID         int
	}
	Method struct {
		Selector            Selector
		Comment             string
		Name                string
		Signature           string
		Body                string
		Type                string
		IsExported          bool
		InInterfaceDecl     bool
		ID                  int
		Parameters, Returns []Param
	}
	Func struct {
		Selector            Selector
		Comment             string
		Name                string
		Signature           string
		Type                string
		FuncBody            string
		IsExported          bool
		ID                  int
		Parameters, Returns []Param
	}
	Const struct {
		Selector   Selector
		Comment    string
		Name       string
		Type       string
		IsExported bool
		Enum       bool
		ID         int
	}
	Var struct {
		Selector   Selector
		Comment    string
		Name       string
		Type       string
		IsExported bool
		ID         int
	}
	Unknown struct {
		Selector   Selector
		Comment    string
		Name       string
		Type       string
		IsExported bool
		ID         int
	}

	StructField struct {
		Comment    string
		Tag        string
		Name       string
		Type       string
		IsExported bool
	}
	Param struct {
		Name string
		Type string
	}
)
