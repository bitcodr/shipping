build: 
		protoc -I $(GOPATH)/src/github.com/amiralii/shipping/vessel-service/ $(GOPATH)/src/github.com/amiralii/shipping/vessel-service/proto/vessel/vessel.proto --go_out=plugins=micro:$(GOPATH)/src/github.com/amiralii/shipping/vessel-service/

		docker build -t vessel-service .
run:
		docker run -p 8686:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns vessel-service




