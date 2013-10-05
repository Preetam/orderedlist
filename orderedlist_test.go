package orderedlist

import (
	"fmt"
	"testing"
)

type ComparableString string
type NotComparableString int

func (ncs NotComparableString) Compare(c Comparable) int {
	return -1
}

func (cs ComparableString) Compare(c Comparable) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ComparableString Compare() panicked")
			result = 0
		}
	}()
	if cs > c.(ComparableString) {
		result = 1
	}
	if cs < c.(ComparableString) {
		result = -1
	}

	return
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

func TestPanicRecovery(t *testing.T) {
	ol := New()

	ol.Insert(ComparableString("c"))
	ol.Insert(ComparableString("a"))
	ol.Insert(ComparableString("b"))
	ol.Insert(ComparableString("aa"))
	ol.Insert(NotComparableString(1))
	ol.Insert(ComparableString("\x05"))
	ol.Remove(ComparableString("\x05"))
	if out := fmt.Sprint(ol.GetRange(ComparableString("1"), ComparableString("b"))); out != "[1 a aa]" {
		t.Errorf("Expected `[1 a aa]`, got `%v`", out)
	}
}
