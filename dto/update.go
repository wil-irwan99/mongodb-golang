package dto

type UpdateForm struct {
	Id    string                 `json:"id"`
	Table string                 `json:"table"`
	Value map[string]interface{} `json:"value"`
}
