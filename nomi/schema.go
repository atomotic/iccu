package nomi

import (
	"encoding/json"
	"strings"
)

type MainResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q      string `json:"q"`
			Indent string `json:"indent"`
			Fl     string `json:"fl"`
			Start  string `json:"start"`
			QOp    string `json:"q.op"`
			Rows   string `json:"rows"`
			Wt     string `json:"wt"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response Response `json:"response"`
}

type Response struct {
	NumFound      int   `json:"numFound"`
	Start         int   `json:"start"`
	NumFoundExact bool  `json:"numFoundExact"`
	Docs          []Doc `json:"docs"`
}

type Doc struct {
	ID      string          `json:"id"`
	Unimarc json.RawMessage `json:"unimarc"`

	// cached indexed fields to avoid repeated unmarshaling and iteration
	fieldIndex map[string]any
}

// parse unmarshals the UNIMARC data once and caches the result in an indexed map.
// Subsequent calls return the cached data.
func (d *Doc) parse() map[string]any {
	if d.fieldIndex != nil {
		return d.fieldIndex
	}

	d.fieldIndex = make(map[string]any)

	var unimarc struct {
		Fields []json.RawMessage `json:"fields"`
	}
	if err := json.Unmarshal(d.Unimarc, &unimarc); err != nil {
		return d.fieldIndex
	}

	for _, field := range unimarc.Fields {
		var fieldMap map[string]any
		if err := json.Unmarshal(field, &fieldMap); err != nil {
			continue
		}
		// Index each field by its number
		for key, val := range fieldMap {
			d.fieldIndex[key] = val
		}
	}

	return d.fieldIndex
}

// Get retrieves a value from a UNIMARC field by field number.
// For control fields (like "003"), it returns the string value directly.
// For data fields (like "200"), it returns the field structure (map[string]any).
func (d *Doc) Get(fieldNum string) any {
	fields := d.parse()
	return fields[fieldNum]
}

func (d *Doc) Bid() string {
	val := d.Get("003")
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

func (d *Doc) Name() string {
	val := d.Get("200")
	field200, ok := val.(map[string]any)
	if !ok {
		return ""
	}

	subfields, ok := field200["subfields"].([]any)
	if !ok {
		return ""
	}

	var name []string
	for _, sub := range subfields {
		subMap, ok := sub.(map[string]any)
		if !ok {
			continue
		}
		if a, ok := subMap["a"].(string); ok && a != "" {
			name = append(name, a)
		}
		if b, ok := subMap["b"].(string); ok && b != "" {
			name = append(name, b)
		}
	}
	return strings.Join(name, "")
}

// Date extracts birth/death dates from UNIMARC field 200, subfield $f.
// Returns the date string (e.g., "1265-1321", "1875-", "n. 1950") or empty string if not found.
func (d *Doc) Date() string {
	val := d.Get("300")
	field300, ok := val.(map[string]any)
	if !ok {
		return ""
	}

	subfields, ok := field300["subfields"].([]any)
	if !ok {
		return ""
	}

	for _, sub := range subfields {
		subMap, ok := sub.(map[string]any)
		if !ok {
			continue
		}
		if a, ok := subMap["a"].(string); ok && a != "" {
			return a
		}
	}
	return ""
}
