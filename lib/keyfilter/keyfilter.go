package keyfilter

// KeyFilter holds facter keys to filter by
type KeyFilter struct {
	keys []string
}

// NewFilter returns new KeyFilter
func NewFilter() *KeyFilter {
	return &KeyFilter{}
}

// Add a facter key
func (ff *KeyFilter) AddOne(k string) {
	ff.keys = append(ff.keys, k)
}

// Add a slice of keys
func (kf *KeyFilter) AddMany(k []string) {
	kf.keys = append(kf.keys, k...)
}

// Get all filter keys
func (kf *KeyFilter) Get() map[string]bool {
	mykeys := make(map[string]bool)
	for _, v := range kf.keys {
		mykeys[v] = true
	}
	return mykeys
}
