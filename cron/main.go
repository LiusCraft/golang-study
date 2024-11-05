package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 0/5 * * * ?", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
	<-make(chan int)
}
