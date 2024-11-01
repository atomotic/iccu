package client

type FullResponse struct {
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

type Doc map[string]interface{}
