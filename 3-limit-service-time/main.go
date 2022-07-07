//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

//******************** - Добавленный код

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

const AvailableTime int64 = 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	//********************
	var throttle <-chan time.Time
	var s int64
	//Заведем таймер на 10 секунд для пользователя у которого нет премиума
	if !u.IsPremium {
		//Если доступное время вышло, то "убиваем" процесс
		if u.TimeUsed >= AvailableTime {
			return false
		}
		//Заводим таймер на оставшееся время
		throttle = time.Tick(time.Second * time.Duration(AvailableTime-u.TimeUsed))
	}

	done := make(chan bool, 1)
	go func() {
		s = time.Now().Unix()
		process()
		done <- true
	}()

	select {
	case <-throttle:
		u.TimeUsed += AvailableTime
		return false
	case <-done:
		u.TimeUsed += time.Now().Unix() - s
		return true
	}

	//********************
	//было
	// process()
	// return true
}

func main() {
	RunMockServer()
}
