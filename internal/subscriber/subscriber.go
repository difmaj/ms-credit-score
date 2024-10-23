package subscriber

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/difmaj/ms-credit-score/internal/pkg/logger"
	"go.uber.org/zap"
)

// HandlerFunction type.
type HandlerFunction func(msg *message.Message) error

// Subscriber struct.
type Subscriber struct {
	subscriber *kafka.Subscriber
	messages   []*Message
}

// Message struct.
type Message struct {
	handlers []HandlerFunction
	ch       <-chan *message.Message
}

// New creates a new instance of the Subscriber struct.
func New() (*Subscriber, error) {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"kafka:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		subscriber: subscriber,
		messages:   make([]*Message, 0),
	}, nil
}

// Subscribe subscribes to a topic.
func (s *Subscriber) Subscribe(context context.Context, topic string, handlers ...HandlerFunction) error {
	messages, err := s.subscriber.Subscribe(context, topic)
	if err != nil {
		return err
	}

	s.messages = append(s.messages, &Message{
		handlers: handlers,
		ch:       messages,
	})
	return nil
}

// Close closes the subscriber.
func (s *Subscriber) Close() error {
	return s.subscriber.Close()
}

// Process processes the messages.
func (s *Subscriber) Process() {
	for _, messages := range s.messages {
		go func(messages *Message) {
			for msg := range messages.ch {
				for _, handler := range messages.handlers {
					if err := handler(msg); err != nil {
						logger.Logger.Error("Subscriber.Process.handler", zap.Error(err), zap.String("topic", msg.Metadata.Get("topic")))
						return
					}
				}
				msg.Ack()
			}
		}(messages)
	}
}
