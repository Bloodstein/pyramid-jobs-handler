package service

import "log"

func PopJob(guid string, queue string) {
	log.Printf("[%s] I am a worker of queue %s", guid, queue)
}
