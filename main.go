package main

import (
	"ecommerce/util"
	"fmt"
)

// "ecommerce/cmd"

func main() {
	// cmd.Serve()

	jwt, err := util.CreateJwt("my-sercet", util.Payload{
		Sub: 47,
		FirstName: "Rifat",
		LastName: "Kobir",
		Email:     "rifat@test.com",
		IsSopOwner: false,
	})

	if err != nil{
		fmt.Println("This is the error", err)
		return
	}

	fmt.Println(jwt)
}
