package domain

type SignHmacSha512 struct {
	Text string `json:"text" validate:"required,gte=1"`
	Key  string `json:"key" validate:"required,gte=1"`
}

func (s *SignHmacSha512) Validate() error {
	return validate.Struct(s)
}
