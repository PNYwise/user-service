package internal

import (
	"github.com/IBM/sarama"
	"github.com/PNYwise/user-service/internal/domain"
	"github.com/PNYwise/user-service/internal/handler"
	"github.com/PNYwise/user-service/internal/repository"
	"github.com/PNYwise/user-service/internal/service"
	user_service "github.com/PNYwise/user-service/proto"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func InitGrpc(srv grpc.ServiceRegistrar, db *pgx.Conn, producer sarama.SyncProducer, extConf *domain.ExtConf) {
	userMessagingRepo := repository.NewPostMessagingRepository(producer, extConf)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, userMessagingRepo)
	userHandlers := handler.NewUserHandler(userService)
	user_service.RegisterUserServer(srv, userHandlers)
}
