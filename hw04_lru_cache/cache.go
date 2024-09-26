package hw04lrucache

type Key string

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

type cacheElement struct {
	cacheKey Key
	value    interface{}
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
		item.Value = cacheElement{key, value} // Обновили значение и положили элемент кэша с этим значением
		c.queue.MoveToFront(item)             // Переместили в начало очереди
		return true
	}

	if c.queue.Len() >= c.capacity { // Елемента нет в словаре и длинна очереди больше вместимости кэша
		last := c.queue.Back()                              // Получаем последний элемент очереди
		delete(c.items, last.Value.(cacheElement).cacheKey) // удаляем последний элемент из мапы
		c.queue.Remove(last)                                // Удаляем последний элемент из очереди
	}

	item := c.queue.PushFront(cacheElement{key, value}) // Перемещаем в начало очереди
	c.items[key] = item                                 // Сохраняем в мапу
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok { // Элемент присутсвует в словаре
		c.queue.MoveToFront(c.items[key]) // Перемещаем элемент в начало очереди

		return item.Value.(cacheElement).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
