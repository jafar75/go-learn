package utils

import (
	"fmt"
)


type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(value int) {
	// newNode := &Node{data: value};	// eqauls to below
	//var newNode *Node = &Node{data: value};   // equals to below

	newNode := new(Node);
	newNode.Data = value;


	if l.Head == nil {
		l.Head = newNode
		return;
	}
	
	curr := l.Head;
	for curr.Next != nil {
		curr = curr.Next;
	}
	curr.Next = newNode;
}

func (l *LinkedList) AddByNode(node *Node) {
	if l.Head == nil {
		l.Head = node
		return;
	}
	
	curr := l.Head;
	for curr.Next != nil {
		curr = curr.Next;
	}
	curr.Next = node;
}

func (l* LinkedList) Remove(value int) {
	if l.Head == nil {
		return;
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next;	// just shift the head 1
	}

	curr := l.Head;
	for curr.Next != nil && curr.Next.Data != value {
		curr = curr.Next;
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next;
	}
}

func SliceToLinkedList(slc []int) *LinkedList {
	list := &LinkedList{};
	for _, v := range slc {
		list.Add(v);
	}
	return list;
}

func PrintList(l *LinkedList) {
    curr := l.Head
    for curr != nil {
        fmt.Printf("%d ", curr.Data)
        curr = curr.Next
    }
    fmt.Println()
}