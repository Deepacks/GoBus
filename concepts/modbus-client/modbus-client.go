package main

import (
	"fmt"

	"github.com/simonvetter/modbus"
)

func main() {
    var client  *modbus.ModbusClient
    var err     error
    var regs    []uint16

    client, err = modbus.NewClient(&modbus.ClientConfiguration{
        URL: "tcp://0.0.0.0:5502",
    });

    if err != nil {
		fmt.Printf("failed to create modbus client: %v\n", err)
		return
	}

	err = client.Open()
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		return
	}
    
    err = client.WriteRegister(0x0065, 0x0001)
    if err != nil {
		fmt.Printf("failed to write register 0x0065: %v\n", err)
		return
	}

    regs, err = client.ReadRegisters(0x0065, 1, modbus.HOLDING_REGISTER)
    if err != nil {
		fmt.Printf("failed to read register 0x0065: %v\n", err)
	} else {
		fmt.Printf("register 0x0065: 0x%04x\n", regs[0])
	}

    err = client.Close()
	if err != nil {
		fmt.Printf("failed to close connection: %v\n", err)
	}
}