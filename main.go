package main

import (
	webscraper "APKElasticSearch/WebScraper"
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {

	fmt.Println("running task")
	LogExecuteTime("Start")
	webscraper.ReadWebData()
	LogExecuteTime("End")

}
func LogExecuteTime(action string) {
	//open File if file is exists else create file and write append the file
	f, err := os.OpenFile("Scheduler.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	dt := time.Now()
	//fmt.Println("Current date and time is: ", dt.String())
	if _, err = f.WriteString("Executed" + action + "DateTime:" + dt.Format("2006-01-02 15:04:05") + "\n"); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("APK Project")

	//gocron.Every(2).Seconds().Do(task)
	s := gocron.NewScheduler(time.UTC)
	//s.Every(20).Seconds().Do(task)
	//s.Every(1).Day().At(time.Now()).Do(task)

	//job, _ := s.Every(2).Minutes().Do(task)
	job, _ := s.Every(1).Day().At("16:14").Do(task)
	s.StartAsync()
	fmt.Println("Last run:", job.LastRun())
	// starts the scheduler and blocks current execution path
	s.StartBlocking()

	for true {
		fmt.Println("Infinite Loop")
		time.Sleep(time.Second)
	}

}
