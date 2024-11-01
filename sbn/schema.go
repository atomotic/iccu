package sbn

import (
	"fmt"
	"net/url"
	"strings"
)

func Host(sourceurl string) (string, error) {
	parsedURL, err := url.Parse(sourceurl)
	if err != nil {
		return "", err
	}

	host := parsedURL.Host
	if host == "" {
		return "", fmt.Errorf("no host")
	}
	protocol := parsedURL.Scheme

	return fmt.Sprintf("%s://%s", protocol, host), nil
}

type Doc struct {
	Tiporec       string   `json:"tiporec"`
	Pres          []string `json:"pres"`
	Autore        string   `json:"autore"`
	ID            string   `json:"id"`
	IDMOL         string   `json:"id_mol"`
	DigPreview    []string `json:"dig_preview"`
	DigPreviewURL []string `json:"dig_previewUrl"`
	DigManifest   []string `json:"dig_manifest"`
	Level         string   `json:"level"`
	FormaCont     []string `json:"forma_cont"`
	Lingua        []string `json:"lingua"`
	DigCover      []string `json:"dig_cover"`
	Baseprov      string   `json:"baseprov"`
	TipoTit       string   `json:"tipo_tit"`
	DataAgg       string   `json:"data_agg"`
	DataDa        int      `json:"datada"`
	DataA         int      `json:"dataa"`
	Unimarc       struct {
		Leader string `json:"leader"`
		Fields []struct {
			Num100 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"100,omitempty"`
			Num101 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"101,omitempty"`
			Num102 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"102,omitempty"`
			Num125 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"125,omitempty"`
			Num181 struct {
				Subfields []struct {
					Num6 string `json:"6,omitempty"`
					A    string `json:"a,omitempty"`
					B    string `json:"b,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"181,omitempty"`
			Num182 struct {
				Subfields []struct {
					Num6 string `json:"6,omitempty"`
					A    string `json:"a,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"182,omitempty"`
			Num183 struct {
				Subfields []struct {
					Num2 string `json:"2,omitempty"`
					Num6 string `json:"6,omitempty"`
					A    string `json:"a,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"183,omitempty"`
			Num200 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"200,omitempty"`
			Num208 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"208,omitempty"`
			Num210 struct {
				Subfields []struct {
					D string `json:"d"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"210,omitempty"`
			Num215 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					D string `json:"d,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"215,omitempty"`
			Num300 struct {
				Subfields []struct {
					A string `json:"a"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"300,omitempty"`
			Num500 struct {
				Subfields []struct {
					Num3 string `json:"3,omitempty"`
					Num9 string `json:"9,omitempty"`
					A    string `json:"a,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"500,omitempty"`
			Num510 struct {
				Subfields []struct {
					Num9 string `json:"9,omitempty"`
					A    string `json:"a,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"510,omitempty"`
			Num517 struct {
				Subfields []struct {
					Num9 string `json:"9,omitempty"`
					A    string `json:"a,omitempty"`
					E    string `json:"e,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"517,omitempty"`
			Num700 struct {
				Subfields []struct {
					Num3 string `json:"3,omitempty"`
					Num4 string `json:"4,omitempty"`
					A    string `json:"a,omitempty"`
					B    string `json:"b,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"700,omitempty"`
			Num702 struct {
				Subfields []struct {
					Num3 string `json:"3,omitempty"`
					Num4 string `json:"4,omitempty"`
					A    string `json:"a,omitempty"`
					B    string `json:"b,omitempty"`
					F    string `json:"f,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"702,omitempty"`
			Num790 struct {
				Subfields []struct {
					Num3 string `json:"3,omitempty"`
					A    string `json:"a,omitempty"`
					B    string `json:"b,omitempty"`
					Z    string `json:"z,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"790,omitempty"`
			Num801 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					B string `json:"b,omitempty"`
					C string `json:"c,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"801,omitempty"`
			Num899 struct {
				Subfields []struct {
					Num1 string `json:"1,omitempty"`
					Num2 string `json:"2,omitempty"`
					Num3 string `json:"3,omitempty"`
					F    string `json:"f,omitempty"`
					C    string `json:"c,omitempty"`
					E    string `json:"e,omitempty"`
					T    string `json:"t,omitempty"`
					U    string `json:"u,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"899,omitempty"`
			Num856 struct {
				Subfields []struct {
					Num1 string `json:"1,omitempty"`
					Num2 string `json:"2,omitempty"`
					Num3 string `json:"3,omitempty"`
					F    string `json:"f,omitempty"`
					C    string `json:"c,omitempty"`
					E    string `json:"e,omitempty"`
					T    string `json:"t,omitempty"`
					U    string `json:"u,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"856,omitempty"`
			Num923 struct {
				Subfields []struct {
					B string `json:"b,omitempty"`
					C string `json:"c,omitempty"`
					D string `json:"d,omitempty"`
					E string `json:"e,omitempty"`
					G string `json:"g,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"923,omitempty"`
			Num926 struct {
				Subfields []struct {
					C string `json:"c,omitempty"`
					F string `json:"f,omitempty"`
					G string `json:"g,omitempty"`
					H string `json:"h,omitempty"`
					I string `json:"i,omitempty"`
					L string `json:"l,omitempty"`
					M string `json:"m,omitempty"`
					O string `json:"o,omitempty"`
					P string `json:"p,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"926,omitempty"`
			Num927 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					B string `json:"b,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"927,omitempty"`
			Num928 struct {
				Subfields []struct {
					A string `json:"a,omitempty"`
					B string `json:"b,omitempty"`
					C string `json:"c,omitempty"`
					Z string `json:"z,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"928,omitempty"`
			Num929 struct {
				Subfields []struct {
					D string `json:"d,omitempty"`
					G string `json:"g,omitempty"`
					Z string `json:"z,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"929,omitempty"`
			Num999 struct {
				Subfields []struct {
					Num1 string `json:"1,omitempty"`
					Num2 string `json:"2,omitempty"`
					Num3 string `json:"3,omitempty"`
					Num9 string `json:"9,omitempty"`
				} `json:"subfields"`
				Ind1 string `json:"ind1"`
				Ind2 string `json:"ind2"`
			} `json:"999,omitempty"`
			Num001 string `json:"001,omitempty"`
			Num003 string `json:"003,omitempty"`
			Num005 string `json:"005,omitempty"`
		} `json:"fields"`
	} `json:"unimarc"`
	Isbd     string   `json:"isbd"`
	PreTitle string   `json:"pre_titolo"`
	PaeseMus []string `json:"paese_mus"`
	Tipomat  []string `json:"tipomat"`
	Publish  string   `json:"publish"`
	// Timestamp time.Time `json:"timestamp"`
}

func (d *Doc) Bid() string {
	for _, field := range d.Unimarc.Fields {
		if field.Num001 != "" {
			return field.Num001
		}
	}
	return ""
}

func (d *Doc) BidLink() string {
	for _, field := range d.Unimarc.Fields {
		if field.Num003 != "" {
			return field.Num003
		}
	}
	return ""
}

func (d *Doc) IDMol() string {
	if d.IDMOL != "" {
		return d.IDMOL
	}
	return ""
}

func (d *Doc) Title() string {
	return fmt.Sprintf("%s%s", d.PreTitle, d.Isbd)
}

func (d *Doc) Type() string {
	if d.Tiporec != "" {
		return DocumentTypes[d.Tiporec]
	}
	return ""
}

func (d *Doc) Material() (material []string) {
	for _, mat := range d.Tipomat {
		if mat != "" {
			material = append(material, MaterialTypes[mat])
		}
	}
	return material
}

func (d *Doc) StartDate() int {
	return d.DataDa
}

func (d *Doc) EndDate() int {
	return d.DataA
}

func (d *Doc) Manifests() (manifests []string) {
	for i, manifest := range d.DigManifest {
		if manifest != "" {
			if len(d.DigPreviewURL) == 0 {
				manifests = append(manifests, fmt.Sprintf("%s%s", DefaultIIIFHost, manifest))
			} else if len(d.DigPreviewURL) > i {
				host, err := Host(strings.Trim(d.DigPreviewURL[i], "\""))
				if err != nil {
					continue
				}
				manifests = append(manifests, fmt.Sprintf("%s/%s", host, manifest))
			}
		}
	}
	return manifests
}

func (d *Doc) Thumbnails() (thumbnails []string) {
	if len(d.DigPreview) > 0 {
		for _, thumb := range d.DigPreview {
			_, err := Host(thumb)
			if err != nil {
				thumbnails = append(thumbnails, fmt.Sprintf("%s%s", DefaultIIIFHost, thumb))
			} else {
				thumbnails = append(thumbnails, thumb)
			}
		}
	} else if len(d.DigPreviewURL) > 0 {
		thumbnails = append(thumbnails, d.DigPreviewURL...)
	}
	return thumbnails
}

func (d *Doc) ExternalLink() (links []string) {
	for _, field := range d.Unimarc.Fields {
		if field.Num899.Subfields != nil {
			for _, subfield := range field.Num899.Subfields {
				if subfield.U != "" {
					parts := strings.Split(subfield.U, " | ")
					links = append(links, parts...)
				}
			}
		}
		if field.Num856.Subfields != nil {
			for _, subfield := range field.Num856.Subfields {
				if subfield.U != "" {
					parts := strings.Split(subfield.U, " | ")
					links = append(links, parts...)
				}
			}
		}

	}
	return links
}
