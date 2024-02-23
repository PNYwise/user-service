package internal

import (
	"github.com/PNYwise/user-service/internal/handler"
	"github.com/PNYwise/user-service/internal/repository"
	"github.com/PNYwise/user-service/internal/service"
	user_service "github.com/PNYwise/user-service/proto"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func InitGrpc(srv grpc.ServiceRegistrar, db *pgx.Conn) {
	userRepostory := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepostory)
	userHandlers := handler.NewUserHandler(userService)
	user_service.RegisterUserServer(srv, userHandlers)
}
