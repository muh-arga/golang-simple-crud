package model

type Todo struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Completed bool   `json:"completed" form:"completed"`
}

var Todos = []*Todo{
	{Id: 1, Name: "Makan Malam", Completed: false},
	{Id: 2, Name: "Makan Siang", Completed: false},
}