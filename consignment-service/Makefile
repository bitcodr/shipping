build:
		protoc -I $(GOPATH)/src/github.com/amiralii/shipping/consignment-service/ $(GOPATH)/src/github.com/amiralii/shipping/consignment-service/proto/consignment/consignment.proto --go_out=plugins=micro:$(GOPATH)/src/github.com/amiralii/shipping/consignment-service/

		docker build -t consignment-service .
 run:
		docker run -p 8787:50051	\
			-e MICRO_SERVER_ADDRESS=:50051	\
			-e MICRO_REGISTRY=mdns consignment-service



# go get -u github.com/golang/protobuf/protoc-gen-go
# brew install grpc
# brew install protobuf	


# 8787:50051  firt part port using on host and second part is internal port