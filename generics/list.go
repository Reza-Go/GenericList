package generics

import "github.com/google/go-cmp/cmp"

type List[T any] struct {
	Items []T
}

func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}

func (list *List[T]) InsertAt(index int, item T) {
	list.Items = append(list.Items, item)
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = item

}

func (list *List[T]) RemoveAt(index int) {
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
}

func (list *List[T]) Remove(item T) {
	index := list.Find(item)
	if index != -1 {
		list.RemoveAt(index)
	}
}

func (list *List[T]) Find(item T) int {
	for i, v := range list.Items {
		if cmp.Equal(v, item) {
			return i
		}
	}
	return -1

}
