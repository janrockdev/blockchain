package engine

type Storage interface {
	Put(*Block) error
}

type MemoryStore struct {
}

func NewMemoryStorage() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Put(b *Block) error {
	return nil
}
