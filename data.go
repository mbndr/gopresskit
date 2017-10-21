// here are structs stored that are only used for unmarshalling
package presskit

// game represents the data of a game
type game struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	History     []struct {
		Header string `xml:"header"`
		Text   string `xml:"text"`
	} `xml:"history>paragraph"`
	Videos           []video  `xml:"videos>video"`
	ReleaseDates     []string `xml:"release-dates>release-date"`
	Website          link     `xml:"website"`
	PressRequestCopy bool     `xml:"press-can-request-copy"`
	Platforms        []link   `xml:"platforms>platform"`
	Prices           []string `xml:"prices>price"`
	Features         []string `xml:"features>feature"`
	//Social  []link   `xml:"social>link"`
	Quotes      []quote      `xml:"quotes>quote"`
	Awards      []award      `xml:"awards>award"`
	Additionals []additional `xml:"additionals>additional"`
	Credits     []credit     `xml:"credits>credit"`
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
	Address     []string     `xml:"address>line"`
	Phone       string       `xml:"phone"`
	Social      []link       `xml:"social>link"`
	Videos      []video      `xml:"videos>video"`
	Quotes      []quote      `xml:"quotes>quote"`
	Awards      []award      `xml:"awards>award"`
	Additionals []additional `xml:"additionals>additional"`
	Credits     []credit     `xml:"credits>credit"`
	Contacts    []struct {
		Name string `xml:"name"`
		Link link   `xml:"link"`
	} `xml:"contacts>contact"`
}

// link representation
type link struct {
	Label string `xml:"label"`
	Link  string `xml:"link"`
}

// quote representation
type quote struct {
	Text string `xml:"text"`
	From string `xml:"from"`
	Link link   `xml:"link"`
}

// award representation
type award struct {
	Description string `xml:"description"`
	Info        string `xml:"info"`
}

// additional representation
type additional struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        link   `xml:"link"`
}

// credit representation
type credit struct {
	Person string `xml:"person"`
	Role   string `xml:"role"`
	Link   string `xml:"link"`
}
