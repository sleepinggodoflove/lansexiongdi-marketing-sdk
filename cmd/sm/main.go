package main

import (
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
)

func main() {
	priK, pubK, err := sm.GenerateKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(priK)
	fmt.Println(pubK)
	sm4key, err := sm.GenerateSM4Key()
	if err != nil {
		panic(err)
	}
	fmt.Println(sm4key)
}
