package main

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("GoDo.json")
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	storage.Save(todos)
}
