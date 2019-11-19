package events

const AgvInfoTopic = "go.micro.topic.agvinfo"

type VehicleType uint32

const (
	_ VehicleType = iota
	VehicleTypeI
	VehicleTypeII
	VehicleTypeIII
	VehicleTypeIV
	VehicleTypeV
	VehicleTypeVI
	VehicleTypeVII
	VehicleTypeVIII
)

type AgvInfo struct {
	AgvID             int
	Type              VehicleType
	Point             int32
	Segment           int32
	X                 int32
	Y                 int32
	Angle             int32
	QuantityOfBattery uint16
	IsCharging        bool
}
