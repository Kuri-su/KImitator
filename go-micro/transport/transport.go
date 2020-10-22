// Package transport is an interface for synchronous connection based communication
package transport

import (
	"time"
)

// Transport is an interface which is used for communication between
// services. It uses connection based socket send/recv semantics and
// has various implementations; http, grpc, quic.
type Transport interface {
	// Actual effect
	Init(...Option) error
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)

	// simple
	Options() Options
	String() string
}

type Message struct {
	Header map[string]string
	Body   []byte
}

type Socket interface {
	Recv(*Message) error
	Send(*Message) error

	// simple
	Close() error
	Local() string
	Remote() string
}

// Semantics
type Client interface {
	Socket
}

type Listener interface {
	Accept(func(Socket)) error

	// simple
	Close() error
	Addr() string
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport Transport = newHTTPTransport()

	DefaultDialTimeout = time.Second * 5
)

func NewTransport(opts ...Option) Transport {
	return newHTTPTransport(opts...)
}
