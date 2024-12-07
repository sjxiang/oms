

gen_proto:
	echo "proto 代码生成中..."
	protoc \
		--proto_path=proto \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		proto/*.proto
  

.PHONY: gen_proto


