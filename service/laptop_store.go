package service

import (
	"errors"
	"fmt"
	"pcbook/pb"
	"sync"

	"github.com/jinzhu/copier"
)

//ErrAlreadyExists is returned when the laptop ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

//LaptopStore is an interface to abstract laptop storage
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

//InMemoryLaptopStore in memory storage implementation
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

//NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

//Save writes the laptop to the Store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	//we need to deep copy to prevent the outside from changing our store
	laptop2 := &pb.Laptop{}
	copier.Copy(laptop2, laptop)

	store.data[laptop2.Id] = laptop2
	return nil
}

//Find returns a stored laptop matching the id
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	laptop2 := &pb.Laptop{}
	err := copier.Copy(laptop2, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %v", err)
	}

	return laptop2, nil
}
