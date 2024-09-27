package main

import (
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
)

func main() {
	n := rsa.NewGenerateKey()
	err := n.SavePem("../../pem")
	if err != nil {
		panic(err)
	}
	privateKeyStr, publicKeyStr := n.GetKey()
	fmt.Println("privateKeyStr=", privateKeyStr)
	fmt.Println("publicKeyStr=", publicKeyStr)
}
