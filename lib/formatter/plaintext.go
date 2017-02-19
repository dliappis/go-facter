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
func (pf PlainTextFormatter) Print(facts map[string]interface{}, keyfilters [][]string, level int) error {
	ThePrinter(facts, keyfilters, level)
	// var keys []string

	// for k := range facts {
	// 	if len(keyfilters) == 0 || keyfilters[k] {
	// 		keys = append(keys, k)
	// 	}
	// }
	// sort.Strings(keys)
	// for _, k := range keys {
	// 	fmt.Printf("%v => %v\n", k, facts[k])
	// }
	return nil
}

func ThePrinter(facts map[string]interface{}, keyfilters [][]string, level int) {
	var keys []string
	for k := range facts {
		// User didn't provide filters. Present all facts.
		if len(keyfilters) == 0 {
			keys = append(keys, k)
			break
		}

		for _, keyslice := range keyfilters {
			//fmt.Printf("We are at level %d and keyslice is %+v\n", level, keyslice)
			if level < len(keyslice) {
				keyflt := keyslice[level]
				if len(keyslice) == 0 || keyflt == k {
					//fmt.Printf("Good news! Just added %+v\n", k)
					keys = append(keys, k)
					break
				}
			} else if level >= len(keyslice) {
				// No more filters for facts of this depth, append all
				keys = append(keys, k)
				break
			}
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		// if _, ok := facts[k].(map[string]interface{}); ok {
		// 	fmt.Printf("Hm there's a map here!")
		// 	ThePrinter(facts[k].(map[string]interface{}), keyfilters)
		// } else {
		// 	fmt.Printf("%v => %v\n", k, facts[k])
		// }
		switch facts[k].(type) {
		case map[string]interface{}:
			fmt.Printf("%v => ", k)
			ThePrinter(facts[k].(map[string]interface{}), keyfilters, level+1)
		default:
			fmt.Printf("%v => %v\n", k, facts[k])
		}
	}
}
