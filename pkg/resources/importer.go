package resources

type ResourceImporter interface {
	Import(raw []byte) (ResourceHandle, error)
	Delete(handle ResourceHandle) error
}
