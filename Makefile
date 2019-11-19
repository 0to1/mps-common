
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/role/role.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/warehouse/warehouse.proto
