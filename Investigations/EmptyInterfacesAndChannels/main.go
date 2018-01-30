package main

import (
	"fmt"
	"reflect"
	"time"
)

func emptyInterface(i interface{}) string {
	value := reflect.ValueOf(i)

	switch value.Kind() {
	case reflect.Int:
		return fmt.Sprintf("Field Type: %s \t Value is: %d", value.Type(), value.Int())
	case reflect.Bool:
		return fmt.Sprintf("Field Type: %s \t Value is: %v", value.Type(), value.Bool())
	default:
		return ""
	}

}

func channelFUnc(referenceToMainBridge chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		receivedInt := <-referenceToMainBridge // the recieved value is not used, its just disposed
		//main is blocked till its recieved here
		fmt.Printf("In the Go routine %d \n", receivedInt)
	}

}

func main() {
	var i = 5
	fmt.Println(emptyInterface(i))

	var b = true
	fmt.Println(emptyInterface(b))

	bridge := make(chan int) //creating the channel
	go channelFUnc(bridge)   //creating a new go routine and sharing the reference of the channel
	for i := 0; i < 10; i++ {
		bridge <- i //sending integer through the bridge(channel), this is blocked till its recieved in the goroutine
		fmt.Printf("in the main function %d \n", i)

	}
	time.Sleep(1 * time.Second)

}
