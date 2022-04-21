package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// массив процессов и их каналов, чтобы в канал передать сигнал об остановке процесса

var runnedProcesses map[string]Process

func init() {
	runnedProcesses = make(map[string]Process)
}

func main() {

	router := gin.New()

	router.GET("run", func(ctx *gin.Context) {
		log.Println("Running...")
		newGuid := uuid.NewString()
		p := Process{
			Guid:  newGuid,
			Abort: make(chan bool),
		}
		go p.Execute()
		runnedProcesses[newGuid] = p
		ctx.JSON(200, map[string]string{
			"result": "ok",
		})
	})
	router.GET("status", func(ctx *gin.Context) {
		result := make([]string, len(runnedProcesses))
		for _, p := range runnedProcesses {
			result = append(result, p.Guid)
		}
		ctx.JSON(200, result)
	})
	router.GET("stop", func(ctx *gin.Context) {
		log.Println("Stopping...")
		for _, p := range runnedProcesses {
			p.Abort <- true
		}
		ctx.JSON(200, map[string]string{
			"result": "ok",
		})
	})

	router.Run(":8000")
}
