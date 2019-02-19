package expression

import "strings"

type Contains struct {
	keyword string
}

func (c Contains) Eval(params Params) bool {
	return strings.Contains(params.Get(""), c.keyword)
}
