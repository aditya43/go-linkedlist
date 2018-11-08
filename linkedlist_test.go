package linkedlist

import (
	"testing"
)

func TestEmptyList(t *testing.T) {
	xs := List{}
	s := xs.String()
	if s != "[]" {
		t.Errorf("expected empty list to be [] but got %s", s)
	}
}

func TestAppend(t *testing.T) {
	xs := List{}
	xs.Append(1)
	xs.Append(2)
	xs.Append(3)
	s := xs.String()
	if s != "[3, 2, 1]" {
		t.Errorf("expected list to be [3, 2, 1] but got %s", s)
	}
}

func TestReverse(t *testing.T) {
	xs := List{}
	xs.Append(1)
	xs.Append(2)
	xs.Append(3)
	xs.Reverse()
	s := xs.String()
	if s != "[1, 2, 3]" {
		t.Errorf("expected list to be [1, 2, 3] but got %s", s)
	}
}
