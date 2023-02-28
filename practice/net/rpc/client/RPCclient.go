package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/wizardpisces/dispatch-logic/practice/net/rpc/sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port sting!")
		return
	}

	Address := arguments[1]
	// c, err := rpc.Dial("tcp", Address)
	c, err := rpc.DialHTTP("tcp", Address)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{A1: 3, A2: 2}
	var reply float64

	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)
	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
