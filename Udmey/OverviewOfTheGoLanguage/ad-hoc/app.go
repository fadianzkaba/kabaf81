package app

import (
	"log"
)

type Service struct {
	TaskQuery string
	Fname     string
	Sname     string
}

func app() {
	/*currTime := time.Now()
	fmt.Println("The current time is:", currTime)

	addHours := 24 * time.Hour
	fmt.Println("This the hour:", addHours)

	newTime := currTime.Add(addHours)
	fmt.Println("The new time is", newTime)
	fmt.Println(newTime.Format("2006-01-02T15:04:05.000Z"))*/

	var b Service
	addServices(&b)
	log.Println(b.TaskQuery)

}

func addServices(s *Service) {
	s.TaskQuery = "Testing"
}
