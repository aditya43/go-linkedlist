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

func TestValue(t *testing.T) {
	v, err := NewValue(1)
	if err != nil {
		t.Errorf("[NewValue(1)]: unexpected error: %s", err)
	}
	s := v.String()
	if s != "1" {
		t.Errorf("[v.String()]: expecting v to be 1 but got %s", s)
	}
}

func TestUnsupportedValue(t *testing.T) {
	_, err := NewValue(3.14)
	if err == nil {
		t.Errorf("[NewValue(3.14): unexpected non-error: %s", err)
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

func TestAppendUnsupportedValue(t *testing.T) {
	xs := List{}
	err := xs.Append(3.13)
	if err == nil {
		t.Errorf("[NewValue(3.14): unexpected non-error: %s", err)
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
