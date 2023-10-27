package domain

type RedisInput struct {
	Key   string `json:"key" validate:"required,gte=1"`
	Value int    `json:"value" validate:"required,gte=1"`
}

func (r *RedisInput) Validate() error {
	return validate.Struct(r)
}

type RedisOutput struct {
	Value int `json:"value"`
}
