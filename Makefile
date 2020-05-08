
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/role/role.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/config/config.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/area/area.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/lpr/lpr.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/material/material.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/opc/opc.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/agv/agv.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/system/system.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/order/order.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/point/point.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/station/station.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/segment/segment.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/route/route.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/rack/rack.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/racktype/racktype.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/racklot/racklot.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/task/task.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/script/script.proto

	protoc --proto_path=. --micro_out=. --go_out=. proto/log/log.proto
