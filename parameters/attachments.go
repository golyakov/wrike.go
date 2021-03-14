package parameters

type Date struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type QueryTaskAttachments struct {
	Versions  *bool `url:"versions,omitempty"`
	StartDate *Date `url:"startDate,omitempty"`
	WithUrls  *bool `url:"withUrls,omitempty"`
}
