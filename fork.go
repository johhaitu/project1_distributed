package main

import "fmt"

func Fork(chIn, chOut chan string) {
	for {
		<-chIn
		chOut <- "you can eat"
		<-chIn
		fmt.Println("im done")
	}

}
