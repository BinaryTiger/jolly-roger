package stream

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

type StreamEngine interface {
	Pub(string, []byte) error
}

type natsEngine struct {
	conn *nats.Conn
}

func NewFromViperSettings() (StreamEngine, error) {
	engine := viper.GetString("stream.engine")

	switch engine {
	case "nats":
		return newNatsEngine(viper.GetString("stream.connection_string"))
	case "":
		return nil, errors.New("[stream.engine] needs to be defined in the config file, not assuming any default value")
	default:
		return nil, errors.New(fmt.Sprintf("%s is not a supported stream option", engine))
	}
}

func newNatsEngine(connection_url string) (natsEngine, error) {
	conn, err := nats.Connect(connection_url)

	return natsEngine{
		conn: conn,
	}, err
}

func (s natsEngine) Pub(subject string, message []byte) error {
	err := s.conn.Publish("default", []byte(message))
	return err
}
