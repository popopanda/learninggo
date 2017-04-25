package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	kafkaCMD := "/opt/kafka/kafka_2.10-0.10.2.0/bin/kafka-topics.sh --zookeeper 172.16.0.23:2181 --describe | /bin/grep ReplicationFactor:1"
	cmd, err := exec.Command("/bin/bash", "-c", kafkaCMD).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(cmd))
}

// "/opt/kafka/kafka_2.10-0.10.2.0/bin/kafka-topics.sh", "--zookeeper", "172.16.0.23:2181", "--describe", "|", "/bin/grep", "ReplicationFactor:1"
