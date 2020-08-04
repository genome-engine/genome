package units

import (
	"github.com/genome-engine/genome/engine/types"
)

func (u *Package) GetId() int            { return u.ID }
func (u *Package) GetSelector() Selector { return u.Selector }
func (u *Package) GetName() string       { return u.Name }
func (u *Package) GetType() types.Type   { return u.Type }
func (u *Package) SetType(types.Type)    { return } //non-settable

func (u *Import) GetId() int            { return u.ID }
func (u *Import) GetSelector() Selector { return u.Selector }
func (u *Import) GetName() string       { return u.Name }
func (u *Import) GetType() types.Type   { return u.Type }
func (u *Import) SetType(types.Type)    { return } //non-settable

func (u *Structure) GetId() int            { return u.ID }
func (u *Structure) GetSelector() Selector { return u.Selector }
func (u *Structure) GetName() string       { return u.Name }
func (u *Structure) GetType() types.Type   { return u.Type }
func (u *Structure) SetType(types.Type)    { return } //non-settable

func (u *Interface) GetId() int            { return u.ID }
func (u *Interface) GetSelector() Selector { return u.Selector }
func (u *Interface) GetName() string       { return u.Name }
func (u *Interface) GetType() types.Type   { return u.Type }
func (u *Interface) SetType(types.Type)    { return } //non-settable

func (u *Custom) GetId() int            { return u.ID }
func (u *Custom) GetSelector() Selector { return u.Selector }
func (u *Custom) GetName() string       { return u.Name }
func (u *Custom) GetType() types.Type   { return u.Type }
func (u *Custom) SetType(t types.Type)  { u.Type = t }

func (u *Method) GetId() int            { return u.ID }
func (u *Method) GetSelector() Selector { return u.Selector }
func (u *Method) GetName() string       { return u.Name }
func (u *Method) GetType() types.Type   { return u.Type }
func (u *Method) SetType(types.Type)    { return } //non-settable

func (u *Function) GetId() int            { return u.ID }
func (u *Function) GetSelector() Selector { return u.Selector }
func (u *Function) GetName() string       { return u.Name }
func (u *Function) GetType() types.Type   { return u.Type }
func (u *Function) SetType(types.Type)    { return } //non-settable

func (u *Constant) GetId() int            { return u.ID }
func (u *Constant) GetSelector() Selector { return u.Selector }
func (u *Constant) GetName() string       { return u.Name }
func (u *Constant) GetType() types.Type   { return u.Type }
func (u *Constant) SetType(t types.Type)  { u.Type = t }

func (u *Variable) GetId() int            { return u.ID }
func (u *Variable) GetSelector() Selector { return u.Selector }
func (u *Variable) GetName() string       { return u.Name }
func (u *Variable) GetType() types.Type   { return u.Type }
func (u *Variable) SetType(t types.Type)  { u.Type = t }

func (u *Unknown) GetId() int            { return u.ID }
func (u *Unknown) GetName() string       { return u.Name }
func (u *Unknown) GetSelector() Selector { return u.Selector }
func (u *Unknown) GetType() types.Type   { return u.Type }
func (u *Unknown) SetType(t types.Type)  { u.Type = t }
