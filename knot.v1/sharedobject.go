package knot

import "sync"

type SharedObject struct {
	data map[string]interface{}
}

var instance *SharedObject
var once sync.Once

func GetSharedObject() *SharedObject {
	once.Do(func() {
		instance = &SharedObject{}
		instance.data = map[string]interface{}{}
	})

	return instance
}

func (s *SharedObject) Get(key string, defaultValue interface{}) interface{} {
	if data, isExist := s.data[key]; isExist {
		return data
	}

	return defaultValue
}

func (s *SharedObject) Set(key string, value interface{}) {
	s.data[key] = value
}

func (s *SharedObject) All() map[string]interface{} {
	return s.data
}
