package amqputils

import (
	"absolutetech/farm_report/global-lib/envutils"
	"absolutetech/farm_report/global-lib/errors"
	"encoding/json"
	"fmt"

	std_error "errors"

	"github.com/streadway/amqp"
)

var amqpURLs = map[envutils.Env]string{
	envutils.Testing:     "amqp://guest:guest@localhost:5672",
	envutils.Development: "amqp://guest:guest@localhost:5672",
	envutils.Staging:     "amqp://guest:guest@localhost:5672",
	envutils.Production:  "",
}

const (
	QueuePrefetchCount                        = 1
	QueueFarmReport                           = "Q_farm_report"
	WaitingQueueFetchProcess                  = "Q_farm_report_waiting"
	DeadQueueFarmReport                       = "Q_farm_report_dead"
	QueueFetchProcessWaitingMessageExpiration = "60000" // 1 minute
	QueueMaxRetries                           = 3
)

func getAMQPURLs(env envutils.Env) (string, error) {
	amqpURL, ok := amqpURLs[env]
	if !ok {
		return "", std_error.New("unknown environment")
	}
	return amqpURL, nil
}

// NewAMQPConnection returns a new *amqp.Connection.
func NewAMQPConnection(env envutils.Env) (*amqp.Connection, error) {
	amqpURL, err := getAMQPURLs(env)
	if err != nil {
		return nil, errors.Wrap(err, "get URLs")
	}
	amqpConnection, err := amqp.Dial(amqpURL)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	return amqpConnection, nil
}

// Open a channel to our RabbitMQ instance
func OpenAMQPChannel(amqpConnection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := amqpConnection.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// Declare Queue
func DeclareAMQPQueue(ch *amqp.Channel, QueueFarmReprtProcess string) (amqp.Queue, error) {
	amqpQueue, err := ch.QueueDeclare(
		QueueFarmReport,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return amqp.Queue{}, err
	}
	return amqpQueue, nil
}

func PublishMessageInQueue(ch *amqp.Channel, qu amqp.Queue, msg amqp.Publishing) error {
	// attempt to publish a message to the queue!
	err := ch.Publish(
		"",
		qu.Name,
		false,
		false,
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}

// AMQPMsg represents an AMQP message.
type AMQPMsg struct {
	FarmID     int64 `json:"farm_id"`
	TemplateID int64 `json:"template_id"`
	Retries    int   `json:"retries,omitempty"`
}

func SendMessage(ch *amqp.Channel, qu amqp.Queue, data *AMQPMsg) error {
	body, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "JSON marshal")
	}
	msg := amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	}

	publishMessageErr := PublishMessageInQueue(ch, qu, msg)
	if publishMessageErr != nil {
		return publishMessageErr
	}
	return nil
}

func SendRetryMessage(ch *amqp.Channel, qu amqp.Queue, data *AMQPMsg) error {
	if data.Retries >= QueueMaxRetries {
		return std_error.New("max retries reached")
	}
	data = copyAMQPMsg(data)
	data.Retries++
	body, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "JSON marshal")
	}

	msg := amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	}

	publishMessageErr := PublishMessageInQueue(ch, qu, msg)
	if publishMessageErr != nil {
		return nil
	}
	return nil
}

func SendMessageInDeadQueue(ch *amqp.Channel, qu amqp.Queue, data *AMQPMsg) error {
	if data.Retries >= QueueMaxRetries {
		err := std_error.New("max retries reached")
		return err
	}
	data = copyAMQPMsg(data)
	data.Retries++
	body, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "JSON marshal")
	}

	msg := amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	}

	publishMessageErr := PublishMessageInQueue(ch, qu, msg)
	if publishMessageErr != nil {
		return nil
	}
	return nil
}

func copyAMQPMsg(data *AMQPMsg) *AMQPMsg {
	tmp := *data
	return &tmp
}
