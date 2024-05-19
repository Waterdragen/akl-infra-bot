package util

import (
	"testing"

	"github.com/akl-infra/bot/internal/util/assert"
	"github.com/akl-infra/bot/internal/util/parser"
)

func TestKwarg(t *testing.T) {
	cmdKwargs := map[string]parser.KwargType{
		"mylist": parser.List,
		"mybool": parser.Bool,
		"mystr":  parser.Str,
	}
	args, err := parser.ParseKwargs("", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "")
	assert.Ok(err)

	args, err = parser.ParseKwargs("a list", parser.List, cmdKwargs)
	assert.Eq(args["args"].AsList(), []string{"a", "list"})

	args, err = parser.ParseKwargs("hello mylist --mylist 1 2 3", parser.List, cmdKwargs)
	assert.Eq(args["args"].AsList(), []string{"hello", "mylist"})
	assert.Eq(args["mylist"].AsList(), []string{"1", "2", "3"})

	args, err = parser.ParseKwargs("hello str --mystr bogos binted", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "hello str")
	assert.Eq(args["mystr"].AsStr(), "bogos binted")

	args, err = parser.ParseKwargs("hello bool --mybool", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "hello bool")
	assert.Eq(args["mybool"].AsBool(), true)

	args, err = parser.ParseKwargs("hello all --mylist a b --mystr c d --mybool", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "hello all")
	assert.Eq(args["mylist"].AsList(), []string{"a", "b"})
	assert.Eq(args["mystr"].AsStr(), "c d")
	assert.Eq(args["mybool"].AsBool(), true)

	args, err = parser.ParseKwargs("hello none --invalid --flag", parser.Str, cmdKwargs)
	assert.Fail(err)
	assert.Eq(err.Error(), "invalid kwarg: `invalid`")
	assert.Eq(args, make(map[string]parser.Kwarg))

	args, err = parser.ParseKwargs("--mybool", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "")
	assert.Eq(args["mybool"].AsBool(), true)
	assert.Eq(args["mylist"].AsList(), []string{})
	assert.Eq(args["mystr"].AsStr(), "")

	args, err = parser.ParseKwargs("many           whitespaces      --mylist    many      whitespaces", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "many whitespaces")
	assert.Eq(args["mylist"].AsList(), []string{"many", "whitespaces"})

	args, err = parser.ParseKwargs("--mylist former is original --mylist latter is duplicate", parser.Str, cmdKwargs)
	assert.Eq(args["mylist"].AsList(), []string{"latter", "is", "duplicate"})

	args, err = parser.ParseKwargs("--MYSTR UPPERCASE", parser.Str, cmdKwargs)
	assert.Eq(args["mystr"].AsStr(), "UPPERCASE")

	args, err = parser.ParseKwargs("--", parser.Str, cmdKwargs)
	assert.Fail(err)

	args, err = parser.ParseKwargs("—mystr em dash ––mylist en dash", parser.Str, cmdKwargs)
	assert.Eq(args["mystr"].AsStr(), "em dash")
	assert.Eq(args["mylist"].AsList(), []string{"en", "dash"})

	args, err = parser.ParseKwargs("--mybool this text is ignored", parser.Str, cmdKwargs)
	assert.Eq(args["args"].AsStr(), "")
	assert.Eq(args["mybool"].AsBool(), true)
	assert.Eq(args["mystr"].AsStr(), "")
	assert.Eq(args["mylist"].AsList(), []string{})

}
