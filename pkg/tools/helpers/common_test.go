package helpers

import (
	"crypto/md5"
	"fmt"
	"testing"
)

var testAbs = []struct {
	in     int
	answer int
}{
	{0, 0}, {-1, 1}, {1, 1},
}

var testStrToBool = []struct {
	in     string
	answer bool
}{
	{"да", true}, {"true", true}, {"false", false}, {"", false},
}

var testStrToInt = []struct {
	in     string
	answer int
}{
	{"1", 1}, {"-25", -25}, {"25.5", 0}, {"wrong string", 0},
}

var testStrToFloat = []struct {
	in     string
	answer float64
}{
	{"2", 2}, {"2.54321", 2.54321},
	{"1.797693134862315708145274237317043567981e+308", 1.797693134862315708145274237317043567981e+308},
	{"wrong string", 0.0},
	{"-125.521", -125.521},
}

var testFind = []struct {
	in     []string
	find   string
	answer int
}{
	{[]string{"string", "str"}, "str", 1}, {[]string{"wrong string"}, "hello", -1},
	{[]string{"correct string", "correct"}, "correct", 1},
	{[]string{"1", "2", "3", "4", "5"}, "5", 4},
}

var testContains = []struct {
	in     []string
	find   string
	answer bool
}{
	{[]string{"string", "str"}, "str", true}, {[]string{"wrong string"}, "hello", false},
	{[]string{"correct string", "correct"}, "correct", true},
	{[]string{"1", "2", "3", "4", "5"}, "5", true},
}

var testMd5 = []struct {
	in     string
	answer string
}{
	{"md5_test_string", fmt.Sprintf("%x", md5.Sum([]byte("md5_test_string")))},
}

func TestAbs(t *testing.T) {
	for _, testValue := range testAbs {
		res := Abs(testValue.in)
		if res != testValue.answer {
			t.Errorf("Abs(%d), expected %d", testValue.in, res)
		}
	}
}

func TestStrToBool(t *testing.T) {
	for _, testValue := range testStrToBool {
		res := StrToBool(testValue.in)
		if res != testValue.answer {
			t.Errorf("StrToBool(%s), expected %t", testValue.in, res)
		}
	}
}

func TestStrToInt(t *testing.T) {
	for _, testValue := range testStrToInt {
		res := StrToInt(testValue.in)
		if res != testValue.answer {
			t.Errorf("StrToInt(%s), excpected %v", testValue.in, res)
		}
	}
}

func TestStrToFloat(t *testing.T) {
	for _, testValue := range testStrToFloat {
		res := StrToFloat(testValue.in)
		if res != testValue.answer {
			t.Errorf("StrToFloat(%s), expected %v", testValue.in, res)
		}
	}
}

func TestFind(t *testing.T) {
	for _, testValue := range testFind {
		res := Find(testValue.in, testValue.find)
		if res != testValue.answer {
			t.Errorf("Find(%v, %s), expected %d", testValue.in, testValue.find, res)
		}
	}
}

func TestContains(t *testing.T) {
	for _, testValue := range testContains {
		res := Contains(testValue.in, testValue.find)
		if res != testValue.answer {
			t.Errorf("Contains(%v, %s), expected %t", testValue.in, testValue.find, res)
		}
	}
}

func TestMd5Hash(t *testing.T) {
	for _, testValue := range testMd5 {
		res := Md5Hash(testValue.in)
		if res != testValue.answer {
			t.Errorf("Md5Hash(%s), expected %s", testValue.in, res)
		}
	}
}
