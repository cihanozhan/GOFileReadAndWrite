package main 

import (
    "fmt"
    "io/ioutil"
)

func main(){

    fileName := "./dummy.txt"

    content, err := ioutil.ReadFile(fileName)
    checkError(err)

    result := string(content)

    fmt.Printf("Okunan Veri: %v", result)

}

func checkError(err error){
    if err != nil{
        panic(err)
    }
}