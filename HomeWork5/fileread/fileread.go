package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Println(rf("fileread.go"))
}
//вынесем чтение файла в отдельную функцию
//так удобнее читать несколько файлов


func rf(name string) string{
    bs, err := ioutil.ReadFile(name)
    if err != nil {
        return ""
    }
    return(string(bs))
}
