package main

import (
	"ryg-email-service/conf"
	"ryg-email-service/rabbit_mq"
)

func main() {
	cnf := conf.LoadConfig()

	qm := rabbit_mq.NewQueueConsumerManager(cnf)
	defer qm.Close()

	qm.Start()
}
