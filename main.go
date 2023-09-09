package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	batterydataprocess "github.com/aniket0951/testmqtt/battery-data-process"
	vehicledataprocess "github.com/aniket0951/testmqtt/vehicle-data-process"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	batteryHardwareTopic = "spiro/battery/in/hw"
	vehicleHardwareTopic = "spiro/vehicle/in/hw"
	vehicleInfoTopic     = "spiro/vehicle/in/hw/boot"
)

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if msg.Topic() == batteryHardwareTopic {
		fmt.Println("Processing a battery data \n", msg.Topic())
		batterydataprocess.ProcessBatteryData(client, msg)
	} else if msg.Topic() == vehicleHardwareTopic {
		fmt.Println("processing a vehicle data \n", msg.Topic())
		vehicledataprocess.ProcessVehicleData(client, msg)
	} else if msg.Topic() == vehicleInfoTopic {
		fmt.Println("Processing a vehicle info data \n", msg.Topic())
		vehicledataprocess.ProcessVehicleInfoData(client, msg)
	}
}

func main() {
	//defer dbconfig.CloseClientDB()

	opts := mqtt.NewClientOptions()
	//opts.AddBroker("tcp://test.mosquitto.org:1883")
	opts.AddBroker("mqtt://v-tro.in:1883")
	opts.SetClientID("go_mqtt_client")
	opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to MQTT broker:", token.Error())
		os.Exit(1)
	}

	fmt.Println("Connected to MQTT broker")

	token := client.Subscribe(batteryHardwareTopic, 0, nil)
	token2 := client.Subscribe(vehicleHardwareTopic, 0, nil)
	token3 := client.Subscribe(vehicleInfoTopic, 0, nil)

	t1 := token.Wait()
	t2 := token2.Wait()
	t3 := token3.Wait()

	fmt.Println("Subscribed to topics ", t1, t2, t3)
	//go printSensorDataLoop()

	// go func() {
	// 	publish(client)
	// 	publishDummyData(client)
	// }()

	// go publish(client)
	// go publishDummyData(client)

	// Keep the application running
	for {
		time.Sleep(time.Second)
	}
}

type SensorData struct {
	MSG string `json:"msg"`
}

var sensorDataList []SensorData

func printSensorDataLoop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		fmt.Println("Current sensor data:")
		for _, data := range sensorDataList {
			fmt.Printf("Temperature: %s\n", data.MSG)
		}
	}
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)

		token := client.Publish(batteryHardwareTopic, 0, false, text)
		token.Wait()
		time.Sleep(2 * time.Second)
		fmt.Println(i+1, " loop has been done")
	}
}

func publishDummyData(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)

		token := client.Publish(vehicleHardwareTopic, 0, false, text)
		token.Wait()
		time.Sleep(2 * time.Second)
		fmt.Println(i+1, " loop has been done")
	}
}
