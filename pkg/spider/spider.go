package spider

import (
	"log"
	"time"
)

func DownCourse() {
	for {
		log.Println("in downCourse")
		time.Sleep(60 * time.Second)
	}
}
