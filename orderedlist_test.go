package orderedlist

import (
	"fmt"
	"testing"
)

type ComparableString string

func (cs ComparableString) Compare(c Comparable) int {
	if cs > c.(ComparableString) {
		return 1
	}
	if cs < c.(ComparableString) {
		return -1
	}
	return 0
}

func TestSetGet(t *testing.T) {
	ol := New()

	ol.Insert(ComparableString("c"))
	ol.Insert(ComparableString("a"))
	ol.Insert(ComparableString("b"))
	ol.Insert(ComparableString("aa"))
	ol.Insert(ComparableString("1"))
	ol.Insert(ComparableString("\x05"))
	ol.Remove(ComparableString("\x05"))
	if out := fmt.Sprint(ol.GetRange(ComparableString(""), ComparableString("\xff"))); out != "[1 a aa b c]" {
		t.Errorf("Expected `[1 a aa b c]`, got `%v`", out)
	}
}

func TestGetRange(t *testing.T) {
	ol := New()

	ol.Insert(ComparableString("c"))
	ol.Insert(ComparableString("a"))
	ol.Insert(ComparableString("b"))
	ol.Insert(ComparableString("aa"))
	ol.Insert(ComparableString("1"))
	ol.Insert(ComparableString("\x05"))
	ol.Remove(ComparableString("\x05"))
	if out := fmt.Sprint(ol.GetRange(ComparableString("1"), ComparableString("b"))); out != "[1 a aa]" {
		t.Errorf("Expected `[1 a aa]`, got `%v`", out)
	}
}
