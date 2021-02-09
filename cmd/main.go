package main

import (
	"fmt"
	_ "job_server/internal/services"
	"job_server/pkg/job"
	"job_server/store"
	"time"
)

func main() {
	redis := new(store.Redis)
	err := redis.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	mysql := new(store.Mysql)
	mysql.Get()
	log := new(store.Log)
	log.Get().Debug("server started at %v", time.Now())
	fmt.Printf("server started at %v \n", time.Now())
	go job.BillNotice()
	select {}
}
