package kafka

import (
	gen "cards-operations/api/gen"
	"cards-operations/core"
	"context"
	"encoding/json/v2"
	"fmt"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaClient(config core.Config) (*kgo.Client, error) {
	kafka := config.Kafka

	client, err := kgo.NewClient(
		kgo.SeedBrokers(kafka.Brokers...),
		kgo.ConsumerGroup(kafka.GroupID),
		kgo.ConsumeTopics(kafka.Topic),
	)

	return client, err
}

func Listen(ctx context.Context, client *kgo.Client) {
	for {
		messages := client.PollFetches(ctx)
		if errors := messages.Errors(); len(errors) > 0 {
			log.Printf("errors: %v", errors)
		}

		iter := messages.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			event, err := readEvent(record)
			if err != nil {
				continue
			}
			fmt.Printf("%+v\n", *event)
		}

	}
}

func readEvent(record *kgo.Record) (*gen.OperationEvent, error) {
	var event gen.OperationEvent
	err := json.Unmarshal(record.Value, &event)
	return &event, err
}
