package nomi

import "strings"

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
	ID      string `json:"id"`
	Unimarc struct {
		Leader string `json:"leader"`
		Fields []struct {
			Num100 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"100,omitempty"`
			Num152 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"152,omitempty"`
			Num200 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					B string `json:"b,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"200,omitempty"`
			Num801 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					B string `json:"b,omitempty"`
					C string `json:"c,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"801,omitempty"`
			Num001 string `json:"001,omitempty"`
			Num005 string `json:"005,omitempty"`
			Num003 string `json:"003,omitempty"`
		} `json:"fields"`
	} `json:"unimarc"`
}

func (d *Doc) Bid() string {
	for _, field := range d.Unimarc.Fields {
		if field.Num003 != "" {
			return field.Num003
		}
	}
	return ""
}

func (d *Doc) Name() string {
	var name []string

	for _, field := range d.Unimarc.Fields {
		for _, sub := range field.Num200.Subfields {
			if sub.A != "" {
				name = append(name, sub.A)
			}
			if sub.B != "" {
				name = append(name, sub.B)
			}

		}
	}
	return strings.Join(name, "")
}
