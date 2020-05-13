package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

type recordsT [][]string

func main() {
	inString := `key,val1,val2
"key1","val11", val12
"key2","val21", val22
"key3","val31", val32
"key4","val41", val42
"key5","val51", val52
`

	fmt.Println(readCsv(strings.NewReader(inString)))

}

func readCsv(sr io.Reader) recordsT {
	var records recordsT
	r := csv.NewReader(sr)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
