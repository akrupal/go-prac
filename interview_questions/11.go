// merge 2 sorted linked lists

package main

import "fmt"

type list struct {
	data int
	next *list
}

func mergeLists(l1 *list, l2 *list) {
	var l *list
	var e *list
	var temp *list
	for l1 != nil && l2 != nil {
		if l1.data > l2.data {
			temp = &list{l2.data, nil}
			l2 = l2.next
		} else {
			temp = &list{l1.data, nil}
			l1 = l1.next
		}
		if l == nil {
			l = temp
			e = l
		} else {
			e.next = temp
			e = e.next
		}
	}
	for l1 != nil {
		temp = &list{l1.data, nil}
		l1 = l1.next
		if l == nil {
			l = temp
			e = l
		} else {
			e.next = temp
			e = e.next
		}
	}
	for l2 != nil {
		temp = &list{l2.data, nil}
		l2 = l2.next
		if l == nil {
			l = temp
			e = l
		} else {
			e.next = temp
			e = e.next
		}
	}
	display(l)
}

func display(l *list) {
	for l != nil {
		fmt.Println(l.data)
		l = l.next
	}
}

// func main() {
// 	list1 := &list{1, &list{3, &list{5, nil}}}
// 	list2 := &list{2, &list{4, &list{6, nil}}}
// 	mergeLists(list1, list2)
// }
