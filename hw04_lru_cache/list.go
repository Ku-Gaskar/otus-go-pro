package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.head == nil { // Установка первого элемента списка
		l.head = item
		l.tail = item
	} else {
		item.Next = l.head
		l.head.Prev = item
		l.head = item
	}
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.tail == nil { // Установка первого элемента списка
		l.head = item
		l.tail = item
	} else {
		item.Prev = l.tail
		l.tail.Next = item
		l.tail = item
	}
	l.len++
	return item
}

func (l *list) Remove(item *ListItem) {
	switch {
	case item == l.head && l.head.Next != nil:
		l.head.Next.Prev = nil
		l.head = l.head.Next
	case item == l.head && l.head.Next == nil:
		l.tail = nil
		l.head = nil
	case item == l.tail:
		l.tail.Prev.Next = nil
		l.tail = l.tail.Prev
	default:
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
	}
	l.len--
}

func (l *list) MoveToFront(item *ListItem) {
	if l.head == item {
		return
	}
	l.Remove(item)

	item.Next = l.head
	item.Prev = nil
	l.head.Prev = item
	l.head = item
	l.len++
}

func NewList() List {
	return new(list)
}
