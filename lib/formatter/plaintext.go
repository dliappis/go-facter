package formatter

import (
	"fmt"
	"sort"
)

// PlainTextFormatter prints-out facts in k=>v format
type PlainTextFormatter struct {
}

// NewFormatter returns new plain-text formatter
func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

// Print prints-out facts in k=>v format
func (pf PlainTextFormatter) Print(facts map[string]interface{}, keyfilters map[string]bool) error {
	var keys []string
	for k := range facts {
		if keyfilters[k] {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v => %v\n", k, facts[k])
	}
	fmt.Println(keyfilters)
	return nil
}
