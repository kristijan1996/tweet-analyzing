package main

import (
	"strings"
	"time"
)

// Tweet structure
type Tweet struct {
	User string
	Text string
}

// IsTalkingAboutGo - Mock process which pretends to be a sophisticated procedure
// to analyse whether tweet is talking about go or not
func (t *Tweet) IsTalkingAboutGo() bool {
	// Simulate fixed delay of the process
	// so that the system is deterministic on each run
	time.Sleep(330 * time.Millisecond)

	// Check whether Golang or Gopher has been mentioned in
	// the tweet
	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher
}

// Fake array of Tweets, allegedly pulled down from
// Twitter itself. This is mock data used to imitate real situation
var twitterData = []Tweet{
	{
		"davecheney",
		"#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		"beertocode",
		"Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		"ironzeb",
		"Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		"beertocode",
		"Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		"vampirewalk666",
		"I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}

// Stream structure, which holds all of
// downloaded tweets, as well as the position of
// current tweet that has been used by the system
type Stream struct {
	pos    int
	tweets []Tweet
}

// GetMockStream is a blackbox function which returns
// a mock stream for demonstration purposes
func GetMockStream() Stream {
	return Stream{0, twitterData}
}

// Next is a mock process which reads the next tweet from the stream
func (s *Stream) Next() *Tweet {
	// Simulate delay
	time.Sleep(320 * time.Millisecond)

	// If every tweet from the stream has been read
	// alert it with a nil pointer
	if s.pos >= len(s.tweets) {
		return nil
	}

	// Otherwise, return the next tweet
	tweet := s.tweets[s.pos]
	s.pos++

	return &tweet
}
