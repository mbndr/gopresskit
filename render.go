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

// MonetizeText returns the monetization permission text depending on the config
// TODO a bit messy
func (gd gameData) MonetizeText() template.HTML {
	switch gd.Game.Monetize {
	case "false":
		return html("%s does currently not allow for the contents of %s to be published through video broadcasting services.", gd.Company.Title, gd.Game.Title)
	case "non-commercial":
		return html("%s allows for the contents of %s to be published through video broadcasting services for non-commercial purposes only. Monetization of any video created containing assets from %s is not allowed.", gd.Company.Title, gd.Game.Title, gd.Game.Title)
	case "monetize":
		return html("%s allows for the contents of %s to be published through video broadcasting services for any commercial or non-commercial purposes. Monetization of videos created containing assets from %s is legally & explicitly allowed by %s. This permission can be found in <a href=\"#contact\">writing at %s</a>.", gd.Company.Title, gd.Game.Title, gd.Game.Title, gd.Company.Title, gd.Company.Title)
	default: // originally "ask", bit if something fails, you should always ask
		return html("%s does allow the contents of this game to be published through video broadcasting services only with direct written permission from <a href=\"#contact\">%s</a>.", gd.Company.Title, gd.Company.Title)
	}
}
