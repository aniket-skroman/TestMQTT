package dto

type InfoDTO struct {
	Dt     int    `json:"dt"`
	Txn    string `json:"txn"`
	Msgkey int    `json:"msgkey"`
	Msgid  int    `json:"msgid"`
	Cmdkey string `json:"cmdkey"`
	Cmdval string `json:"cmdval"`
}

type GPSDTO struct {
	Fix   string    `json:"fix"`
	Loc   []float64 `json:"loc"`
	Speed int       `json:"speed"`
	Sat   int       `json:"sat"`
	Alt   int       `json:"alt"`
	Dir   int       `json:"dir"`
	Odo   int       `json:"odo"`
}

type IoDTO struct {
	Box    int   `json:"box"`
	Ign    int   `json:"ign"`
	Gpi    int   `json:"gpi"`
	Status int   `json:"status"`
	Analog []int `json:"analog"`
}

type PWRDTO struct {
	Main  int     `json:"main"`
	Batt  int     `json:"batt"`
	Volt  int     `json:"volt"`
	Mvolt float64 `json:"mvolt"`
}

type DBGDTO struct {
	Status []int    `json:"status"`
	Ver    []string `json:"ver"`
	Lib    string   `json:"lib"`
}

type VehicleDTO struct {
	UID  string  `json:"uid"`
	Info InfoDTO `json:"info"`
	Gps  GPSDTO  `json:"gps"`
	Io   IoDTO   `json:"io"`
	Pwr  PWRDTO  `json:"pwr"`
	Dbg  DBGDTO  `json:"dbg"`
}
