//
// {
//   "remoteAddress":"10.0.0.1",
//   "eventType":"event_stream_attached",
//   "timestamp":"2018-02-12T23:44:37.684Z"
// }

package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	RemoteAddress string `json:"remoteAddress"`
	EventType     string `json:"eventType"`
	TimeStamp     string `json:"timestamp"`
}

func main() {

	dataJson := `[{"remoteAddress":"10.0.0.1","eventType":"event_stream_attached","timestamp":"2018-02-12T23:44:37.684Z"}]`
	var app []App
	json.Unmarshal([]byte(dataJson), &app)
	fmt.Printf("data : %+v", app)
}
