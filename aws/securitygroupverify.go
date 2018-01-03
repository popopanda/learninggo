package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	groupIds := []string{"groupa", "groupb", "groupc"}

	//loads credentials from default locations
	//the environment variables, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role
	sess := session.Must(session.NewSession())

	svc := ec2.New(sess)

	result, err := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupIds: aws.StringSlice(groupIds),
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, securitygroup := range result.SecurityGroups {
		for _, ipPerms := range securitygroup.IpPermissions {
			switch *ipPerms.FromPort {
			case *aws.Int64(80):
				writeIPFile(ipPerms.IpRanges, *ipPerms.FromPort, *securitygroup.GroupId)
				fmt.Println("Port 80:")
				getIP(ipPerms.IpRanges)
			case *aws.Int64(443):
				writeIPFile(ipPerms.IpRanges, *ipPerms.FromPort, *securitygroup.GroupId)
				fmt.Println("\nPort 443:")
				getIP(ipPerms.IpRanges)
			case *aws.Int64(4433):
				writeIPFile(ipPerms.IpRanges, *ipPerms.FromPort, *securitygroup.GroupId)
				fmt.Println("\nPort 4433:")
				getIP(ipPerms.IpRanges)
			default:
				fmt.Println("No Matches Found!")
			}
		}
	}
}

func getIP(ipList []*ec2.IpRange) {
	for _, ip := range ipList {
		fmt.Println(*ip.CidrIp)
	}
}

func writeIPFile(writeIPs []*ec2.IpRange, port int64, groupID string) {
	homePath := os.Getenv("HOME")
	filePath := fmt.Sprintf("%v/Desktop/test.txt", homePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist... Creating...")
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		fmt.Println("File Created...")

	}
	output, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()
	fmt.Fprintf(output, "\n%v:\n", groupID)
	fmt.Fprintf(output, "Port %v:\n", port)

	for _, IP := range writeIPs {
		_, err := fmt.Fprintf(output, "%v\n", *IP.CidrIp)

		if err != nil {
			log.Fatal(err)
		}
	}
}
