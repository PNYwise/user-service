package config

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/PNYwise/user-service/internal/domain"
)

var producer sarama.SyncProducer

func kafkaConn(extConf *domain.ExtConf) sarama.SyncProducer {
	// Kafka broker address
	brokerList := []string{
		fmt.Sprintf("%s:%d", extConf.Kafka.Host, extConf.Kafka.Port),
	}

	// Initialize Kafka producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatal("error to creates a new sync producer")
	}
	log.Println("Connected to Kafka")
	return producer
}
func GetKafkaProducer(extConf *domain.ExtConf) sarama.SyncProducer {
	if producer == nil {
		producer = kafkaConn(extConf)
	}
	return producer
}
