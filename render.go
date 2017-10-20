package presskit

import (
	"bytes"
	"html/template"
)

// companyData is given to a company template
type companyData struct {
	Company company
	Media   media
}

// gameData is given to a game template
type gameData struct {
	companyData
	Game game
}

// renderCompany renders a company html
// TODO: base and company template (remove company-game redundancy)
func renderCompany(c company, m media) ([]byte, error) {
	var buf bytes.Buffer

	data := companyData{
		Company: c,
		Media:   m,
	}

	// read template
	tplData, err := Asset("templates/base.html")
	if err != nil {
		return nil, err
	}

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

// renderGame renders a game html
func renderGame(c company, g game) ([]byte, error) {
	return []byte(""), nil
}
