package sbn

const DefaultIIIFHost = "https://jmms.iccu.sbn.it"

var DocumentTypes = map[string]string{
	"a": "Testo",
	"b": "Testo manoscritto",
	"c": "Musica a stampa",
	"d": "Musica manoscritta",
	"e": "Risorsa cartografica a stampa",
	"f": "Risorsa cartografica manoscritta",
	"g": "Risorsa da proiettare o video",
	"i": "Registrazione sonora non musicale",
	"j": "Registrazione sonora musicale",
	"k": "Risorsa grafica",
	"l": "Risorsa elettronica",
	"r": "Oggetto tridimensionale",
	"m": "Risorsa multimediale",
}

var MaterialTypes = map[string]string{
	"v": "Audiovisivi",
	"c": "Cartografia",
	"g": "Grafica",
	"V": "Audiovisivi",
	"C": "Cartografia",
	"G": "Grafica",
	"A": "Libro antico",
	"N": "Libro moderno",
	"M": "Musica",
}
