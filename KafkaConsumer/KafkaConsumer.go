package main

import (
	"fmt"
	"regexp"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := "localhost"
	groupID := "myGroup"
	
	// Create an AdminClient
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}

	// List topics
	metadata, err := adminClient.GetMetadata(nil, true, 5000)  // 5000 is the timeout in ms
	if err != nil {
		panic(err)
	}

	// Find topics based on your criteria
	var topicsToSubscribe []string
	for _, topic := range metadata.Topics {
		// Check for "test" topic
		if topic.Topic == "test" {
			topicsToSubscribe = append(topicsToSubscribe, topic.Topic)
		}
		
		// Check for regex match
		match, _ := regexp.MatchString("^aRegex.*[Tt]opic", topic.Topic)
		if match {
			topicsToSubscribe = append(topicsToSubscribe, topic.Topic)
		}

		fmt.Println("Subscribed To Topics: ", topicsToSubscribe)
	}
	
	// Create a new consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	// Subscribe to the topics
	if len(topicsToSubscribe) > 0 {
		c.SubscribeTopics(topicsToSubscribe, nil)
	} else {
		fmt.Println("No topics found to subscribe!")
		return
	}

	// Consumer logic (kept simple for the example)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// Handle errors, but don't panic
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
