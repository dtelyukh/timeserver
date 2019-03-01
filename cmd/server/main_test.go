package main

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"
	"time"
)

var bs []byte

type ConnMock struct {
}

func (c ConnMock) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (c ConnMock) Write(b []byte) (n int, err error) {
	bs = b
	return 0, nil
}

func (c ConnMock) Close() error {
	return nil
}

func (c ConnMock) LocalAddr() net.Addr {
	return nil
}

func (c ConnMock) RemoteAddr() net.Addr {
	return AddrMock{}
}

func (c ConnMock) SetDeadline(t time.Time) error {
	return nil
}

func (c ConnMock) SetReadDeadline(t time.Time) error {
	return nil
}

func (c ConnMock) SetWriteDeadline(t time.Time) error {
	return nil
}

type AddrMock struct{}

func (a AddrMock) Network() string {
	return "localhost"
}

func (a AddrMock) String() string {
	return "localhost"
}

func TestHandlerOK(t *testing.T) {
	c := ConnMock{}
	handler(c)

	var received int32

	buf := bytes.NewReader(bs)
	err := binary.Read(buf, binary.BigEndian, &received)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second)
	now := int32(time.Now().Unix())
	if now <= received {
		t.Errorf("Received time is greather than expected. Received: %d, now (received +1 sec): %d", received, now)
	}
}
