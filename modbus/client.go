package modbus

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

type Client struct {
	client *modbus.ModbusClient
	config *modbus.ClientConfiguration
	isOpen bool
}

func NewModbusClientRTU(url string) (*Client, error) {
	cfg := &modbus.ClientConfiguration{
		URL:      url,
		Speed:    9600,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  5 * time.Second,
	}

	client, err := NewModbusClientRTUWithConfig(cfg)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewModbusClientRTUWithConfig(cfg *modbus.ClientConfiguration) (*Client, error) {
	client, err := modbus.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create modbus client %w", err)
	}

	if err = client.SetEncoding(modbus.BIG_ENDIAN, modbus.LOW_WORD_FIRST); err != nil {
		return nil, fmt.Errorf("failde to modbus client encoding setting %w", err)
	}

	return &Client{client: client, config: cfg, isOpen: false}, nil
}

func (m *Client) ConnectionOpen() error {
	if m.nilCheck() {
		return fmt.Errorf("modbus client is not initialized")
	}

	if err := m.client.Open(); err != nil {
		return fmt.Errorf("failed to open modbus connection %w", err)
	}

	m.isOpen = true

	return nil
}

func (m *Client) ConnectionClose() error {
	if m.nilCheck() {
		return fmt.Errorf("modbus client is not initialized")
	}

	if err := m.client.Close(); err != nil {
		return fmt.Errorf("failed to close modbus connection %w", err)
	}

	m.isOpen = false

	return nil
}

func (m *Client) ReadRegister(addr uint16) (uint16, error) {
	if m.isConnectionOpen() {
		return 0, fmt.Errorf("modbus connection not opened")
	}

	reg, err := m.client.ReadRegister(addr, modbus.HOLDING_REGISTER)
	if err != nil {
		return 0, fmt.Errorf("failed to read register because: %w from %d", err, addr)
	}

	return reg, nil
}

func (m *Client) ReadRegisters(addr uint16, quantity uint16) ([]uint16, error) {
	if m.isConnectionOpen() {
		return nil, fmt.Errorf("modbus connection not opened")
	}

	regs, err := m.client.ReadRegisters(addr, quantity, modbus.HOLDING_REGISTER)
	if err != nil {
		return nil, fmt.Errorf("failde to read registers because: %w from %d", err, addr)
	}

	return regs, nil
}

func (m *Client) nilCheck() bool {
	return m == nil || m.client == nil
}

func (m *Client) isConnectionOpen() bool {
	//todo 커넥션 상태 확인 로직 추가 필요
	return m.isOpen
}
