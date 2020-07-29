# tweet-analyzing

A sample program which shows the advantage of Go's built-in concurency capability

It creates a mock stream of tweets (which, in a real situation, would be pulled
down from Twitter) and then processes them

Three solutions are implemented and two of them show how utilizing Go's concurency,
channels and WaitGroups can help perform the processing much faster that what
traditional approach would allow

It is a slight modification and solution of a problem posted here 
https://github.com/loong/go-concurrency-exercises/tree/master/1-producer-consumer
