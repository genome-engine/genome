package units

import (
	"github.com/genome-engine/genome/engine/types"
)

func (u *Package) GetId() int            { return u.id }
func (u *Package) GetSelector() Selector { return u.selector }
func (u *Package) GetName() string       { return u.name }
func (u *Package) GetType() types.Type   { return u.Type }
func (u *Package) SetType(types.Type)    { return } //non-settable

func (u *Import) GetId() int            { return u.id }
func (u *Import) GetSelector() Selector { return u.selector }
func (u *Import) GetName() string       { return u.name }
func (u *Import) GetType() types.Type   { return u.Type }
func (u *Import) SetType(types.Type)    { return } //non-settable

func (u *Structure) GetId() int            { return u.id }
func (u *Structure) GetSelector() Selector { return u.selector }
func (u *Structure) GetName() string       { return u.name }
func (u *Structure) GetType() types.Type   { return u.Type }
func (u *Structure) SetType(types.Type)    { return } //non-settable

func (u *Interface) GetId() int            { return u.id }
func (u *Interface) GetSelector() Selector { return u.selector }
func (u *Interface) GetName() string       { return u.name }
func (u *Interface) GetType() types.Type   { return u.Type }
func (u *Interface) SetType(types.Type)    { return } //non-settable

func (u *Custom) GetId() int            { return u.id }
func (u *Custom) GetSelector() Selector { return u.selector }
func (u *Custom) GetName() string       { return u.name }
func (u *Custom) GetType() types.Type   { return u.Type }
func (u *Custom) SetType(t types.Type)  { u.Type = t }

func (u *Method) GetId() int            { return u.id }
func (u *Method) GetSelector() Selector { return u.selector }
func (u *Method) GetName() string       { return u.name }
func (u *Method) GetType() types.Type   { return u.Type }
func (u *Method) SetType(types.Type)    { return } //non-settable

func (u *Function) GetId() int            { return u.id }
func (u *Function) GetSelector() Selector { return u.selector }
func (u *Function) GetName() string       { return u.name }
func (u *Function) GetType() types.Type   { return u.Type }
func (u *Function) SetType(types.Type)    { return } //non-settable

func (u *Constant) GetId() int            { return u.id }
func (u *Constant) GetSelector() Selector { return u.selector }
func (u *Constant) GetName() string       { return u.name }
func (u *Constant) GetType() types.Type   { return u.Type }
func (u *Constant) SetType(t types.Type)  { u.Type = t }

func (u *Variable) GetId() int            { return u.id }
func (u *Variable) GetSelector() Selector { return u.selector }
func (u *Variable) GetName() string       { return u.name }
func (u *Variable) GetType() types.Type   { return u.Type }
func (u *Variable) SetType(t types.Type)  { u.Type = t }

func (u *Unknown) GetId() int            { return u.id }
func (u *Unknown) GetName() string       { return u.name }
func (u *Unknown) GetSelector() Selector { return u.selector }
func (u *Unknown) GetType() types.Type   { return u.Type }
func (u *Unknown) SetType(t types.Type)  { u.Type = t }

func (u *EnumSeries) GetId() int            { return u.id }
func (u *EnumSeries) GetName() string       { return u.name }
func (u *EnumSeries) GetSelector() Selector { return u.selector }
func (u *EnumSeries) GetType() types.Type   { return u.Type }
func (u *EnumSeries) SetType(t types.Type)  { u.Type = t }
func (u *EnumSeries) SetName(name string)   { u.name = name } //non-interface
func (u *EnumSeries) SetId(id int)          { u.id = id }     //non-interface
