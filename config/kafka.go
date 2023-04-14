
package config

import (
	"os"
	"WorkerWithCheckHealth/exception"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConnection() *kafka.Producer {

	configMap := kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST")}

	// Variable p holds the new Producer instance.
	producer, err := kafka.NewProducer(&configMap)
	exception.PanicIfNeeded(err)

	return producer
}
