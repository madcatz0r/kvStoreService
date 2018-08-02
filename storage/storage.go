package storage

type Storage interface {
	Put(string, []byte)
	Get(string) ([]byte, error)
	Delete(string) error
}
