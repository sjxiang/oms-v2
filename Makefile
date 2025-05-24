

# ✅ 正确写法（去掉单引号）

# 目录
input_dir  = proto
output_dir = internal/common/pb

# 颜色转义符
red     = \033[1;31m
white   = \033[1;37m
yellow  = \033[1;33m
nocolor = \033[0m


gen-proto:
	@echo "$(red)正在清理输出目录...$(nocolor)"
	rm -rf $(output_dir)/*.go
	@echo "$(yellow)正在从 proto 文件生成 Go 代码和 gRPC 服务代码...$(nocolor)"
	@echo "$(white)$(input_dir) => $(output_dir)$(nocolor)"
	protoc --proto_path=$(input_dir) \
		--go_out=$(output_dir) --go_opt=paths=source_relative \
		--go-grpc_out=$(output_dir) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		$(input_dir)/*.proto



container-up:
	@echo "$(white)容器上线...$(nocolor)"
	docker-compose up -d

container-down:
	@echo "$(white)容器下线...$(nocolor)"
	docker-compose -f ./docker-compose.yml down
	

container-stop:
	@echo "$(white)正在暂停容器...$(nocolor)"
	@docker-compose stop


container-start:
	@echo "$(white)正在开启容器...$(nocolor)"
	@docker-compose start


.PHONY: gen-proto container-up container-down container-stop container-start