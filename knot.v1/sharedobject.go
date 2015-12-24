package knot

//--- please remove this, do manual checking instead to reduce memory cost allocation
import "sync"

//--- please use M for data. M already has all method requried
//--- change variable name to sharedObject to make it not exposable to outside
type SharedObject struct {
	data map[string]interface{}
}

var instance *SharedObject
var once sync.Once

//--- change to SharedObject
func GetSharedObject() *SharedObject {
	//--- use check if nil
	once.Do(func() {
		instance = &SharedObject{}
		instance.data = map[string]interface{}{}
	})

	return instance
}

//-- will be handled by M
func (s *SharedObject) Get(key string) interface{} {
	if data, isExist := s.data[key]; isExist {
		return data
	}

	return nil
}

//-- will be handled by M
func (s *SharedObject) GetWithDefaultValue(key string, defaultValue interface{}) interface{} {
	if data, isExist := s.data[key]; isExist {
		return data
	}

	return defaultValue
}

//-- will be handled by M
func (s *SharedObject) Set(key string, value interface{}) {
	s.data[key] = value
}

//-- not neccesarily needed, can access knot.SharedObject() directly
func (s *SharedObject) All() map[string]interface{} {
	return s.data
}
