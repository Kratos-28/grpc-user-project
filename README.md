# gRPC User Service

This is a sample gRPC service for managing user details with search functionality.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Build and Run](#build-and-run)
3. [Accessing gRPC Endpoints](#accessing-grpc-endpoints)

## Prerequisites
- Go 1.18 or higher installed on your machine.
- Docker and Docker Compose installed (optional, for containerization).

## Build and Run

### Using Go
1. Clone this repository to your local machine.
2. Navigate to the server directory: `cd grpc-userservice/server`.
3. Run the following command to build the application:
   ```bash
   go build -o main .
4. Run the executable: ./main

##Accessing gRPC Endpoints

Once the server is running, you can access the gRPC service endpoints using any gRPC client. By default, the server listens on port 50051.
To run server
   cd cd grpc-userservice/server
    go run main.go
To run the client:
  cd grpc-userservice/client
  go run client.go

    


