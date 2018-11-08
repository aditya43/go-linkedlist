package main

import (
	"bytes"
	"fmt"
)

// Value represents a value object and must implement thr Stringer interface
type Value interface {
	String() string
}

// NewValue constructs a new Value object based on the type of the underlying
// value `v` passed. Unsupported types results in an error.
func NewValue(v interface{}) (Value, error) {
	switch t := v.(type) {
	case int:
		return &IntValue{value: v.(int)}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %q", t)
	}
}

// IntValue is a Value that holds an int
type IntValue struct {
	value int
}

// String returns a stringified version of the underlying value
func (v *IntValue) String() string {
	return fmt.Sprintf("%d", v.value)
}

// Node is an object that holds the structure of a singlely linked list
// and supporting methods for read and write operations
type Node struct {
	/* This adds support for doubly linked list
	prev  *Node
	*/

	next  *Node
	value Value
}

// String returns the value as a string
func (n *Node) String() string {
	return n.value.String()
}

// List represents a linked list of connected nodes
type List struct {
	head *Node
	tail *Node
}

// Append appends the value to the list creating a new appropriate value
// object and node and linking the new node to the tail of the list
func (l *List) Append(v interface{}) error {
	value, err := NewValue(v)
	if err != nil {
		return err
	}

	node := &Node{next: l.head, value: value}
	/* This adds support for doubly linked list
	if l.head != nil {
		l.head.prev = node
	}
	*/

	l.head = node

	node = l.head
	for node.next != nil {
		node = node.next
	}
	l.tail = node

	return nil
}

// Reverse reverses the order of the nodes in the list
func (l *List) Reverse() {
	var (
		curr *Node
		prev *Node
	)

	curr = l.head
	l.tail = l.head

	for curr != nil {
		curr.next, curr, prev = prev, curr.next, curr
	}

	l.head = prev
}

// String returns a stringified represenation of the list
func (l *List) String() string {
	var buf bytes.Buffer

	node := l.head
	buf.WriteString("[")
	for node != nil {
		buf.WriteString(node.value.String())
		node = node.next
		if node != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}

func main() {
	xs := &List{}
	fmt.Printf("xs = %s\n", xs.String())

	fmt.Println("appending 3 elements to xs ...")
	xs.Append(1)
	xs.Append(2)
	xs.Append(3)

	fmt.Printf("xs[head] = %s\n", xs.head.String())
	fmt.Printf("xs[tail] = %s\n", xs.tail.String())
	fmt.Printf("xs = %s\n", xs.String())

	fmt.Println("reverse the list ...")
	xs.Reverse()
	fmt.Printf("xs = %s\n", xs.String())
}
