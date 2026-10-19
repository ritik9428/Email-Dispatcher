package main

import "sync"

type Recipient struct {
	Name  string
	Email string
}

func main() {

	var wg sync.WaitGroup
	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipient("./emails.csv", recipientChannel)
	}()

	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}
	wg.Wait()
}
