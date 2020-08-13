package units

func (u *Package) GetId() int            { return u.ID }
func (u *Package) GetSelector() Selector { return u.Selector }
func (u *Package) GetName() string       { return u.Name }

func (u *Import) GetId() int            { return u.ID }
func (u *Import) GetSelector() Selector { return u.Selector }
func (u *Import) GetName() string       { return u.Name }

func (u *Structure) GetId() int            { return u.ID }
func (u *Structure) GetSelector() Selector { return u.Selector }
func (u *Structure) GetName() string       { return u.Name }

func (u *Interface) GetId() int            { return u.ID }
func (u *Interface) GetSelector() Selector { return u.Selector }
func (u *Interface) GetName() string       { return u.Name }

func (u *Custom) GetId() int            { return u.ID }
func (u *Custom) GetSelector() Selector { return u.Selector }
func (u *Custom) GetName() string       { return u.Name }

func (u *Method) GetId() int            { return u.ID }
func (u *Method) GetSelector() Selector { return u.Selector }
func (u *Method) GetName() string       { return u.Name }

func (u *Function) GetId() int            { return u.ID }
func (u *Function) GetSelector() Selector { return u.Selector }
func (u *Function) GetName() string       { return u.Name }

func (u *Constant) GetId() int            { return u.ID }
func (u *Constant) GetSelector() Selector { return u.Selector }
func (u *Constant) GetName() string       { return u.Name }

func (u *Variable) GetId() int            { return u.ID }
func (u *Variable) GetSelector() Selector { return u.Selector }
func (u *Variable) GetName() string       { return u.Name }

func (u *Unknown) GetId() int            { return u.ID }
func (u *Unknown) GetName() string       { return u.Name }
func (u *Unknown) GetSelector() Selector { return u.Selector }
