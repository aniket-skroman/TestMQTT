package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type InfoDTO struct {
	Dt     int    `json:"dt" bson:"dt"`
	Txn    string `json:"txn" bson:"txn"`
	Msgkey int    `json:"msgkey" bson:"msgkey"`
	Msgid  int    `json:"msgid" bson:"msgid"`
	Cmdkey string `json:"cmdkey" bson:"cmdkey"`
	Cmdval string `json:"cmdval" bson:"cmdval"`
}

type GPSDTO struct {
	Fix   string    `json:"fix" bson:"fix"`
	Loc   []float64 `json:"loc" bson:"loc"`
	Speed int       `json:"speed" bson:"speed"`
	Sat   int       `json:"sat" bson:"sat"`
	Alt   int       `json:"alt" bson:"alt"`
	Dir   int       `json:"dir" bson:"dir"`
	Odo   int       `json:"odo" bson:"odo"`
}

type IoDTO struct {
	Box    int   `json:"box" bson:"box"`
	Ign    int   `json:"ign" bson:"ign"`
	Gpi    int   `json:"gpi" bson:"gpi"`
	Status int   `json:"status" bson:"status"`
	Analog []int `json:"analog" bson:"analog"`
}

type PWRDTO struct {
	Main  int     `json:"main" bson:"main"`
	Batt  int     `json:"batt" bson:"batt"`
	Volt  int     `json:"volt" bson:"volt"`
	Mvolt float64 `json:"mvolt" bson:"mvolt"`
}

type DBGDTO struct {
	Status []int    `json:"status" bson:"status"`
	Ver    []string `json:"ver" bson:"ver"`
	Lib    string   `json:"lib" bson:"lib"`
}

type VehicleDTO struct {
	UID        string             `json:"uid" bson:"uid"`
	Info       InfoDTO            `json:"info" bson:"info"`
	Gps        GPSDTO             `json:"gps" bson:"gps"`
	Io         IoDTO              `json:"io" bson:"io"`
	Pwr        PWRDTO             `json:"pwr" bson:"pwr"`
	Dbg        DBGDTO             `json:"dbg" bson:"dbg"`
	Created_at primitive.DateTime `json:"created_at" bson:"created_at"`
	Updated_at primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
