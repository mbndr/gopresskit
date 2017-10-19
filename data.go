package presskit

// here are structs stored that are only used for unmarshalling

// link represents a link with a label
type link struct {
	Label string `xml:"label"`
	Link  string `xml:"link"`
}

// company represents the data of a company
type company struct {
	Title        string `xml:"title"`
	Website      link   `xml:"website"`
	FoundingDate string `xml:"founding-date"`
	BasedIn      string `xml:"based-in"`
	PressContact string `xml:"press-contact"`
	Description  string `xml:"description"`
	History      []struct {
		Header string `xml:"header"`
		Text   string `xml:"text"`
	} `xml:"history>paragraph"`
	Address []string `xml:"address>line"`
	Phone   string   `xml:"phone"`
	Social  []link   `xml:"social>link"`
	// videos are loaded with a media object
	Videos  []video  `xml:"videos>video"`
	Quotes  []struct {
		Text string `xml:"text"`
		From string `xml:"from"`
		Link link   `xml:"link"`
	} `xml:"quotes>quote"`
	Awards []struct {
		Description string `xml:"description"`
		Info        string `xml:"info"`
	} `xml:"awards>award"`
	Additionals []struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        link   `xml:"link"`
	} `xml:"additionals>additional"`
	Credits []struct {
		Person string `xml:"person"`
		Role   string `xml:"role"`
		Link   string `xml:"link"`
	} `xml:"credits>credit"`
	Contacts []struct {
		Name string `xml:"name"`
		Link   link `xml:"link"`
	} `xml:"contacts>contact"`
}

type game struct {
}
