package main

import (
    "fmt"
    "io/ioutil"
)

const (
    DECLARATION_OF_INDEPENDENCE string = "data/declaration_of_independence.txt"
    CONSTITUTION                string = "data/constitution.txt"
)

func main() {
    document, err := ioutil.ReadFile(CONSTITUTION)
    if err != nil {
        fmt.Println("Err is ", err)
    }
    strBuffer := string(document)
    fmt.Println(strBuffer)
}
