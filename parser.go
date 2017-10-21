package presskit

import (
	"encoding/xml"
)

// Parser parses all data files
type Parser interface {
	Extension() string
	Company([]byte) (*company, error)
	Game([]byte) (*game, error)
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

// Game parses data of a game
func (p XMLParser) Game(data []byte) (*game, error) {
	var g game
	err := xml.Unmarshal(data, &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
