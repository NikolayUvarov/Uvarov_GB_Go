package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main_() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Alex2"] = []int{89996543211}
	addressBook["Alex4"] = []int{89996543212}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)
	addressBook["Bob"] = append(addressBook["Bob"], 89155243623)

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}

	b, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Errorf("Cannot marshal data")
	}
	fmt.Println("Marsalled data:\n", b)

	fmt.Println("Delete Alex")
	delete(addressBook, "Alex")
	b, err = json.Marshal(addressBook)
	if err != nil {
		fmt.Errorf("Cannot marshal data")
	}
	fmt.Println("Marsalled data:\n", b)

	fname := "phonebook.dat"
	//Save data
	err = ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		fmt.Errorf("Cannot save data")
	}

	//Read it again

	readB, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Errorf("Cannot read data")
	}
	fmt.Println("Read data:\n", readB)

	fmt.Println("Saved and read data:")
	aB := make(map[string][]int)

	err = json.Unmarshal(readB, &aB)
	if err != nil {
		fmt.Errorf("Cannot unmarshall data")
	}
	for name, numbers := range aB {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
	fmt.Println("Alex2", aB["Alex2"])
}
