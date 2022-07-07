//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

//******************** - Добавленный код

package main

import (
	"fmt"
	"time"
)

//********************
var tweets chan Tweet

//********************

//Исходная версия
// func producer(stream Stream) (tweets []*Tweet) {
// 	for {
// 		tweet, err := stream.Next()
// 		if err == ErrEOF {
// 			return tweets
// 		}
// 		tweets = append(tweets, tweet)
// 	}
// }

// func consumer(tweets []*Tweet) {
// 	for _, t := range tweets {
// 		if t.IsTalkingAboutGo() {
// 			fmt.Println(t.Username, "\ttweets about golang")
// 		} else {
// 			fmt.Println(t.Username, "\tdoes not tweet about golang")
// 		}
// 	}
// }

func producer(stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return
		}
		tweets <- *tweet
	}
}

func consumer() {
	for {
		t, ok := <-tweets
		//Если канал закрыт, то больше нечего читать
		if !ok {
			return
		}
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	//********************
	//Создадим буферизированный канал
	tweets = make(chan Tweet, 1)
	//********************

	//Исходная версия
	// // Producer
	// tweets := producer(stream)

	// // Consumer
	// consumer(tweets)

	go producer(stream)
	go consumer()

	fmt.Printf("Process took %s\n", time.Since(start))
}
