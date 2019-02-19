package expression


type Or struct {
	Left , Right Exp
}

func (o Or) Eval(params Params) bool {
	return o.Left.Eval(params) || o.Right.Eval(params)
}
