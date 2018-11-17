package models

// Summary type for the summary API call
type Summary struct {
	TotalGrand      int `json:"total_grand"`
	TotalBillable   int `json:"total_billable"`
	TotalCurrencies []struct {
		Currency string  `json:"currency"`
		Amount   float32 `json:"amount"`
	} `json:"total_currencies"`
	Data []struct {
		ID    int `json:"id"`
		Title struct {
			Project  string `json:"project"`
			Client   string `json:"client"`
			Color    string `json:"color"`
			HexColor string `json:"hex_color"`
		} `json:"title"`
		Time            int `json:"time"`
		TotalCurrencies []struct {
			Currency string  `json:"currency"`
			Amount   float32 `json:"amount"`
		} `json:"total_currencies"`
		Items []struct {
			Title struct {
				TimeEntry string `json:"time_entry"`
			} `json:"title"`
			Time int     `json:"time"`
			Cur  string  `json:"cur"`
			Sum  float32 `json:"sum"`
			Rate float32 `json:"rate"`
		} `json:"items"`
	} `json:"data"`
}
