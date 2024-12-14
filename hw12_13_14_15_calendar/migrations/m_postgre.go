package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/configs"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	"io/ioutil"
	"strings"
)

func EnsureDatabase(ctx context.Context, logg logger.Log, dbConf configs.DatabaseConf, dsn string) error {

	// Если работа с памятью миграция не нужна
	if dbConf.InMemory {
		return nil
	}

	// Подключение без указания имени базы данных
	adminDSN := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port)

	db, err := sql.Open("postgres", adminDSN)
	if err != nil {
		msg := fmt.Sprintf("failed to connect to database server: %s\n", err)
		logg.Error(msg)
		return fmt.Errorf(msg)
	}
	defer db.Close()

	// Проверяем, существует ли база данных
	var exists bool
	err = db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbConf.Name).Scan(&exists)
	if err != nil {
		msg := fmt.Sprintf("failed to check database existence: %s\n", err)
		logg.Error(msg)
		return fmt.Errorf(msg)
	}

	if !exists {
		// Создаём базу данных
		_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", dbConf.Name))
		if err != nil {
			msg := fmt.Sprintf("failed to create database: %s\n", err)
			logg.Error(msg)
			return fmt.Errorf(msg)
		}
		logg.Info("Database created successfully")
	}

	// Выполняем миграции
	if err = runMigrations(dsn, dbConf.FileMigrations, logg); err != nil {

		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// runMigrations выполняет миграции из файла
func runMigrations(dsn, migrationsFile string, logg logger.Log) error {
	// Подключение к базе данных
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Чтение содержимого файла миграций
	migrationSQL, err := ioutil.ReadFile(migrationsFile)
	if err != nil {
		return fmt.Errorf("failed to read migrations file: %w", err)
	}

	// Разделение файла на отдельные SQL-команды
	migrationCommands := strings.Split(string(migrationSQL), ";")
	for _, command := range migrationCommands {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		// Выполнение каждой SQL-команды
		_, err = db.Exec(command)
		if err != nil {
			return fmt.Errorf("failed to execute migration command (%s): %w", command, err)
		}
	}

	logg.Info("All migrations executed successfully")
	return nil
}
