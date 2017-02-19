package facter

import (
	"github.com/dliappis/go-facter/lib/formatter"
	"github.com/dliappis/go-facter/lib/keyfilter"
)

// Facter struct holds Facter-related attributes
type Facter struct {
	facts     map[string]interface{}
	formatter Formatter
	keyfilter KeyFilter
}

// Config struct serves to pass Facter configuration
type Config struct {
	Formatter Formatter
	KeyFilter KeyFilter
}

// Formatter interface
type Formatter interface {
	Print(map[string]interface{}, [][]string, int) error
}

// KeyFilter interface
type KeyFilter interface {
	// Get() map[string]bool
	Get() [][]string
	AddOne(k string)
	AddMany(k []string)
}

// New returns new instance of Facter
func New(userConf *Config) *Facter {
	var conf *Config
	if userConf != nil {
		conf = userConf
	} else {
		conf = &Config{
			Formatter: formatter.NewFormatter(),
			KeyFilter: keyfilter.NewFilter(),
		}
	}
	f := &Facter{
		facts:     make(map[string]interface{}),
		formatter: conf.Formatter,
		keyfilter: conf.KeyFilter,
	}
	return f
}

// Add adds a fact
func (f *Facter) Add(k string, v interface{}) {
	f.facts[k] = v
}

// Delete deletes given fact
func (f *Facter) Delete(k string) {
	delete(f.facts, k)
}

// Get returns value of given fact, if it exists
func (f *Facter) Get(k string) (interface{}, bool) {
	value, ok := f.facts[k]
	return value, ok
}

// Print prints-out facts by calling formatter
func (f *Facter) Print() {
	f.formatter.Print(f.facts, f.keyfilter.Get(), 0)
}
