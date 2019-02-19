package expression

type OptType int

const (
	CONTAINS OptType = iota
	AND
	OR
)

type Exp interface {
	Eval(params Params) bool
}

type Params interface {
	Get(key string) string
}

type Opt struct {
	Left, Right *Opt
	Operator    OptType
	Value       string
}

func (o *Opt) compare(dst *Opt) bool {
	return o.Operator >= dst.Operator
}

func AddOpt(current, exp *Opt) *Opt {
	if current == nil {
		return exp
	}
	if exp.compare(current) {
		temp := current
		current = exp
		current.Left = temp
		return current
	}

	if !exp.compare(current) && current.Left == nil {
		current.Left = exp
		return current
	}

	if !exp.compare(current) && current.Right == nil {
		current.Right = exp
		return current
	}

	return current

}

func (o *Opt) Get() Exp {
	switch o.Operator {
	case OR:
		return Or{
			o.Left.Get(),
			o.Right.Get(),
		}
	case AND:
		return And{
			o.Left.Get(),
			o.Right.Get(),
		}
	case CONTAINS:
		return Contains{o.Value}
	}
	return False{}
}
