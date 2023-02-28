package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/wizardpisces/dispatch-logic/practice/net/rpc/sharedRPC"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}
func main() {
	PORT := ":1234"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	myInterface := new(MyInterface)
	rpc.Register(myInterface)
	rpc.HandleHTTP()
	// t, err := net.ResolveTCPAddr("tcp4", PORT)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// l, err := net.ListenTCP("tcp4", t)
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// c, err := l.Accept()
		// if err != nil {
		// 	continue
		// }
		// fmt.Printf("%s\n", c.RemoteAddr())
		// rpc.ServeConn(c)
		http.Serve(l, nil)
	}
}
