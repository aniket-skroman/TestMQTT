package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type BatteryTemperature struct {
	PreStart     int `json:"pre_start" bson:"pre_start"`
	Cell1        int `json:"cell_1" bson:"cell_1"`
	Cell2        int `json:"cell_2" bson:"cell_2"`
	Cell3        int `json:"cell_3" bson:"cell_3"`
	ChargeMos    int `json:"charge_mos" bson:"charge_mos"`
	DisChargeMos int `json:"discharge_mos" bson:"discharge_mos"`
	PCB          int `json:"pcb" bson:"pcb"`
	PreCharge    int `json:"pre_charge" bson:"pre_charge"`
}

type BatteryErrorCount struct {
	PreStartFail                          int `json:"pre_start_fail" bson:"pre_start_fail"`
	PrimaryOverDischargeError             int `json:"primary_over_discharge_error" bson:"primary_over_discharge_error"`
	SecoundaryOverDischargeError          int `json:"secondary_over_discharge_error" bson:"secondary_over_discharge_error"`
	ThirdOverCurrentError                 int `json:"third_over_current_error" bson:"third_over_current_error"`
	OverChargeTemperatureError            int `json:"over_charge_temperature_error"`
	OverTemperatureOfPreStartCircuitError int `json:"over_temperature_of_pre_start_circuit_error"`
	ProtectionChipError                   int `json:"protection_chip_error"`
	OverChargeError                       int `json:"over_charge_error"`
}

type BatteryHistory struct {
	MaxDischargeCurrent int `json:"max_discharge_current" bson:"max_discharge_current"`
	MaxTemperature      int `json:"max_temperature" bson:"max_temperature"`
	MinCellVoltage      int `json:"min_cell_voltage" bson:"min_cell_voltage"`
	MinTemperature      int `json:"min_temperature" bson:"min_temperature"`
	MaxCellVoltage      int `json:"max_cell_voltage" bson:"max_cell_voltage"`
	MaxChargeCurrent    int `json:"max_charge_current" bson:"max_charge_current"`
}

type Status struct {
	ChargeMosfetStatus        int `json:"charge_mosfet_status" bson:"charge_mosfet_status"`
	ChargeIdentifyStatus      int `json:"charger_identify_status" bson:"charger_identify_status"`
	CycleCount                int `json:"cycle_count" bson:"cycle_count"`
	DischargeMosfetStatus     int `json:"discharge_mosfet_status" bson:"discharge_mosfet_status"`
	LowestCellTemperature     int `json:"lowest_cell_temperature" bson:"lowest_cell_temperature"`
	MosfetTemperature         int `json:"mosfet_temperature" bson:"mosfet_temperature"`
	PreChargeCircuitStatus    int `json:"pre_charge_circuit_status" bson:"pre_charge_circuit_status"`
	RealLightStatus           int `json:"real_light_status" bson:"real_light_status"`
	SOH                       int `json:"soh" bson:"soh"`
	UsbPowerStatus            int `json:"usb_power_status" bson:"usb_power_status"`
	HighestCellTemperature    int `json:"highest_cell_temperature" bson:"highest_cell_temperature"`
	PreDischargeCircuitStatus int `json:"pre_discharge_circuit_status" bson:"pre_discharge_circuit_status"`
	RealtimeCurrent           int `json:"realtime_current" bson:"realtime_current"`
	ChargerConnectionStatus   int `json:"charger_connection_status" bson:"charger_connection_status"`
	PreStartTemperature       int `json:"pre_start_temperature" bson:"pre_start_temperature"`
	PackVoltage               int `json:"pack_voltage" bson:"pack_voltage"`
	SOC                       int `json:"soc" bson:"soc"`
}

// data receive in this
type BatteryData struct {
	Type                      int                `json:"type"`
	BmsID                     string             `json:"bms_id"`
	GsmSignalStrength         int                `json:"gsm_signal_strength"`
	GpsSignalStrength         int                `json:"gps_signal_strength"`
	GpsSatelliteInViewCount   int                `json:"gps_satellite_in_view_count"`
	GnssSatelliteUsedCount    int                `json:"gnss_satellite_used_count"`
	GpsStatus                 int                `json:"gps_status"`
	LocationLongitude         int                `json:"location_longitude"`
	LocationLatitude          int                `json:"location_latitude"`
	LocationSpeed             int                `json:"location_speed"`
	LocationAngle             int                `json:"location_angle"`
	IotTemperature            int                `json:"iot_temperature"`
	GprsTotalDataUsed         int                `json:"gprs_total_data_used"`
	SoftwareVersion           string             `json:"software_version"`
	BmsSoftwareVersion        string             `json:"bms_software_version"`
	Iccid                     string             `json:"iccid"`
	Imei                      string             `json:"imei"`
	GprsApn                   string             `json:"gprs_apn"`
	IsFirstFill               bool               `json:"is_first_fill" bson:"is_first_fill"`
	BatteryVoltage            int                `json:"battery_voltage"`
	BatteryCurrent            float64            `json:"battery_current"`
	BatterySoc                int                `json:"battery_soc"`
	BatteryTemperature        BatteryTemperature `json:"battery_temperature"`
	BatteryRemainingCapacity  int                `json:"battery_remaining_capacity"`
	BatteryFullChargeCapacity int                `json:"battery_full_charge_capacity"`
	BatteryCycleCount         int                `json:"battery_cycle_count"`
	BatteryRatedCapacity      int                `json:"battery_rated_capacity"`
	BatteryRatedVoltage       int                `json:"battery_rated_voltage"`
	BatteryVersion            string             `json:"battery_version"`
	BatteryManufactureDate    interface{}        `json:"battery_manufacture_date"`
	BatteryManufactureName    string             `json:"battery_manufacture_name"`
	BatteryName               string             `json:"battery_name"`
	BatteryChemID             string             `json:"battery_chem_id"`
	BmsBarCode                string             `json:"bms_bar_code"`
	IsSecondFill              bool               `json:"is_second_fill" bson:"is_second_fill"`
	CellVoltageList0          interface{}        `json:"cell_voltage_list_0"`
	CellVoltageList1          interface{}        `json:"cell_voltage_list_1"`
	History                   BatteryHistory     `json:"history"`
	ErrorCount                BatteryErrorCount  `json:"error_count"`
	Status                    Status             `json:"status"`
	IsThirdFill               bool               `json:"is_third_fill" bson:"is_third_fill"`
}

// send a response to hardwareData
type DataReceiveResponse struct {
	BmsId     string             `json:"bms_id"`
	Status    bool               `json:"status"`
	TimeStamp primitive.DateTime `json:"time_stamp"`
}
