create_timeline:
	protoc -I ./grpc-gateway/pkg/proto \
	--go_out ./gen/proto/ --go_opt paths=source_relative \
	--go-grpc_out ./gen/proto/ --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen/proto/ --grpc-gateway_opt paths=source_relative \
	grpc-gateway/pkg/proto/timeline/timeline.proto



clean_timeline:
	rm -rf gen/proto/timeline