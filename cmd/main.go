package main

import "ryg-email-service/rabbit_mq"

func main() {
	conn, ch, q := rabbit_mq.ConnectAMQ()
	defer ch.Close()
	defer conn.Close()

	rabbit_mq.ConsumeForever(ch, &q)
}
