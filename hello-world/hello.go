package main

import (
	"fmt"
	"github.com/alexwally/my-go-app/common"
)

func main() {
	message := common.Hello("Gladys")
	
	fmt.Println(message)
}
