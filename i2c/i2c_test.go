package i2c

import (
	"testing"
)

type mock struct{}

func (m *mock) SetAddress(_ byte) error                        { return nil }
func (m *mock) ReadBytes(addr byte, num int) ([]byte, error)   { return []byte{}, nil }
func (m *mock) WriteBytes(addr byte, value []byte) error       { return nil }
func (m *mock) ReadFromReg(addr, reg byte, value []byte) error { return nil }
func (m *mock) WriteToReg(addr, reg byte, value []byte) error  { return nil }

func MockBus() Bus { return new(mock) }

func TestI2c(t *testing.T) {
	var b bus
	if b.f != nil {
		t.Error()
	}
}
