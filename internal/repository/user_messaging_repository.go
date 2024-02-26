package repository

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/PNYwise/user-service/internal/domain"
)

type userMessagingRepository struct {
	producer sarama.SyncProducer
	extConf  *domain.ExtConf
}

func NewPostMessagingRepository(producer sarama.SyncProducer, extConf *domain.ExtConf) domain.IUserMessagingRepository {
	return &userMessagingRepository{
		producer: producer,
		extConf:  extConf,
	}
}

// PublishMessage implements domain.KafkaPostRepository.
func (u *userMessagingRepository) PublishMessage(post *domain.User) error {
	jsonMessage, err := json.Marshal(post)
	if err != nil {
		fmt.Printf("error to mashal %v", err)
		return err
	}
	_, _, err = u.producer.SendMessage(&sarama.ProducerMessage{
		Topic: u.extConf.Kafka.Topic,
		Value: sarama.ByteEncoder(jsonMessage),
	})
	return err
}
