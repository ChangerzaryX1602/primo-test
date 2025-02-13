package infrastructure

import (
	"fmt"
	"log"
	"net"
	"test/api/pb"
	"test/api/server"
	"test/internal/repository"
	"test/internal/usecase"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Resources struct {
	*gorm.DB
}

func NewServer(version, buildTag, runEnv string) (servers *Resources, err error) {
	// mainDbConn, err := ConnectDb(DbConfig{
	// 	DbDriver: "postgres",
	// 	DbName:   viper.GetString("db.postgres.db_name"),
	// 	Host:     viper.GetString("db.postgres.host"),
	// 	Username: viper.GetString("db.postgres.username"),
	// 	Password: viper.GetString("db.postgres.password"),
	// 	Port:     viper.GetInt("db.postgres.port"),
	// 	Timezone: "Asia/Bangkok",
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// AutoMigrate(mainDbConn)
	return nil, nil
}
func (s *Resources) Run() {
	// Start GRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("grpc.port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	userRepository := repository.NewTestRepository(nil)
	userUsecase := usecase.NewTestUsecase(userRepository)
	userServer := server.NewTestServer(userUsecase)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(LogResponsesInterceptor))
	pb.RegisterTestServiceServer(grpcServer, userServer)
	log.Println("Server started on port :", viper.GetInt("grpc.port"))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
