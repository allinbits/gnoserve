package bmp

type RegistryCallback = func(key string, value interface{}) bool
type RegistryMatcher = func(value interface{}) bool

func HasImageApi(value interface{}) bool {
	_, ok := value.(ImageApi)
	return ok
}

func IterateRegistry(callback RegistryCallback, matcher RegistryMatcher) {
	registry.IterateByOffset(0, registry.Size(), func(key string, value interface{}) bool {
		if matcher(value) {
			return callback(key, value)
		}
		return false
	})
}
