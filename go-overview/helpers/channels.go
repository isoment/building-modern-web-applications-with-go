package helpers

import (
	"fmt"
	"math/rand"
)

const numPool = 10

func randomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

func calculateValue(intChan chan int) {
	randomNumber := randomNumber(numPool)
	intChan <- randomNumber
}

/*
The defer keyword simply says that we will execute whatever comes after the keyword
as soon as the current function is done. In this case we want to ensure that the intChan
is closed when this function is done running.
*/
func Channels() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}
