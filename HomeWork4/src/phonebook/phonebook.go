package main

import (
	"fmt"
	"sort"
	"strings"
)

type recordPB struct {
	name   string
	phones []int
}

type phonebookSortableT []recordPB

func (pB phonebookSortableT) String() string {
	rezult := ""
	for num, rec := range pB {
		rezult = rezult + " " + fmt.Sprintln(num, "Абонент:", rec.name)
		for i, number := range rec.phones {
			rezult = rezult + " " + fmt.Sprintf("\t %v) %v \n", i+1, number)
		}
		rezult = rezult + "\n"
	}
	return "myPhoneBook\n " + rezult
}

// Len возвращает количество элементов в коллекции.
func (pB phonebookSortableT) Len() int {
	return len(pB)
}

// Less отдает true, если элементы i, j  необходимо поменять местами.
func (pB phonebookSortableT) Less(i, j int) bool {
	return strings.Compare(pB[i].name, pB[j].name) == -1
}

// Swap меняет местами элементы с индексами i,j.
func (pB phonebookSortableT) Swap(i, j int) {
	fmt.Println("Swap", i, j)
	pB[i], pB[j] = pB[j], pB[i]
}

func main() {
	aBook := &phonebookSortableT{}
	*aBook = append(*aBook, recordPB{name: "Ivan", phones: []int{22223322, 28223344}})
	*aBook = append(*aBook, recordPB{name: "Andrew", phones: []int{22223323, 27223344}})
	*aBook = append(*aBook, recordPB{name: "Onotole", phones: []int{22223326, 24223344}})
	*aBook = append(*aBook, recordPB{name: "Zeus", phones: []int{22223365, 22223344}})
	*aBook = append(*aBook, recordPB{name: "Xen", phones: []int{22223322, 21223444}})

	fmt.Println(aBook)
	sort.Sort(aBook)
	fmt.Println(aBook)

}
