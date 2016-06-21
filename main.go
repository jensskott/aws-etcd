package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Connect() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})
	return
}

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	connect()
	params := &ec2.DescribeTagsInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{ // Required
				Name: aws.String("key"),
				Values: []*string{
					aws.String("Name"), // Required
					// More values...
				},
			},
			{ // Required
				Name: aws.String("value"),
				Values: []*string{
					aws.String("etcd-master"), // Required
					// More values...
				},
			},
		},
		MaxResults: aws.Int64(10),
		//NextToken:  aws.String("key"),
	}
	resp, err := svc.DescribeTags(params)

	check(err)

	// Pretty-print the response data.
	fmt.Println(resp)
}

func runningInstances() {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}
}
