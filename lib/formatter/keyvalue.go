package formatter

import (
	"fmt"
	"sort"
)

// KeyValueFormatter prints-out facts in k:v format
type KeyValueFormatter struct {
}

// NewKeyValueFormatter returns new key-value formatter
func NewKeyValueFormatter() *KeyValueFormatter {
	return &KeyValueFormatter{}
}

// Print prints-out facts in k:v format
func (kvf KeyValueFormatter) Print(facts map[string]interface{}, keyfilters map[string]bool) error {
	fmt.Println("Len is", len(keyfilters))
	var keys []string
	for k := range facts {
		if len(keyfilters) == 0 || keyfilters[k] {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, facts[k])
	}
	return nil
}
