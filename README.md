# go-linkedlist

[![Build Status](https://cloud.drone.io/api/badges/prologic/go-linkedlist/status.svg)](https://cloud.drone.io/prologic/go-linkedlist)
[![CodeCov](https://codecov.io/gh/prologic/go-linkedlist/branch/master/graph/badge.svg)](https://codecov.io/gh/prologic/go-linkedlist)
[![Go Report Card](https://goreportcard.com/badge/prologic/go-linkedlist)](https://goreportcard.com/report/prologic/go-linkedlist)
[![GoDoc](https://godoc.org/github.com/prologic/go-linkedlist?status.svg)](https://godoc.org/github.com/prologic/go-linkedlist) 
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/go-linkedlist/-/badge.svg)](https://sourcegraph.com/github.com/prologic/go-linkedlist?badge)


A simple singly linked list written in [Go](https://golang.org). This is
often a common interview question and one worth practicing and understanding.

To build run `make`.

```#!bash
$ go get -u github.com/prologic/go-linkedlist
$ cd $GOPATH/github.com/prologic/go-linkedlist
$ make
```

To run the tests run `make test`

Travis CI and Codecov are also integreated.

## Basics

The basics of a singly linked list is simple A `List` is an object that
holds two pointers:

- `head` -- Points to the head node in the list.
- `tail` -- Points to the tail node in the list.

Each node in the list holds a pointer to the next node and a value:

- `next` -- A pointer to the next connecting node in the list
- `value` -- The underlying value of the node

## Operations

Most operations in a singly linked list are `O(n)` with the exception of
insertion and deletion which are both `O(1)`.

### Iteration

Iterating over a singly linked list is easy. Set your initial pointer `curr`
to the `head` of the list and repeatedly access the next node in the list by
accessing `curr.next` until `curr` is `nil`.

### Reversing

Reversing a singly linked list is also fairly easy to do as you're just
swapping the node's pointers for `next` to the previous node as you iterate
over the list (//with the initial `prev` being `nil`//). At the start of the
iteration the `tail` becomes the list's `head` and once the iteration is
complete the `head` becomes the last node pointed at `prev`.

Hope this helps someone out there :D

## License

This work is licensed under the terms of the MIT License.
