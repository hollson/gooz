package util

import (
	"sort"
	"sync"
)

type Set struct {
	sync.RWMutex
	M map[string]bool
}

// 新建集合对象
func NewSet(items ...string) *Set {
	s := &Set{
		M: make(map[string]bool, len(items)),
	}
	s.Add(items...)
	return s
}

// 添加元素
func (s *Set) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.M[v] = true
	}
}

// 删除元素
func (s *Set) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.M, v)
	}
}

// 判断元素是否存在
func (s *Set) Has(items ...string) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.M[v]; !ok {
			return false
		}
	}
	return true
}

// 元素个数
func (s *Set) Count() int {
	return len(s.M)
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.M = map[string]bool{}
}

// 空集合判断
func (s *Set) Empty() bool {
	return len(s.M) == 0
}

// 无序列表
func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.M))
	for item := range s.M {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *Set) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.M))
	for item := range s.M {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}

// 并集
func (s *Set) Union(sets ...*Set) *Set {
	r := NewSet(s.List()...)
	for _, set := range sets {
		for e := range set.M {
			r.M[e] = true
		}
	}
	return r
}

// 差集
func (s *Set) Minus(sets ...*Set) *Set {
	r := NewSet(s.List()...)
	for _, set := range sets {
		for e := range set.M {
			if _, ok := s.M[e]; ok {
				delete(r.M, e)
			}
		}
	}
	return r
}

// 交集
func (s *Set) stringersect(sets ...*Set) *Set {
	r := NewSet(s.List()...)
	for _, set := range sets {
		for e := range s.M {
			if _, ok := set.M[e]; !ok {
				delete(r.M, e)
			}
		}
	}
	return r
}

// 补集
func (s *Set) Complement(full *Set) *Set {
	r := NewSet()
	for e := range full.M {
		if _, ok := s.M[e]; !ok {
			r.Add(e)
		}
	}
	return r
}
