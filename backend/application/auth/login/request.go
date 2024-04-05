package login

type validationErrors map[string]string

type Request struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

func (r *Request) Validate() (bool, validationErrors) {
	return true, nil
}
