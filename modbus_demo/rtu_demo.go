package modbus_demo

import (
	"fmt"
	"github.com/goburrow/modbus"
	"time"
)

func RtuMasterDemo(){
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 120 * time.Second

	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)

	dataToWrite := []byte{ 0, 11, 0, 22, 0, 33, 0, 44, 0, 55}
	client.WriteMultipleRegisters(0, uint16(len(dataToWrite))/2, dataToWrite)
	//
	results, err := client.ReadHoldingRegisters(0, uint16(len(dataToWrite))/2)

	fmt.Println(results, err)
}
