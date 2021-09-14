package main

import (
	"fmt"
)

func Phil(chIn, chOut chan string) {
	for {
		chOut <- "i wanna eat"
		fmt.Println("hungry")
		<-chIn
		fmt.Println("eating")
		chOut <- "im done"
	}
}
