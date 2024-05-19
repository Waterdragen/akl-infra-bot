package assert

import (
	"log"
	"reflect"
)

func Eq(lhs, rhs interface{}) {
	if !reflect.DeepEqual(lhs, rhs) {
		log.Panicf("Assertion failed:\nLeft != Right\nlhs = `%v`\nrhs = `%v`", lhs, rhs)
	}
}

func Ne(lhs, rhs interface{}) {
	if reflect.DeepEqual(lhs, rhs) {
		log.Panicf("Assertion failed:\nLeft == Right\nlhs = `%v`\nrhs = `%v`", lhs, rhs)
	}
}

func Ok(err error) {
	if err != nil {
		log.Panicf("Assertion failed:\nError: `%v`", err)
	}
}

func Fail(err error) {
	if err == nil {
		log.Panicf("Assertion failed:\nError: `%v` (ok)", err)
	}
}
