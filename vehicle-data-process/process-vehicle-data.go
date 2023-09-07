package vehicledataprocess

import (
	"context"
	"encoding/json"
	"time"

	dbconfig "github.com/aniket0951/testmqtt/db-config"
	"github.com/aniket0951/testmqtt/dto"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var temp_collection = dbconfig.GetCollection(dbconfig.DB, "vehicle_hardware")

func updateTempCollection(hardwareData dto.VehicleDTO) error {
	opts := options.FindOneAndReplace().SetUpsert(true)

	filter := bson.D{
		bson.E{Key: "uid", Value: hardwareData.UID},
	}

	result := temp_collection.FindOneAndReplace(context.Background(), filter, &hardwareData, opts)

	return result.Err()
}

func ProcessVehicleData(client mqtt.Client, msg mqtt.Message) {
	containData := string(msg.Payload())

	msgData := dto.VehicleDTO{}

	err := json.Unmarshal([]byte(containData), &msgData)

	response_data := dto.DataReceiveResponse{}

	if err != nil {
		response_data.BmsId = ""
		response_data.Status = false
		response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
		notifyToHardwareFromSever(client, response_data)
		return
	}

	response_data.BmsId = msgData.UID
	response_data.Status = true
	response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	notifyToHardwareFromSever(client, response_data)

	updateTempCollection(msgData)
}

func notifyToHardwareFromSever(client mqtt.Client, response_data dto.DataReceiveResponse) {
	json_data, _ := json.Marshal(response_data)
	token := client.Publish("spiro/vehicle/SPIROVEH-KRZDEFAM/ack", 0, false, json_data)
	token.Wait()
}
