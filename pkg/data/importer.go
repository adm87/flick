package data

import "github.com/adm87/flick/pkg/resources"

var dataTypes = []resources.ResourceType{"json", "yaml", "yml"}

func DataTypes() []resources.ResourceType {
	return dataTypes
}

type DataImporter struct {
	store *DataStore
}

func NewDataImporter(store *DataStore) *DataImporter {
	return &DataImporter{
		store: store,
	}
}

func (di *DataImporter) Import(raw []byte) (resources.ResourceHandle, error) {
	return di.store.Put(raw), nil
}

func (di *DataImporter) Delete(handle resources.ResourceHandle) error {
	return di.store.Delete(handle)
}
