package requests

type RegisterRequest struct {
	Type string `json:"type"  validate:"oneof=credit debit loan"`
}


func (a *RegisterRequest) GetType() int {
	m := map[string]int{
		"credit": 0,
		"debit":  1,
		"loan":   2,
	}
	return m[a.Type]
}
