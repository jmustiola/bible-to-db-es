package main

type Verse struct {
	VersNum int32  `json:"verNum"`
	Text    string `json:"text"`
}

type Chapter struct {
	NumChapter  int32   `json:"numChapter"`
	TotalVerses int32   `json:"totalVer"`
	Verses      []Verse `json:"vers"`
}

type Book struct {
	Name          string    `json:"name"`
	TotalVerses   int32     `json:"totalVers"`
	TotalChapters int32     `json:"totalChapters"`
	Chapters      []Chapter `json:"chapters"`
}
