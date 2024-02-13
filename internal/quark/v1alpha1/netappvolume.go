package v1alpha1

// SetNoSnapDir sets the NoSnapDir option for dmap
func (n *NetAppVolume) SetNoSnapDir(val bool) {
	n.Spec.NoSnapDir = val
}
// GetNoSnapDir sets the NoSnapDir option for dmap
func (n *NetAppVolume) GetNoSnapDir() bool{
	return n.Spec.NoSnapDir
}