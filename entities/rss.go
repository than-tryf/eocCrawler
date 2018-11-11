package entities

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Base    string   `xml:"base,attr"`
	Dc      string   `xml:"dc,attr"`
	Atom    string   `xml:"atom,attr"`
	Og      string   `xml:"og,attr"`
	Article string   `xml:"article,attr"`
	Book    string   `xml:"book,attr"`
	Profile string   `xml:"profile,attr"`
	Video   string   `xml:"video,attr"`
	Product string   `xml:"product,attr"`
	Content string   `xml:"content,attr"`
	Foaf    string   `xml:"foaf,attr"`
	Rdfs    string   `xml:"rdfs,attr"`
	Sioc    string   `xml:"sioc,attr"`
	Sioct   string   `xml:"sioct,attr"`
	Skos    string   `xml:"skos,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Item        []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Creator     string `xml:"creator"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}
