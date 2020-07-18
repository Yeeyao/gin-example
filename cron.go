package main

import (
	"log"
	"time"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run modles.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run modles.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	// 启动新的计时器，reset 重置定时器
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
