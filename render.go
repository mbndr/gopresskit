package presskit

import (
	"bytes"
	"html/template"
)

// companyData is given to a company template
// TODO not happy with outputPath here
type companyData struct {
	Company    company
	Media      media
	outputPath string
}

// gameData is given to a game template
type gameData struct {
	companyData
	Game game
}

// renderCompany renders a company html
func renderCompany(c company, m media, outputPath string) ([]byte, error) {
	var buf bytes.Buffer

	data := companyData{
		outputPath: outputPath,
		Company:    c,
		Media:      m,
	}

	// read template
	tplData, err := Asset("templates/base.html")
	if err != nil {
		return nil, err
	}

	//t, err := template.ParseFiles("./static/templates/base.html")
	t, err := template.New("company").Parse(string(tplData))
	if err != nil {
		return nil, err
	}

	err = t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// renderCompany renders a game html
func renderGame(c company, g game) ([]byte, error) {
	return []byte(""), nil
}
