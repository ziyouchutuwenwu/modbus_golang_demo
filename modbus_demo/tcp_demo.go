package modbus_demo

import (
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"os"
	"time"
)

func TcpMasterDemo(){
	handler := modbus.NewTCPClientHandler("192.168.88.148:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)

	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)

	dataToWrite := []byte{ 0, 11, 0, 22, 0, 33, 0, 44, 0, 55}
	client.WriteMultipleRegisters(0, uint16(len(dataToWrite))/2, dataToWrite)

	results, err := client.ReadHoldingRegisters(0, uint16(len(dataToWrite))/2)

	fmt.Println(results, err)
}

