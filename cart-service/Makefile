build:
	go build -o bin/micro


evans:
	evans --host localhost --port 9091 -r repl
install:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

run: build
	./bin/micro

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

proto_1:
	protoc -I . \
    --go_out=. --go_opt paths=source_relative \
    --go-grpc_out=. --go-grpc_opt paths=source_relative \
    proto/service.proto

proto_2:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger

proto_3:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

proto_jalan:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

proto_jalan_tes:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

docker-update:
	docker rmi farhandwian/cart-service:latest
	docker build -t farhandwian/cart-service:latest .

docker-run:
	echo "Starting core services..."
	docker-compose -f ./docker-compose-core-service.yml up -d zookeeper order-db auth-db redis kafka discovery-server
	echo "Starting and registering API Gateway..."
	sleep 240
	docker-compose -f ./docker-compose-core-service.yml up -d api-gateway
	sleep 240
	echo "Starting and registering all microservices..."
	docker-compose -f ./docker-compose-core-service.yml up -d auth-service product-service order-service payment-service
	echo "All services started."


docker-run-dev:
	echo "Starting core services..."
	docker-compose -f ./docker-compose-core-service.yml up -d redis discovery-server
	echo "Starting and registering API Gateway..."
	sleep 240
	docker-compose -f ./docker-compose-core-service.yml up -d api-gateway
	sleep 240
	echo "Starting and registering all microservices..."
	docker-compose -f ./docker-compose-core-service.yml up -d cart-service 
	echo "All services started."
	

docker-run-specific-container:
	docker-compose -f docker-compose-core-service.yml up -d auth-db


docker-build:
	docker build -t farhandwian/cart-service:latest .


docker-push-hub:
	docker tag farhandwian/cart-service:latest alciquairev/farhandwian:latest
	docker push alciquairev/farhandwian:latest	

start:
	cd cmd
	cd grpc
	go run .	

.PHONY: proto