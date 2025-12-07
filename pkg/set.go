package pkg

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Set[T comparable, V any] struct {
	items map[T]V
}

type NumericalSet[T comparable, V Number] struct {
	*Set[T, V]
}

func CreateSet[T comparable, V any](init map[T]V) *Set[T, V] {
	if init == nil {
		init = make(map[T]V)
	}

	return &Set[T, V]{
		items: init,
	}
}

func CreateNumericalSet[T comparable, V Number](init map[T]V) *NumericalSet[T, V] {
	return &NumericalSet[T, V]{
		Set: CreateSet(init),
	}
}

func (s *Set[T, V]) Add(key T, value V) {
	s.items[key] = value
}

func (s *Set[T, V]) Delete(key T) {
	delete(s.items, key)
}

func (s *Set[T, V]) GetList() []T {
	list := make([]T, 0, len(s.items))

	for v := range s.items {
		list = append(list, v)
	}

	return list
}

func (s *Set[T, V]) Exists(key T) bool {
	_, ok := s.items[key]
	return ok
}

func (s *Set[T, V]) GetMap() map[T]V {
	return s.items
}

func (s *Set[T, V]) GetSize() int {
	return len(s.items)
}

func (s *NumericalSet[T, V]) AddOrSet(key T, value V) V {
	if v, ok := s.items[key]; ok {
		s.items[key] = v + value
	} else {
		s.items[key] = value
	}

	return s.items[key]
}
