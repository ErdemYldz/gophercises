package main

type chapters map[string]Chapter

// Chapter holds a chapter in the file
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

// Options holds the options the user may choose
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
