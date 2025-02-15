package structs

type TodoItem struct {
	Id         string
	Item       string
	Item_order int
}

type TodoItemList struct {
	Items []TodoItem
	Count int
}
