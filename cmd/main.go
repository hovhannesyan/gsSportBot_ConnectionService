package main

import (
	"fmt"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/config"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/db"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/pb"
	"github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/services"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var f *os.File

func init() {
	var err error
	t := time.Now().Format("2006-01-02")

	logFile := fmt.Sprintf("./log/log_%s.txt", t)
	f, err = os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	logrus.SetOutput(f)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := config.LoadConfig(); err != nil {
		logrus.Fatalln(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalln("error loading.env file: %s", err.Error())
	}
}

func main() {
	defer f.Close()

	dbHandler := db.Init(os.Getenv("REDIS_URL"))

	lis, err := net.Listen("tcp", viper.GetString("port"))

	if err != nil {
		logrus.Fatalln("failed to listing", err.Error())
	}

	logrus.Println("Connection on", viper.GetString("port"))

	s := services.Server{
		DbHandler: dbHandler,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterConnectionServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalln("failed to serve", err.Error())
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Shutting down server...")

	grpcServer.GracefulStop()

	logrus.Println("Server exiting")
}
