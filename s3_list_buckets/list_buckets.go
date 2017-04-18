package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("There was no buckets listed")
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

}
