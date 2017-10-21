package presskit

import (
	"bytes"
	"html/template"
)

// companyData is given to a company template
type companyData struct {
	Company  company
	Media    media
	GameList map[string]string
}

// gameData is given to a game template
type gameData struct {
	companyData
	Game game
}

// render renders a template with given data
func render(tplName string, data interface{}) ([]byte, error) {
	var buf bytes.Buffer

	// read template
	tplData, err := Asset("templates/" + tplName + ".html")
	if err != nil {
		return nil, err
	}

	t, err := template.New(tplName).Parse(string(tplData))
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
func renderGame(c company, g game, m media) ([]byte, error) {
	return render("game", gameData{
		companyData: companyData{
			Company: c,
			Media:   m,
		},
		Game: g,
	})
}

// renderCompany renders a company html
func renderCompany(c company, m media, gl map[string]string) ([]byte, error) {
	return render("company", companyData{
		Company:  c,
		Media:    m,
		GameList: gl,
	})
}
