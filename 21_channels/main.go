package main

import (
	"fmt"
	// "time"
)

func processNum(num chan int) {
	fmt.Println("Processing the number", <-num)
}

func sum(result chan int, a int, b int) {
	ans := a + b
	result <- ans
}

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("Processing...")
}

func emailSender(emailChan <-chan string, done chan<- bool){
		defer func() {
		done <- true
	}()
	// emailChan<-"Hello@example.com" -> this is invalid 
	for email:=range emailChan{
		fmt.Println("Sending email to: ", email)
	}
}

func main() {

	chan1:=make(chan int)
	chan2:=make(chan string)


	go func(){
		chan1<-10
	}()

	go func(){
		chan2<-"Golang"
	}()

	for i:=0;i<2;i++{
		select{
		case chan1Val:=<-chan1:
			fmt.Println("Recieved data from chan1:", chan1Val)
		case chan2Val:=<-chan2:
			fmt.Println("Recieved data from chan1:", chan2Val)
		}
	}


// 	//Buffered Channels Example:
// // this will be non blocking until the limit given: in this example we see 100
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)

// 	go emailSender(emailChan, done)

// 	for i := 0; i < 5; i++ {
// 		emailChan <- fmt.Sprintf("%d@example.com", i)
// 		time.Sleep(time.Second)
// 	}
// 	close(emailChan)

// 	fmt.Println("Done Sending.. ")

// 	<-done

	// // Unbuffered Channels Example

	// // Create a channel of type bool
	// done := make(chan bool)
	// // Start a goroutine that signals when done
	// go task(done)
	// // Wait for the goroutine to finish (blocks until done receives a value)
	// ans := <-done
	// fmt.Println(ans) // Output: true

	// // Create a channel of type int
	// numChan := make(chan int)
	// // Start a goroutine to process a number from the channel
	// go processNum(numChan)
	// // Send the number 5 to the channel (blocks until processNum receives it)
	// numChan <- 5

	// // Create a channel to receive the sum result
	// result := make(chan int)
	// // Start a goroutine to compute the sum of 5 and 6
	// go sum(result, 5, 6)
	// // Receive the result from the channel (blocks until sum sends it)
	// res := <-result
	// fmt.Println(res) // Output: 11

	// Deadlock Example (commented out)
	// messageChan := make(chan string)
	// messageChan <- "ping" // This line will cause a deadlock because there is no goroutine receiving from the channel
	// msg := <-messageChan
	// fmt.Println(msg)
}
