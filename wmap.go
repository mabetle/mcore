package mcore

type KeyValue map[string]interface{}

func NewKeyValue() KeyValue {
	return make(map[string]interface{})
}

func (c KeyValue) GetKeys() []string {
	return GetMapKeys(c)
}

func (c KeyValue) IsHasKey(key string) bool {
	return IsMapHasKey(c, key)
}

func GetMapKeys(m map[string]interface{}) (r []string) {
	for k, _ := range m {
		r = append(r, k)
	}
	return
}

func IsMapHasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
