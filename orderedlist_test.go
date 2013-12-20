package orderedlist

import (
	"fmt"
	"testing"
)

type ComparableString string

func CompareStrings(a, b interface{}) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
		}
	}()

	aStr := a.(string)
	bStr := b.(string)

	if aStr > bStr {
		result = 1
	}

	if aStr < bStr {
		return -1
	}

	return
}

func TestSetGet(t *testing.T) {
	ol := New(CompareStrings)

	ol.Insert("c")
	ol.Insert("a")
	ol.Insert("b")
	ol.Insert("aa")
	ol.Insert("1")
	ol.Insert("\x05")
	ol.Remove("\x05")
	if out := fmt.Sprint(ol.GetRange("", "\xff")); out != "[1 a aa b c]" {
		t.Errorf("Expected `[1 a aa b c]`, got `%v`", out)
	}
}

func TestGetRange(t *testing.T) {
	ol := New(CompareStrings)

	ol.Insert("c")
	ol.Insert("a")
	ol.Insert("b")
	ol.Insert("aa")
	ol.Insert("1")
	ol.Insert("\x05")
	ol.Remove("\x05")
	if out := fmt.Sprint(ol.GetRange("1", "b")); out != "[1 a aa]" {
		t.Errorf("Expected `[1 a aa]`, got `%v`", out)
	}
}

func TestGetRangeIterator(t *testing.T) {
	ol := New(CompareStrings)

	ol.Insert("c")
	ol.Insert("a")
	ol.Insert("b")
	ol.Insert("aa")
	ol.Insert("1")
	ol.Insert("\x05")
	ol.Remove("\x05")

	i := ol.GetRangeIterator("b", "\xff")

	if e := i.Next(); e != nil {
		if CompareStrings(e.Value(), "c") != 0 {
			t.Errorf("Expected iterator value to be c, got %v", e.Value())
		}
	}
	if i == nil {
		t.Error(i)
	} else if e := i.Prev(); e != nil {
		t.Errorf("Expected iterator to be nil, got non-nil with value %v", e.Value())
	}
}
