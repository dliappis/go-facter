package keyfilter

import (
	"strings"
)

// KeyFilter holds slices of facter keys to filter by
type KeyFilter struct {
	keys [][]string
}

// NewFilter returns new KeyFilter
func NewFilter() *KeyFilter {
	return &KeyFilter{}
}

// Add a facter key
func (ff *KeyFilter) AddOne(k string) {
	ff.keys = append(ff.keys, []string{k})
}

/*
Add all cliargs as dotseparated slices to the slice `keys`
e.g. the cli args "os.kernelversion facterversion" will create
[["os","kernerlversion"], ["facterversion"]]
*/
func (kf *KeyFilter) AddMany(k []string) {
	for _, cliarg := range k {
		cliargslice := strings.Split(cliarg, ".")
		kf.keys = append(kf.keys, cliargslice)
	}
	// left over when we treated each cliarg as a string regardless of dots
	//kf.keys = append(kf.keys, k...)
}

// Get all filter keys
func (kf *KeyFilter) Get() [][]string {
	// mykeys := make(map[string]bool)
	// for _, v := range kf.keys {
	// 	mykeys[v] = true
	// }
	// return mykeys
	return kf.keys
}
