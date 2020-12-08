package lepton

// OnPrem provider for ops
type OnPrem struct{}

// Initialize on prem provider
func (p *OnPrem) Initialize(config *ProviderConfig) error {
	return nil
}

// GetStorage returns storage interface for cloud provider
func (p *OnPrem) GetStorage() Storage {
	return nil
}
