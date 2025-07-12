package responses

type GetConfigDatesResponse struct {
	Months []int `json:"months"`
	Years  []int `json:"years"`
}
