package parser

import "github.com/akl-infra/bot/internal/util/assert"

type KwargType int

const (
	Bool KwargType = iota
	List
	Str
)

type Kwarg struct {
	_type KwargType
	bool  *bool
	list  *[]string
	str   *string
}

func NewBoolValue(b bool) Kwarg {
	return Kwarg{
		_type: Bool,
		bool:  &b,
	}
}

func NewListValue(l []string) Kwarg {
	return Kwarg{
		_type: List,
		list:  &l,
	}
}

func NewStrValue(s string) Kwarg {
	return Kwarg{
		_type: Str,
		str:   &s,
	}
}

func (k Kwarg) AsBool() bool {
	assert.Eq(k._type, Bool)
	return *k.bool
}

func (k Kwarg) AsList() []string {
	assert.Eq(k._type, List)
	return *k.list
}

func (k Kwarg) AsStr() string {
	assert.Eq(k._type, Str)
	return *k.str
}
