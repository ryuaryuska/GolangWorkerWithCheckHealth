
package helper

import (
	"encoding/json"
	"fmt"
	"WorkerWithCheckHealth/config"
	"WorkerWithCheckHealth/exception"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func PublishData(data interface{}, topic string) {
	p := config.KafkaConnection()

	msg, _ := json.Marshal(data)

	delivery_chan := make(chan kafka.Event, 10000)
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          msg},
		delivery_chan,
	)
	exception.PanicIfNeeded(err)

	e := <-delivery_chan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(delivery_chan)
}
