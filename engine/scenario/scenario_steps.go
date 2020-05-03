package scenario

type Step interface {
	Execute() error
}
