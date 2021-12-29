package e

type ApiError struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err ApiError) Error() string {
	return err.Message
}
