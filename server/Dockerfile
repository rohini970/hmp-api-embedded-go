# GRPC-demo
# build command : docker build . -t grpc-demo/client
# run command : docker run -it grpc-demo/client
FROM golang:latest

# Install grpc
RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

# Copy the grpc proto file and generate the go module
RUN mkdir -p /go/src/hmp-api-embedded-go
COPY homegateway.proto /go/src/hmp-api-embedded-go
RUN cd /go/src && \
    protoc -I hmp-api-embedded-go/ hmp-api-embedded-go/homegateway.proto --go_out=plugins=grpc:hmp-api-embedded-go

# Copy the project to be executed
RUN mkdir -p /go/src/grpc-client 
COPY server.go /go/src/grpc-client

ENTRYPOINT cd /go/src/grpc-client && go run server.go
