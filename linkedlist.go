package main

import "fmt"

type node struct {
	data string
	next *node
}

func push(list *node, new_data string) {
	tail := list
	for tail.next != nil {
		tail = tail.next
	}
	new_node := node{
		data: new_data,
		next: nil,
	}
	tail.next = &new_node
}

func print_list(list *node) {
	iter := list
	for iter.next != nil {
		fmt.Println(iter.data)
		iter = iter.next
	}
}

func delete_match(head *node, match string) {
	if head.data == match {
		head = head.next
		// this doesn't work
	}
	iter := head
	for iter.next != nil {
		if iter.next.data == match {
			iter.next = iter.next.next
		} else {
			iter = iter.next
		}
	}
}

func main() {
	list := node{
		data: "First",
		next: nil,
	}
	push(&list, "Second")
	push(&list, "Third")
	push(&list, "Fourth")
	push(&list, "Fourth")
	push(&list, "Fifth")
	push(&list, "Sixth")
	push(&list, "Seventh")
	print_list(&list)
	delete_match(&list, "First")
	delete_match(&list, "Fourth")
	print_list(&list)
}
