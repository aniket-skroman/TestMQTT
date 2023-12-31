package batterydataprocess

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

var temp_collection = dbconfig.GetCollection(dbconfig.DB, "battery_temp")

func updateTempCollection(hardwareData dto.BatteryData) error {
	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		bson.E{Key: "bms_id", Value: hardwareData.BmsID},
	}

	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "type", Value: hardwareData.Type},
			bson.E{Key: "bms_id", Value: hardwareData.BmsID},
			bson.E{Key: "gsm_signal_strength", Value: hardwareData.GsmSignalStrength},
			bson.E{Key: "gps_signal_strength", Value: hardwareData.GpsSignalStrength},
			bson.E{Key: "gps_satellite_in_view_count", Value: hardwareData.GpsSatelliteInViewCount},
			bson.E{Key: "gnss_satellite_used_count", Value: hardwareData.GnssSatelliteUsedCount},
			bson.E{Key: "gps_status", Value: hardwareData.GpsStatus},
			bson.E{Key: "location_longitude", Value: hardwareData.LocationLongitude},
			bson.E{Key: "location_latitude", Value: hardwareData.LocationLatitude},
			bson.E{Key: "location_speed", Value: hardwareData.LocationSpeed},
			bson.E{Key: "location_angle", Value: hardwareData.LocationAngle},
			bson.E{Key: "iot_temperature", Value: hardwareData.IotTemperature},
			bson.E{Key: "gprs_total_data_used", Value: hardwareData.GprsTotalDataUsed},
			bson.E{Key: "software_version", Value: hardwareData.SoftwareVersion},
			bson.E{Key: "bms_software_version", Value: hardwareData.BmsSoftwareVersion},
			bson.E{Key: "iccid", Value: hardwareData.Iccid},
			bson.E{Key: "imei", Value: hardwareData.Imei},
			bson.E{Key: "gprs_apn", Value: hardwareData.GprsApn},
			bson.E{Key: "is_first_fill", Value: hardwareData.IsFirstFill},
			bson.E{Key: "battery_voltage", Value: hardwareData.BatteryVoltage},
			bson.E{Key: "battery_current", Value: hardwareData.BatteryCurrent},
			bson.E{Key: "battery_soc", Value: hardwareData.BatterySoc},
			bson.E{Key: "battery_temperature", Value: hardwareData.BatteryTemperature},
			bson.E{Key: "battery_remaining_capacity", Value: hardwareData.BatteryRemainingCapacity},
			bson.E{Key: "battery_full_charge_capacity", Value: hardwareData.BatteryFullChargeCapacity},
			bson.E{Key: "battery_cycle_count", Value: hardwareData.BatteryCycleCount},
			bson.E{Key: "battery_rated_capacity", Value: hardwareData.BatteryRatedCapacity},
			bson.E{Key: "battery_rated_voltage", Value: hardwareData.BatteryRatedVoltage},
			bson.E{Key: "battery_version", Value: hardwareData.BatteryVersion},
			bson.E{Key: "battery_manufacture_date", Value: hardwareData.BatteryManufactureDate},
			bson.E{Key: "battery_manufacture_name", Value: hardwareData.BatteryManufactureName},
			bson.E{Key: "battery_name", Value: hardwareData.BatteryName},
			bson.E{Key: "battery_chem_id", Value: hardwareData.BatteryChemID},
			bson.E{Key: "bms_bar_code", Value: hardwareData.BmsBarCode},
			bson.E{Key: "is_second_fill", Value: true},
			bson.E{Key: "cell_voltage_list_0", Value: hardwareData.CellVoltageList0},
			bson.E{Key: "cell_voltage_list_1", Value: hardwareData.CellVoltageList1},
			bson.E{Key: "history", Value: hardwareData.History},
			bson.E{Key: "error_count", Value: hardwareData.ErrorCount},
			bson.E{Key: "status", Value: hardwareData.Status},
			bson.E{Key: "is_third_fill", Value: true},
			bson.E{Key: "created_at", Value: primitive.NewDateTimeFromTime(time.Now())},
			bson.E{Key: "updated_at", Value: primitive.NewDateTimeFromTime(time.Now())},
		}},
	}

	_, err := temp_collection.UpdateOne(context.Background(), filter, update, opts)

	return err
}

func ProcessBatteryData(client mqtt.Client, msg mqtt.Message) {
	containData := string(msg.Payload())

	msgData := dto.BatteryData{}

	err := json.Unmarshal([]byte(containData), &msgData)

	response_data := dto.DataReceiveResponse{}

	if err != nil {
		response_data.BmsId = ""
		response_data.Status = false
		response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
		notifyToHardwareFromSever(client, response_data)
		return
	}

	response_data.BmsId = msgData.BmsID
	response_data.Status = true
	response_data.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	notifyToHardwareFromSever(client, response_data)

	updateTempCollection(msgData)
}

func notifyToHardwareFromSever(client mqtt.Client, response_data dto.DataReceiveResponse) {
	json_data, _ := json.Marshal(response_data)
	if token := client.Publish("spiro/battery/SPIROBAT-CRFXUQOD/ack", 0, false, json_data); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to ACK to battery data : ", token.Error())
	}
}
