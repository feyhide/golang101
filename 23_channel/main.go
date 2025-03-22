//message communication between wait group

package main

// ((((-1))))

// func processNum(numChannel chan int) {
// 	for num := range numChannel {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// ((((-2))))

// func sum(result chan int, num1 int, num2 int) {
// 	add := num1 + num2
// 	result <- add
// }

// ((((-3))))

// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("Processing...")
// 	time.Sleep(time.Second)
// }

// ((((-4))))
// ALL METHOD CHANNEL == func emailSender(emailChan chan string, done chan bool)
// RECIEVE ONLY CHANNEL == func emailSender(emailChan <-chan string, done chan bool)
// SEND ONLY CHANNEL == func emailSender(emailChan chan<- string, done chan bool)

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

func main() {
	// ((((-0))))

	// messageChannel := make(chan string)
	// messageChannel <- "ping" //sending in , blocking until second side is there to recieved
	// msg := <-messageChannel  // recieving
	// fmt.Println(msg)

	// ((((-1))))
	// ----SENDING MESSAGE INTO PROCESS----

	// numChannel := make(chan int)
	// go processNum(numChannel)
	// for {
	// 	numChannel <- rand.Intn(100)
	// }

	// ((((-2))))
	// ----RECIEVING MESSAGE fROM PROCESS----

	// result := make(chan int)
	// go sum(result, 4, 5)
	// addition := <-result // somewhat blocking
	// fmt.Println("Addition from sum in main", addition)

	// ((((-3))))
	// ----SAME AS WE DONE SYNCHRONIZING IN WAITGROUP
	// DO WAITGROUP IF MULTIPLE GO ROUTINES----

	// done := make(chan bool)
	// go task(done)
	// result := <-done
	// fmt.Println("processed :", result)

	// ((((-4))))
	// ----BUFFERED CHANNEL----

	// 10 message can be stored in buffer as unblocking
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// close(emailChan) // Close channel after sending all emails

	// <-done
	// fmt.Println("Email sended")

	// ((((-5))))
	// ----LISTENING TO MULTIPLE CHANNELS----

	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func() {
	// 	chan1 <- 10
	// }()

	// go func() {
	// 	chan2 <- "bingo"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Val := <-chan1:
	// 		fmt.Println("recieved from channel 1", chan1Val)

	// 	case chan2Val := <-chan2:
	// 		fmt.Println("recieved from channel 2", chan2Val)
	// 	}
	// }

	// ((((-6))))
}
