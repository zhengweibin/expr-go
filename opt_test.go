package expression

import (
	"testing"
	"util"
)

func TestOptEval(t *testing.T) {
	opts := []*Opt{
		{Operator: CONTAINS, Value: "john"},
		{Operator: OR},
		{Operator: CONTAINS, Value: "smith"},
		{Operator: OR},
		{Operator: CONTAINS, Value: "marry"},
	}
	var current *Opt
	for _, opt := range opts {
		current = AddOpt(current, opt)
	}
	line := Line("hello world ")
	exp := current.Get()
	util.Infof("%v", exp.Eval(line))
}
