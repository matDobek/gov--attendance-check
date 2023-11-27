package cache

type Cache interface {
	Get(id string) (string, error)
	Put(id string, content string) error
	Clear() error
}
