package main

import "ryg-email-service/rabbit_mq"

func main() {
	qm := rabbit_mq.NewQueueConsumerManager()
	defer qm.Close()

	qm.Start()
}
