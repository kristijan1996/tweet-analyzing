package main

import (
	"fmt"
	"sync"
	"time"
)

// -------------- First solution ----------------
// Produce function accepts stream of tweets, reads
// them one by one until the last one has be read
// and returns an array of tweets
func produce1(stream Stream) (tweets []*Tweet) {
	for {
		tweet := stream.Next()
		if tweet == nil {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

// Consume accepts array of tweets and for each checks if
// it is talking about Go
func consume1(tweets []*Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.User, "\ttweets about golang")
		} else {
			fmt.Println(t.User, "\tdoes not tweet about golang")
		}
	}
}

func solution1() {
	// Start the stopwatch
	start := time.Now()

	// Create a stream of data
	stream := GetMockStream()

	// Call the produce process
	tweets := produce1(stream)
	// and afterwards the consume process
	consume1(tweets)

	// Note that there is no concurency here, so this
	// solution is the slowest one
	fmt.Printf("\nsolution1 took %s\n", time.Since(start))
}

// ------------------- Second solution -----------------
func produce2(stream Stream, c chan<- *Tweet) {
	for {
		// Read the next tweet from the stream
		tweet := stream.Next()

		// If all the tweets have been read, close the
		// communication channel and exit
		if tweet == nil {
			close(c)
			return
		}

		// Send the tweet over the channel
		c <- tweet
	}
}

func consume2(t *Tweet) {
	if t.IsTalkingAboutGo() {
		fmt.Println(t.User, "\ttweets about golang")
	} else {
		fmt.Println(t.User, "\tdoes not tweet about golang")
	}
}

func solution2() {
	// Start the stopwatch
	start := time.Now()

	// Create a stream of data
	stream := GetMockStream()

	// Create a communication channel
	c := make(chan *Tweet)

	// Call the produce as goroutine, creating tweets
	// from the stream in the background and sending them
	// over channel
	go produce2(stream, c)

	for {
		// Blocking wait for a tweet over the channel
		t, open := <-c

		// If there is a tweet to be consumed, do so,
		// otherwise break
		if !open {
			break
		} else {
			consume2(t)
		}
	}

	// Note that produce is here called concurently, which
	// makes this solution faster. The only downside is that
	// bolierplate code over waiting for a tweet to arrive via
	// channel is in the main function
	fmt.Printf("\nsolution2 took %s\n", time.Since(start))
}

// ----------------------- Third solution ------------------------
func produce3(stream Stream, c chan<- *Tweet) {
	// Read the stream one tweet at a time, and when
	// all of them have been pulled from the stream,
	// close the communication channel
	for {
		tweet := stream.Next()

		if tweet == nil {
			close(c)
			return
		}

		c <- tweet
	}
}

func consume3(c <-chan *Tweet) {
	for {
		// Blocking call waiting for a tweet to arrive
		t, open := <-c

		if !open {
			return
		} else if t.IsTalkingAboutGo() {
			fmt.Println(t.User, "\ttweets about golang")
		} else {
			fmt.Println(t.User, "\tdoes not tweet about golang")
		}
	}
}

func solution3() {
	// Start the stopwatch
	start := time.Now()

	// Create a stream of data
	stream := GetMockStream()

	// Create a communication channel
	c := make(chan *Tweet)

	// Create a waitgroup which monitors if the
	// goroutines have finished their job
	var wg sync.WaitGroup

	// Add a mark of a running goroutine before its actual
	// call in order to fix sync issues (it can happen that
	// wg.Done() is called before wg.Add(1) if we place them
	// both inside goroutines)
	// Both goroutines are called like literals since they
	// themselves don't have to know about the WaitGroup, so
	// we avoid passing wg in argument list
	wg.Add(1)
	go func(Stream, chan<- *Tweet) {
		produce3(stream, c)
	}(stream, c)

	go func(<-chan *Tweet) {
		consume3(c)
		wg.Done()
	}(c)

	// Make the main func wait for goroutines to finish
	wg.Wait()
	// Note that this is the full utilization of concurency,
	// channels and WaitGroup
	fmt.Printf("\nsolution3 took %s\n", time.Since(start))
}

// ----------------------- main ------------------------
func main() {
	// Pick which solution to run
	solution1()
	solution2()
	solution3()
}
