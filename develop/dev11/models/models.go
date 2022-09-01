package models

// Event ../
type Event struct {
	Result struct {
		Id       string `json:"id"`
		Datetime string `json:"datetime"`
		Event    string `json:"event"`
	} `json:"result"`
}

// Events ../
type Events []struct {
	Result struct {
		Id       string `json:"id"`
		Datetime string `json:"datetime"`
		Event    string `json:"event"`
	} `json:"result"`
}

// Success ../
type Success struct {
	Result struct {
		Result string `json:"result"`
	}
}

// Err ../
type Err struct {
	Error struct {
		Error string `json:"error"`
	}
}
