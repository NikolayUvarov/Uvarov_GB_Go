/*

create queue
1) read dir by 10 elements and put it in queue
2) process element from queue
3)




*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	tstMyQ()

}

func tstMyQ() {
	q := list.New()

	fmt.Println("q len:", q.Len())
	q.PushBack("one")
	q.PushBack("two")
	fmt.Println("q len:", q.Len())

	q.PushBack("two")
	q.PushBack("tree")
	q.PushBack("tree")
	q.PushBack("tree")
	q.PushBack("four")
	fmt.Println("q len:", q.Len())

	for q.Len() > 0 {
		e := q.Front()
		fmt.Println("dequeue q len:", q.Len(), e.Value)
		q.Remove(e)
	}
}
