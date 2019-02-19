package expression

type And struct {
	Left, Right Exp
}

func (a And) Eval(params Params) bool {
	return a.Left.Eval(params) && a.Right.Eval(params)
}
