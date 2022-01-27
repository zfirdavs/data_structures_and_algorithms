package main

import (
	"fmt"
	"strings"
)

type direction string

const (
	Forward  direction = "forward"
	Backward direction = "backward"
)

type Node struct {
	value      string
	next, prev *Node
}

type LinkedList struct {
	head, tail *Node
}

// add node to front of DLL
func (l *LinkedList) push(value string) *Node {
	n := &Node{value: value}

	if l.head == nil {
		l.head = n
		l.tail = n
		return n
	}

	n.next = l.head
	l.head.prev = n
	l.head = n
	return n
}

// add node to the end of DLL
func (l *LinkedList) append(value string) *Node {
	n := &Node{value: value}

	if l.head == nil {
		l.head = n
		l.tail = n
		return n
	}

	// find the last node
	current := l.head
	for current.next != nil {
		current = current.next
	}

	n.prev = current
	current.next = n
	l.tail = n
	return n
}

func (l *LinkedList) insertAfter(prevNode *Node, value string) {
	if prevNode == nil {
		return
	}

	if l.tail == prevNode {
		l.append(value)
		return
	}

	newNode := &Node{value: value}
	newNode.prev = prevNode
	newNode.next = prevNode.next
	prevNode.next.prev = newNode
	prevNode.next = newNode

}

func (l *LinkedList) insertBefore(currNode *Node, value string) {
	if currNode == nil {
		return
	}

	if l.head == currNode {
		l.push(value)
		return
	}

	newNode := &Node{value: value}
	newNode.next = currNode
	newNode.prev = currNode.prev
	currNode.prev.next = newNode
	currNode.prev = newNode

}

func (l *LinkedList) contains(value string) bool {
	head := l.head
	for head != nil {
		if head.value == value {
			return true
		}
		head = head.next
	}
	return false
}

func (l *LinkedList) traverse(dir direction) {
	if l.head == nil {
		return
	}

	printStr := "prev = %+v, current value = %#v, next = %+v\n"

	switch dir {
	case "forward":
		current := l.head
		for current != nil {
			fmt.Printf(printStr, current.prev, current.value, current.next)
			current = current.next
		}
	case "backward":
		current := l.tail
		for current != nil {
			fmt.Printf(printStr, current.prev, current.value, current.next)
			current = current.prev
		}
	}
}

func main() {
	l := new(LinkedList)
	l.push("3")         // 3
	elem := l.push("2") // 2 <-> 3

	l.push("1")               // 1 <-> 2 <-> 3
	l.append("4")             // 1 <-> 2 <-> 3 <-> 4
	l.append("5")             // 1 <-> 2 <-> 3 <-> 4 <-> 5
	l.insertAfter(elem, "7")  // 1 <-> 2 <-> 7 <-> 3 <-> 4 <-> 5
	l.insertBefore(elem, "0") // 1 <-> 0 <-> 2 <-> 7 <-> 3 <-> 4 <-> 5
	l.traverse(Backward)
	fmt.Println(strings.Repeat("-", 150))
	l.traverse(Forward)

	l.contains("10") // false
	l.contains("7")  // true
}
