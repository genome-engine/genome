package scenario

//A script is a strict sequence of parsing,
//execute intermediate commands and generate code.
type Scenario struct {
	Chain
}

type Chain map[Type]Step
