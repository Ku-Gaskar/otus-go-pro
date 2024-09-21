package hw04lrucache

type Key string

var cache = make(map[*ListItem]Key)

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok { // Если элемент есть в словаре
		item.Value = value        // Обновили значение
		c.queue.MoveToFront(item) // Переместили в начало очереди
		return true
	}

	if c.queue.Len() >= c.capacity { // Елемента нет в словаре и длинна очереди больше вместимости кэша
		last := c.queue.Back()       // Получаем последний элемент очереди
		c.queue.Remove(last)         // Удаляем последний элемент из очереди
		delete(c.items, cache[last]) // Удаляем последний элемент из мапы
		delete(cache, last)
	}

	item := c.queue.PushFront(key) // Перемещаем в начало очереди
	item.Value = value
	cache[item] = key   // Обновляем значение
	c.items[key] = item // Сохраняем в мапу
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok { // Элемент присутсвует в словаре
		c.queue.MoveToFront(item) // Перемещаем элемент в начало очереди
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	cache = make(map[*ListItem]Key)
}
