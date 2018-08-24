build:
		protoc -I $(GOPATH)/src/github.com/amiralii/shipping/user-service/ $(GOPATH)/src/github.com/amiralii/shipping/user-service/proto/user/user.proto --go_out=plugins=micro:$(GOPATH)/src/github.com/amiralii/shipping/user-service/

		docker build -t user-service .
 run:
		docker run -p 8585:50051	\
			-e MICRO_SERVER_ADDRESS=:50051	\
			-e MICRO_REGISTRY=mdns user-service

