package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.

// Структура конфигурации

type Config struct {
	Logger   LoggerConf   `yaml:"logger"`
	Database DatabaseConf `yaml:"database"`
}

type DatabaseConf struct {
	InMemory bool   `yaml:"in_memory"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

type LoggerConf struct {
	Enabled bool   `yaml:"enabled"`
	Level   string `yaml:"level"`
}

func NewConfig(pathConfigFile string) Config {
	// Открываем файл
	file, err := os.Open(pathConfigFile)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Ошибка декодирования YAML: %v", err)
	}

	return config
}
