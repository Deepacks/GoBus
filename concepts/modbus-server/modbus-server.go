package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	var server	*modbus.ModbusServer
	var err		error
	var eh		*exampleHandler
	var ticker	*time.Ticker

	eh = &exampleHandler{}

	server, err = modbus.NewServer(&modbus.ServerConfiguration{
		URL:		"tcp://0.0.0.0:5502",
		Timeout:	30 * time.Second,
		MaxClients:	5,
	}, eh)

	if err != nil {
		fmt.Printf("failed to create server: %v\n", err)
		os.Exit(1)
	}

	err = server.Start()
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Server started at address: tcp://0.0.0.0:5502\n")

	ticker	= time.NewTicker(1 * time.Second)
	for {
		<-ticker.C

		eh.lock.Lock()
		eh.uptime++
		eh.lock.Unlock()
	}
}

type exampleHandler struct {
	lock 		sync.RWMutex
	uptime 		uint32
	holdingReg1 uint16
}

func (eh *exampleHandler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (eh *exampleHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	if req.UnitId != 1 {
		fmt.Printf("Error: Bad unit id: %v\n", req.UnitId)
		err = modbus.ErrBadUnitId
		return
	}

	if (req.Addr != 101) {
		fmt.Printf("Error: Illegal data address: %v\n", req.Addr)
		err = modbus.ErrIllegalDataAddress
		return
	}

	
	if (req.IsWrite) {
		eh.lock.Lock()
		defer eh.lock.Unlock()
		fmt.Printf("Writing value %v to address %v\n", req.Args[0], req.Addr)
		eh.holdingReg1 = req.Args[0]
	} else {
		eh.lock.RLock()
		defer eh.lock.RUnlock()
		fmt.Printf("Reading address %v\n", req.Addr)
	}

	res = append(res, eh.holdingReg1)
	return
}

func (eh *exampleHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (eh *exampleHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	err = modbus.ErrIllegalFunction
	return
}