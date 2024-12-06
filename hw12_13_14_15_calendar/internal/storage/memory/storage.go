package memorystorage

import (
	"context"
	"fmt"
	"sync"
)

type Storage struct {
	data map[string]string
	mu   sync.RWMutex //nolint:unused
}

//func New() *Storage {
//	return &Storage{}
//}

//// InMemoryStorage реализация для хранения данных в памяти.
//type InMemoryStorage struct {
//	data map[string]string // Пример хранения данных в памяти
//}

func (m *Storage) Close(ctx context.Context) error {
	// Для InMemory ничего не требуется
	fmt.Println("Closing in-memory storage")
	return nil
}

func New() *Storage {
	return &Storage{data: make(map[string]string)}
}

func (m *Storage) Connect(ctx context.Context, driverName string, dsn string) error {
	// Для InMemory ничего не требуется
	fmt.Println("Using in-memory storage")
	return nil
}
