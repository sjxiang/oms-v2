
output_dir = './internal/common/pb'

gen_proto: 
	echo "proto 目录下的 proto 文件代码生成 ... "
	protoc \
		--proto_path proto \
		--go_out $(output_dir) --go_opt paths=source_relative \
		--go-grpc_out $(output_dir) --go-grpc_opt paths=source_relative,require_unimplemented_servers=false \
		proto/*.proto
	echo "生成完成"


.PHONY: gen_proto


# 目录 './pb'
# 路径 './pb/*.go'
