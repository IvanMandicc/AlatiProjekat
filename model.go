package main

type Config struct {
	Id      string            `json:"id"`
	Entries map[string]string `json:"entries"`
	Title   string            `json:"title"`
	Text    string            `json:"text"`
	Tags    []string          `json:"tags"`
}

type Service struct {
	Data map[string]*[]Config
}

type Group struct {
	Id      string   `json:"id"`
	Configs []Config `json:"configs"`
}
