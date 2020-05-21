package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Println(rf("fileread.go"))
}
//������� ������ ����� � ��������� �������
//��� ������� ������ ��������� ������


func rf(name string) string{
    bs, err := ioutil.ReadFile(name)
    if err != nil {
        return ""
    }
    return(string(bs))
}
