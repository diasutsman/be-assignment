package responses

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}
