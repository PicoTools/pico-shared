.PHONY: pb
pb:
	@echo "common.proto"
	@protoc -I $(PWD)/proto/schema --go-grpc_out=paths=source_relative:$(PWD)/proto/gen --go_out=paths=source_relative:$(PWD)/proto/gen common/v1/common.proto
	@echo "listener.proto"
	@protoc -I $(PWD)/proto/schema --go-grpc_out=paths=source_relative:$(PWD)/proto/gen --go_out=paths=source_relative:$(PWD)/proto/gen listener/v1/listener.proto
	@echo "operator.proto"
	@protoc -I $(PWD)/proto/schema --go-grpc_out=paths=source_relative:$(PWD)/proto/gen --go_out=paths=source_relative:$(PWD)/proto/gen operator/v1/operator.proto
	@echo "management.proto"
	@protoc -I $(PWD)/proto/schema --go-grpc_out=paths=source_relative:$(PWD)/proto/gen --go_out=paths=source_relative:$(PWD)/proto/gen management/v1/management.proto
