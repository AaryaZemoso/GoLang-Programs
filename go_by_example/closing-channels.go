package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved Job : ", j)
			} else {
				fmt.Println("All jobs are recieved")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sent Job : ", i)
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}
