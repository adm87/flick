package data

import (
	"github.com/adm87/flick/pkg/resources"
	"github.com/adm87/flick/pkg/structures/slotmap"
)

type DataStore struct {
	data *slotmap.SlotMap[[]byte]
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: slotmap.New[[]byte](256),
	}
}

func (is *DataStore) Get(handle resources.ResourceHandle) ([]byte, error) {
	k := slotmap.Unpack(handle)

	data, err := is.data.Get(k)
	if err != nil {
		return nil, resources.GetError(err)
	}

	return data, nil
}

func (is *DataStore) Set(handle resources.ResourceHandle, data []byte) error {
	k := slotmap.Unpack(handle)

	if err := is.data.Set(k, data); err != nil {
		return resources.GetError(err)
	}

	return nil
}

func (is *DataStore) Put(data []byte) resources.ResourceHandle {
	k := is.data.Insert(data)
	return k.Pack()
}

func (is *DataStore) Delete(handle resources.ResourceHandle) error {
	k := slotmap.Unpack(handle)

	if _, err := is.data.Delete(k); err != nil {
		return resources.GetError(err)
	}

	return nil
}
