package vehicledataprocess

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	dbconfig "github.com/aniket0951/testmqtt/db-config"
	"github.com/aniket0951/testmqtt/dto"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	temp_collection     = dbconfig.GetCollection(dbconfig.DB, "vehicle_hardware")
	vehicles_collection = dbconfig.GetCollection(dbconfig.DB, "vehicles")

	VEHICLE_INFO_ACK     = "spiro/vehicle/SPIRO-TEST-1/boot/ack"
	VEHICLE_HARDWARE_ACK = "spiro/vehicle/SPIRO-TEST-1/ack"
)

func updateTempCollection(hardwareData dto.VehicleDTO) error {
	opts := options.FindOneAndReplace().SetUpsert(true)

	filter := bson.D{
		bson.E{Key: "uid", Value: hardwareData.UID},
	}

	hardwareData.Created_at = primitive.NewDateTimeFromTime(time.Now())
	hardwareData.Updated_at = primitive.NewDateTimeFromTime(time.Now())

	result := temp_collection.FindOneAndReplace(context.Background(), filter, &hardwareData, opts)

	return result.Err()
}

func updateVehicleInfo(vehicleInfo dto.VehicleInfo) error {
	opts := options.FindOneAndReplace().SetUpsert(true)

	filter := bson.D{
		bson.E{Key: "vehiclename", Value: vehicleInfo.Vehiclename},
	}

	vehicleInfo.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	vehicleInfo.Created_at = primitive.NewDateTimeFromTime(time.Now())
	vehicleInfo.Updated_at = primitive.NewDateTimeFromTime(time.Now())

	result := vehicles_collection.FindOneAndReplace(context.Background(), filter, &vehicleInfo, opts)
	return result.Err()
}

func ProcessVehicleData(client mqtt.Client, msg mqtt.Message) {
	containData := string(msg.Payload())

	msgData := dto.VehicleDTO{}

	err := json.Unmarshal([]byte(containData), &msgData)

	response_data := dto.DataReceiveForVehicelResponse{}

	if err != nil {
		response_data.UID = ""
		response_data.Status = false
		response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
		notifyToHardwareFromSever(client, response_data)
		return
	}

	response_data.UID = msgData.UID
	response_data.Status = true
	response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	notifyToHardwareFromSever(client, response_data)

	updateTempCollection(msgData)
}

func ProcessVehicleInfoData(client mqtt.Client, msg mqtt.Message) {
	containData := string(msg.Payload())

	vehicleInfoData := dto.VehicleInfo{}

	err := json.Unmarshal([]byte(containData), &vehicleInfoData)
	response_data := dto.DataReceiveForVehicelResponse{}

	if err != nil {
		response_data.UID = ""
		response_data.Status = false
		response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
		notifyToVehicleInfoFromSever(client, response_data)
		return
	}

	response_data.UID = vehicleInfoData.Vehiclename
	response_data.Status = true
	response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	notifyToVehicleInfoFromSever(client, response_data)

	updateVehicleInfo(vehicleInfoData)

}

func notifyToHardwareFromSever(client mqtt.Client, response_data dto.DataReceiveForVehicelResponse) {
	json_data, _ := json.Marshal(response_data)
	fmt.Println("Ack from vehicle : ", response_data)
	if token := client.Publish(VEHICLE_HARDWARE_ACK, 0, false, json_data); token.Wait() && token.Error() != nil {
		fmt.Println("failed to ACK for vehicle hardware data : ", token.Error())
	}

}

func notifyToVehicleInfoFromSever(client mqtt.Client, response_data dto.DataReceiveForVehicelResponse) {
	json_data, _ := json.Marshal(response_data)
	fmt.Println("Ack from vehicle info : ", response_data)
	if token := client.Publish(VEHICLE_INFO_ACK, 0, false, json_data); token.Wait() && token.Error() != nil {
		fmt.Println("failed to ACK for vehicle info data : ", token.Error())
	}

}
