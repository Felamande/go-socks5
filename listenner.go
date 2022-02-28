package socks5

import (
	"net"

	"golang.org/x/net/context"
)

type chanListener struct {
	l        net.Listener
	connChan chan net.Conn
	errChan  chan error
}

func NewChanlisten(ctx context.Context, network, address string) (*chanListener, error) {

	lc := net.ListenConfig{}
	rawl, err := lc.Listen(ctx, network, address)

	return &chanListener{
		l:        rawl,
		connChan: make(chan net.Conn, 16),
		errChan:  make(chan error, 16),
	}, err
}

func (l *chanListener) Accept() (chan net.Conn, chan error) {
	go func() {
		for {
			conn, err := l.l.Accept()
			if err != nil {
				l.errChan <- err
				continue
			}
			l.connChan <- conn
		}

	}()
	return l.connChan, l.errChan
}

func (l *chanListener) Close() error {
	return l.l.Close()
}
