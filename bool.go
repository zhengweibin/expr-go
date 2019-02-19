package expression

type True struct {
}

func (t True) Eval(params Params) bool {
	return true
}

type False struct {
}

func (f False) Eval(params Params) bool {
	return false
}
