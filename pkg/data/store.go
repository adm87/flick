package data

import (
	"github.com/adm87/flick/pkg/resources"
	"github.com/adm87/flick/pkg/types/structures/slotmap"
)

type DataStore struct {
	data *slotmap.SlotMap[[]byte]
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: slotmap.New[[]byte](256),
	}
}

func (ds *DataStore) Get(handle resources.ResourceHandle) ([]byte, error) {
	k := slotmap.Unpack(handle)

	data, err := ds.data.Get(k)
	if err != nil {
		return nil, resources.GetError(err)
	}

	return data, nil
}

func (ds *DataStore) Set(handle resources.ResourceHandle, data []byte) error {
	k := slotmap.Unpack(handle)

	if err := ds.data.Set(k, data); err != nil {
		return resources.GetError(err)
	}

	return nil
}

func (ds *DataStore) Put(data []byte) resources.ResourceHandle {
	k := ds.data.Insert(data)
	return k.Pack()
}

func (ds *DataStore) Delete(handle resources.ResourceHandle) error {
	k := slotmap.Unpack(handle)

	if _, err := ds.data.Delete(k); err != nil {
		return resources.GetError(err)
	}

	return nil
}
