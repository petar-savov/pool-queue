package main

import (
	"fmt"
	"time"
)

type Message struct {
	Key, Value string
}

type Topic struct {
	Name      string
	Partition int
	Messages  []Message
}

type Broker struct {
	ID     int
	Topics map[string][]Topic
}

func (b *Broker) createTopic(topicName string, partitionCount int) {
	topics := make([]Topic, partitionCount)
	for i := 0; i < partitionCount; i++ {
		topics[i] = Topic{
			Name:      topicName,
			Partition: i,
			Messages:  []Message{},
		}
	}
	b.Topics[topicName] = topics
}

func (b *Broker) produceMessage(topicName string, partition int, key, value string) {
	message := Message{
		Key:   key,
		Value: value,
	}
	b.Topics[topicName][partition].Messages = append(b.Topics[topicName][partition].Messages, message)
}

func (b *Broker) consumeMessage(topicName string, partition int) {
	messages := b.Topics[topicName][partition].Messages
	for _, message := range messages {
		fmt.Printf("Consuming message: Key=%s, Value=%s\n", message.Key, message.Value)
		time.Sleep(time.Second) // Simulating message processing
	}
}

func main() {
	broker := Broker{
		ID:     1,
		Topics: make(map[string][]Topic),
	}

	broker.createTopic("my-topic", 3)
	broker.produceMessage("my-topic", 0, "key1", "value1")
	broker.produceMessage("my-topic", 1, "key2", "value2")
	broker.produceMessage("my-topic", 2, "key3", "value3")

	go broker.consumeMessage("my-topic", 0)
	go broker.consumeMessage("my-topic", 1)
	go broker.consumeMessage("my-topic", 2)

	time.Sleep(5 * time.Second) // Wait for consumers to finish processing messages
}
