package safetype

import "sync"

// 类型安装的数组泛型类型
type safeArr[T any] struct {
	sync.Mutex
	data []T
}

func (self *safeArr[T]) Add(item T) {
	self.Lock()
	self.data = append(self.data, item)
	self.Unlock()
}

func (self *safeArr[T]) Remove(index int) {
	self.Lock()
	if index >= 0 && index < len(self.data) {
		self.data = append(self.data[:index], self.data[index+1:]...)
	}
	self.Unlock()
}

func (self *safeArr[T]) Data() []T {
	self.Lock()
	self.Unlock()
	return self.data
}
