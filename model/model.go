package model

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var Todos = []*Todo{
	{Id: 1, Name: "Makan Malam", Completed: false},
	{Id: 2, Name: "Makan Siang", Completed: false},
}