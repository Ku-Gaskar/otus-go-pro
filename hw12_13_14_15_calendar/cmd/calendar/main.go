package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/configs"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/migrations"
	"os"
	"os/signal"
	"syscall"
	//"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/app"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/server/http"
	storage2 "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

func main() {

	flag.Parse()

	//switch flag.Arg(0) {
	//case "version":
	//	printVersion()
	//	return
	//case "migrate":
	//	migrations.Migrations()
	//	return
	//}
	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := configs.NewConfig(configFile)
	logg := logger.New(config.Logger.Enabled, config.Logger.Level, "")

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Name, config.Database.Host, config.Database.Port)

	// Проверка наличия базы данных и миграция в случае отсутствия
	if err := migrations.EnsureDatabase(ctx, logg, config.Database, dsn); err != nil {
		fmt.Printf("Failed to ensure database: %v\n", err)
		logg.Error("Failed to ensure database:", err)
		return
	}

	// Инициализация хранилища
	storage, err := storage2.NewStorage(ctx, config.Database.InMemory, config.Database.DriverName, dsn)
	if err != nil {
		fmt.Printf("Error initializing storage: %v\n", err)
		logg.Error("Error initializing storage:", err)
		return
	}
	defer func() {
		if err := storage.Close(ctx); err != nil {
			fmt.Printf("Error closing storage: %v\n", err)
		}
	}()

	//calendar := app.New(logg, storage)

	// Создание и запуск HTTP-сервера
	//server := internalhttp.NewServer(config.ServerHttp.Host, config.ServerHttp.Port, calendar, logg)
	server := internalhttp.NewServer(config.Server.Http.Host, config.Server.Http.Port, logg)

	go func() {
		<-ctx.Done()

		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
