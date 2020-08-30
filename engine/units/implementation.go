package units

func (u *Pack) GetId() int            { return u.ID }
func (u *Pack) GetSelector() Selector { return u.Selector }
func (u *Pack) GetName() string       { return u.Name }

func (u *Import) GetId() int            { return u.ID }
func (u *Import) GetSelector() Selector { return u.Selector }
func (u *Import) GetName() string       { return u.Name }

func (u *Struct) GetId() int            { return u.ID }
func (u *Struct) GetSelector() Selector { return u.Selector }
func (u *Struct) GetName() string       { return u.Name }

func (u *Iface) GetId() int            { return u.ID }
func (u *Iface) GetSelector() Selector { return u.Selector }
func (u *Iface) GetName() string       { return u.Name }

func (u *Custom) GetId() int            { return u.ID }
func (u *Custom) GetSelector() Selector { return u.Selector }
func (u *Custom) GetName() string       { return u.Name }

func (u *Method) GetId() int            { return u.ID }
func (u *Method) GetSelector() Selector { return u.Selector }
func (u *Method) GetName() string       { return u.Name }

func (u *Func) GetId() int            { return u.ID }
func (u *Func) GetSelector() Selector { return u.Selector }
func (u *Func) GetName() string       { return u.Name }

func (u *Const) GetId() int            { return u.ID }
func (u *Const) GetSelector() Selector { return u.Selector }
func (u *Const) GetName() string       { return u.Name }

func (u *Var) GetId() int            { return u.ID }
func (u *Var) GetSelector() Selector { return u.Selector }
func (u *Var) GetName() string       { return u.Name }

func (u *Unknown) GetId() int            { return u.ID }
func (u *Unknown) GetName() string       { return u.Name }
func (u *Unknown) GetSelector() Selector { return u.Selector }
