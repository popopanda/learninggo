package main

import (
	"log"
	"net/smtp"
	"os/exec"
)

func checkKafka() string {
	kafkaCMD := "/opt/kafka/kafka_2.10-0.10.2.0/bin/kafka-topics.sh --zookeeper 172.16.0.23:2181 --describe | /bin/grep ReplicationFactor:1"
	cmd, err := exec.Command("/bin/bash", "-c", kafkaCMD).Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", string(cmd))
	return cmd
}

func sendMail(message string) {
	username := ""
	password := ""
	to := ""
	subject := ""
	smtpServer := ""
	smtpPort := ""

	msg := message

	err := smtp.SendMail(addr, a, from, to, msg)

}

func main() {

}
