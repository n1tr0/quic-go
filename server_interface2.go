package quic

/*
import (
	"crypto/tls"
	"net"
	"time"
)

type Config struct {
	TLSConfig *tls.Config
	// called in a separate go routine
	ConnState func(*Session, ConnState)
	// Defaults to ...
	HandshakeTimeout time.Duration
	// more timeouts...
	//
	MaxIncomingStreams int
}

type Listener interface {
	Close() error
	Addr() net.Addr
}

func ListenAddr(addr string, config *Config) (Listener, error)     {}
func Listen(conn net.PacketConn, config *Config) (Listener, error) {}

// If config.ConnState is set, returns immediately after a version was negotiated.
// If config.ConnState is nil, returns only after a fordward secure connection is established.
func DialAddr(addr string, config *Config) (Session, error)     {}
func Dial(conn net.PacketConn, config *Config) (Session, error) {}
*/