package formatter

import (
	"encoding/json"
	"fmt"
	"os"
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
	if len(keyfilters) > 0 {
		PrintFilteredFacts(keyfilters, facts)
	} else {
		b, err := json.MarshalIndent(facts, "", "  ")
		if err == nil {
			os.Stdout.Write(b)
			fmt.Println("")
		}
	}

	return nil
}

// loop through all our filters and print matching sections of facts
func PrintFilteredFacts(keyfilters [][]string, facts map[string]interface{}) {
	for _, kfv := range keyfilters {
		res := FilterFacts(kfv, facts)
		switch restype := res.(type) {
		case string:
			fmt.Println(restype)
		case map[string]interface{}:
			bjson, err := json.MarshalIndent(restype, "", "  ")
			if err == nil {
				os.Stdout.Write(bjson)
				fmt.Println("")
			}
		}
	}
}

/* try to match a filter string (e.g. os.release) in our nested map facts
   and return leaf fact element
*/
func FilterFacts(keyfilter []string, facts map[string]interface{}) interface{} {
	myfacts := facts

	for level, filterkey := range keyfilter {
		for factkey, factvalue := range myfacts {
			if filterkey == factkey && level == len(keyfilter)-1 {
				// we've matched the whole filter string. Return the rest of the fact structure as it.
				switch factvalue.(type) {
				case map[string]interface{}:
					return factvalue.(map[string]interface{})
				case string:
					return factvalue.(string)
				}
			} else if filterkey == factkey && level < len(keyfilter)-1 {
				// matched one part of our filter, go deeper in our facts
				myfacts = myfacts[factkey].(map[string]interface{})
				break
			}
		}
	}
	return nil // nothing matched ...
}
