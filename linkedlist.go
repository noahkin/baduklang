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
	for iter != nil {
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
		data: "data",
		next: nil,
	}

	slice := make([]string, 1)
	slice[0] = "data"

	var array [10]string
	array[0] = "data"

	list_chan := make(chan string)
	slice_chan := make(chan string)
	array_chan := make(chan string)

	go func() {
		for i := 1; i < 10; i++ {
			push(&list, "data")
		    list_chan <- "list"
		}
	}()
	go func() {
		for i := 1; i < 10; i++ {
			slice = append(slice, "data")
		    slice_chan <- "slice"
		}
	}()
	go func() {
		for i := 1; i < 10; i++ {
			array[i] = "data"
		    array_chan <- "array"
		}
	}()

	for i := 0; i < 27; i++ {
		select {
		case list_msg := <-list_chan:
			fmt.Println(list_msg)
		case slice_msg := <-slice_chan:
			fmt.Println(slice_msg)
		case array_msg := <-array_chan:
			fmt.Println(array_msg)
		}
	}
}
