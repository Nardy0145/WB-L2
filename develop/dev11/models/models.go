package models

type Event struct {
	Result struct {
		Id       string `json:"id"`
		Datetime string `json:"datetime"`
		Event    string `json:"event"`
	} `json:"result"`
}
type Events []struct {
	Result struct {
		Id       string `json:"id"`
		Datetime string `json:"datetime"`
		Event    string `json:"event"`
	} `json:"result"`
}

type Success struct {
	Result struct {
		Result string `json:"result"`
	}
}

type Err struct {
	Error struct {
		Error string `json:"error"`
	}
}
