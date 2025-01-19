package models

type Department struct {
	ID   string
	Name string
}

func NewDepartment(id, name string) *Department {
	return &Department{
		ID:   id,
		Name: name,
	}
}
