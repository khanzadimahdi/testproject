package createelement

type Response struct {
	ValidationErrors validationErrors `json:"errors,omitempty"`
	UUID             string           `json:"uuid,omitempty"`
}
