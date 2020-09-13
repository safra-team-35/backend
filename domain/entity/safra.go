package entity

type SafraMorningCallsResponse struct {
	Data []struct {
		Channel     string `json:"channel"`
		Data        string `json:"data"`
		Description string `json:"description"`
		ID          string `json:"id"`
		Links       []struct {
			Href  string `json:"href"`
			Rel   string `json:"rel"`
			Title string `json:"title"`
		} `json:"links"`
		Playlist string `json:"playlist"`
		Title    string `json:"title"`
	} `json:"data"`
}
