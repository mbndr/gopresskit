package presskit

import (
	"html/template"
)

const (
	// video embed strings
	youtubeEmbed = "<iframe src=\"https://www.youtube.com/embed/%s\" frameborder=\"0\" allowfullscreen></iframe>"
	vimeoEmbed   = "<iframe src=\"https://player.vimeo.com/video/%s?title=0&byline=0&portrait=0\" frameborder=\"0\" allowfullscreen></iframe>"
	htmlEmbed    = `
		<video controls>
			<source src="%s" type="video/%s">
			Your browser does not support HTML5 video.
		</video>`

	// video link strings
	youtubeLink = "https://www.youtube.com/watch?v="
	vimeoLink   = "https://vimeo.com/"
)

// video holds video information
type video struct {
	Title  string `xml:"title"`
	Source string `xml:"source"` // Youtube etc
	Id     string `xml:"id"`     // e.g. youtube/vimeo id or path to local video
}

// Embed returns the embed html of a video
func (v video) Embed() template.HTML {
	switch v.Source {
	case "youtube":
		return html(youtubeEmbed, v.Id)
	case "vimeo":
		return html(vimeoEmbed, v.Id)
	case "mp4":
		return html(htmlEmbed, v.Id, v.Source)
	}
	return html("<b>No valid video source</b>")
}

// Link returns the direct link to a video
func (v video) Link() template.HTML {
	switch v.Source {
	case "youtube":
		return html("<a href=\"%s%s\">Youtube</a>", youtubeLink, v.Id)
	case "vimeo":
		return html("<a href=\"%s%s\">Vimeo</a>", vimeoLink, v.Id)
	case "mp4":
		return html("<a href=\"%s\">MP4</a>", v.Id)
	}
	return html("")
}

// UseIframeContainer checks if for the video source an css class
// "iframe-container" has to be added to display the video properly
// TODO understand why (css)
func (v video) UseIframeContainer() bool {
	switch v.Source {
	case "youtube":
		return true
	case "vimeo":
		return true
	}
	return false
}
