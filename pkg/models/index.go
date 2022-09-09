package models

type Model struct {
	Model any
}





//Ugly way to register Models for Postgre :)
func RegisterModels() []Model {
	return []Model{
		{Model: Job{}},
	}
}
