package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/leaanthony/clir"
)

func main() {
	cli := clir.NewCli("smallwatch", "Cloudwatch log cleanup for Dev Environments", "v0.0.1")

	listCmd := cli.NewSubCommand("list", "List out CloudWatch logs")

	var listCwPrefix string = "/"

	listCmd.StringFlag("prefix", "The prefix of the log group to list", &listCwPrefix)
	listCmd.Action(func() error {
		cfg, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		client := cloudwatchlogs.NewFromConfig(cfg)

		var cwLogNumber int

		cloudWatchLogsPaginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(client, &cloudwatchlogs.DescribeLogGroupsInput{
			LogGroupNamePrefix: &listCwPrefix,
		})

		for cloudWatchLogsPaginator.HasMorePages() {
			page, err := cloudWatchLogsPaginator.NextPage(context.TODO())
			if err != nil {
				log.Fatalf("ERROR: unable to get next page: %v", err)
			}

			for _, logGroup := range page.LogGroups {
				fmt.Printf("Log Group: %v\n", *logGroup.LogGroupName)
				cwLogNumber += 1
			}
		}
		if cwLogNumber == 0 {
			fmt.Printf("No log groups found with prefix: %v\n", cwLogNumber)
		} else {
			fmt.Printf("Found %v log groups\n", cwLogNumber)
		}
		return nil
	})

	setCwLgCmd := cli.NewSubCommand("reduce", "Reduce CloudWatch logs retention size")

	var setCwLgNoDryRun bool
	var setCwLgPrefix string
	var setCwLgRetention int = 30

	setCwLgCmd.BoolFlag("no-dry-run", "Performs actual modification", &setCwLgNoDryRun)
	setCwLgCmd.StringFlag("prefix", "The prefix of the log group to reduce", &setCwLgPrefix)
	setCwLgCmd.IntFlag("days", "The number of days to set retention. Defaults to 30 days", &setCwLgRetention)

	setCwLgCmd.Action(func() error {
		cfg, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		client := cloudwatchlogs.NewFromConfig(cfg)
		var cwLogNumber int

		cloudWatchLogsPaginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(client, &cloudwatchlogs.DescribeLogGroupsInput{
			LogGroupNamePrefix: &setCwLgPrefix,
		})

		for cloudWatchLogsPaginator.HasMorePages() {
			page, err := cloudWatchLogsPaginator.NextPage(context.TODO())
			if err != nil {
				log.Fatalf("ERROR: unable to get next page: %v", err)
			}

			for _, logGroup := range page.LogGroups {
				if setCwLgNoDryRun {
					fmt.Printf("Reducing Log Group: %v\n", *logGroup.LogGroupName)
					client.PutRetentionPolicy(context.TODO(), &cloudwatchlogs.PutRetentionPolicyInput{
						LogGroupName:    logGroup.LogGroupName,
						RetentionInDays: aws.Int32(int32(setCwLgRetention)),
					})
				} else {
					fmt.Printf("DRY RUN - Reducing Log Group: %v\n", *logGroup.LogGroupName)
				}
				cwLogNumber += 1
			}
		}
		if cwLogNumber == 0 {
			fmt.Printf("No log groups found with prefix: %v\n", setCwLgPrefix)
		} else {
			fmt.Printf("Set %v log groups to %v days retention\n", cwLogNumber, setCwLgRetention)
		}
		return nil
	})

	deleteCmd := cli.NewSubCommand("delete", "List of CloudWatch Logs to delete")

	var deleteCwPrefix string
	var deleteCwNoDryRun bool

	deleteCmd.BoolFlag("no-dry-run", "Performs actual deletion", &deleteCwNoDryRun)
	deleteCmd.StringFlag("prefix", "The prefix of the log group to delete", &deleteCwPrefix)

	deleteCmd.Action(func() error {
		cfg, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatalf("unable to load SDK config, %v\n", err)
		}

		client := cloudwatchlogs.NewFromConfig(cfg)
		var cwLogNumber int
		cloudWatchLogsPaginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(client, &cloudwatchlogs.DescribeLogGroupsInput{
			LogGroupNamePrefix: &deleteCwPrefix,
		})

		for cloudWatchLogsPaginator.HasMorePages() {
			page, err := cloudWatchLogsPaginator.NextPage(context.TODO())
			if err != nil {
				log.Fatalf("ERROR: unable to get next page: %v\n", err)
			}

			for _, logGroup := range page.LogGroups {
				if deleteCwNoDryRun {
					fmt.Printf("Deleting Log Groups: %v\n", *logGroup.LogGroupName)
					client.DeleteLogGroup(context.TODO(), &cloudwatchlogs.DeleteLogGroupInput{
						LogGroupName: logGroup.LogGroupName,
					})
				} else {
					fmt.Printf("DRY RUN - Deleting Log Group: %v\n", *logGroup.LogGroupName)
				}
				cwLogNumber += 1
			}
		}
		if cwLogNumber == 0 {
			fmt.Printf("No log groups found with prefix: %v\n", deleteCwPrefix)
		} else {
			fmt.Printf("Deleted %v log groups\n", cwLogNumber)
		}
		return nil
	})

	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}
