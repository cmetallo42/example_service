package main

import (
	"fmt"
	"io/ioutil"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "client-124")
	if err != nil {
		panic(err)
	}
	file, err := os.Open("model.json")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
   
    data, err := ioutil.ReadAll(file)
    if err != nil {
    	panic(err)
    }
    
	err = sc.Publish("test", data)
	if err != nil {
		panic(err)
	}
	sc.Close()
}
