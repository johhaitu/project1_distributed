package main

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go Phil(ch1, ch2)
	go Fork(ch2, ch1)
	// for {
	// }

}
