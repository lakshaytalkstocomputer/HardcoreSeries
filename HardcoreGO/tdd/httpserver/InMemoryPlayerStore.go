package httpserver


type InMemoryPlayerStore struct {
	Store map[string] int
}

func NewInMemoryStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{ map[string] int {}}
}

func (i *InMemoryPlayerStore) RecordWin(name string){
	i.Store[name]++
}

func(i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.Store[name]
}
