package config

import (
	"context"
	"encoding/json"
	"log"

	"github.com/PNYwise/user-service/internal/domain"
	user_service "github.com/PNYwise/user-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

var grpcConn *grpc.ClientConn

func getConfigConn(conf *viper.Viper) *grpc.ClientConn {
	// Dial the gRPC server
	conn, err := grpc.Dial(
		conf.GetString("config-service.host")+":"+conf.GetString("config-service.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Config Service gRPC server: %v", err)
	}
	log.Println("Connected to Config Service gRPC server")
	grpcConn = conn
	return grpcConn
}

func ConfigFromGrpcServer(ctx context.Context, conf *viper.Viper) *domain.ExtConf {
	conn := getConfigConn(conf)
	client := user_service.NewConfigClient(conn)
	// Add metadata to the context
	response, err := client.Get(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}
	grpcConn.Close()
	extConf, err := parseConfigResponse(response)
	if err != nil {
		log.Fatalf("Error unmarshaling configuration: %v", err)
	}

	return extConf
}

func parseConfigResponse(response *structpb.Value) (*domain.ExtConf, error) {
	extConf := &domain.ExtConf{}
	if stringVal, ok := response.Kind.(*structpb.Value_StringValue); ok {
		err := json.Unmarshal([]byte(stringVal.StringValue), extConf)
		return extConf, err
	}
	return nil, nil
}
