package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestList_2(t *testing.T) {
	l := NewList()

	// Test Len()
	if l.Len() != 0 {
		t.Errorf("Expected Len() to be 0, got %d", l.Len())
	}

	// Test Front() and Back()
	if l.Front() != nil || l.Back() != nil {
		t.Errorf("Expected Front() and Back() to be nil, got %v and %v", l.Front(), l.Back())
	}

	// Test PushFront()
	item1 := l.PushFront("item1")
	if l.Len() != 1 || l.Front() != item1 || l.Back() != item1 {
		t.Errorf("Expected Len() to be 1, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item1, item1, l.Len(), l.Front(), l.Back())
	}

	item2 := l.PushFront("item2")
	if l.Len() != 2 || l.Front() != item2 || l.Back() != item1 {
		t.Errorf("Expected Len() to be 2, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item2, item1, l.Len(), l.Front(), l.Back())
	}

	// Test PushBack()
	item3 := l.PushBack("item3")
	if l.Len() != 3 || l.Front() != item2 || l.Back() != item3 {
		t.Errorf("Expected Len() to be 3, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item2, item3, l.Len(), l.Front(), l.Back())
	}

	item4 := l.PushBack("item4")
	if l.Len() != 4 || l.Front() != item2 || l.Back() != item4 {
		t.Errorf("Expected Len() to be 4, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item2, item4, l.Len(), l.Front(), l.Back())
	}

	// Test Remove()
	l.Remove(item3)
	if l.Len() != 3 || l.Front() != item2 || l.Back() != item4 {
		t.Errorf("Expected Len() to be 3, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item2, item4, l.Len(), l.Front(), l.Back())
	}

	// Test MoveToFront()
	l.MoveToFront(item1)
	if l.Len() != 3 || l.Front() != item1 || l.Back() != item4 {
		t.Errorf("Expected Len() to be 3, Front() to be %v, and Back() to be %v, got %d, %v, and %v",
			item1, item4, l.Len(), l.Front(), l.Back())
	}
}

func TestList_3_Remove_Only_one(t *testing.T) {
	l := NewList()
	l.PushFront("text_1")
	require.Equal(t, 1, l.Len())
	l.Remove(l.Front())
	require.Equal(t, 0, l.Len())
}
