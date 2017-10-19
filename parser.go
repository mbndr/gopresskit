package presskit

import (
	"encoding/xml"
)

type Parser interface {
	Extension() string
	Company([]byte) (*company, error)
	//Games() []game
}

// XMLParser parses the data from xml files
type XMLParser struct{}

// Extension returns the file extension the parser uses
func (p XMLParser) Extension() string {
	return "xml"
}

// Company parses the company data
func (p XMLParser) Company(data []byte) (*company, error) {
	var c company
	err := xml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Games parses all game data
/*func (p XMLParser) Games() []Game {
	return nil
}*/
