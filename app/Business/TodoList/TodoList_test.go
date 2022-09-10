package TodoList

import "testing"

func TestAddTodo(t *testing.T) {
	var todo Todo = Todo{
		Content: "origin",
		Status:  1,
		CreatAt: 0,
		DueDate: 0,
	}

	AddTodo(0, todo)

	if len(TodoList[0]) > 0 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestListByStatus(t *testing.T) {
	n := ListByStatus(0, 1)

	if len(n) > 0 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestDeleteTodo(t *testing.T) {
	DeleteTodo(0, 0)

	if TodoList[0][0].Status == 3 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestEditTodo(t *testing.T) {
	var todo Todo = Todo{
		Content: "after edit",
		Status:  1,
		CreatAt: 0,
		DueDate: 0,
	}

	EditTodo(0, todo)

	if TodoList[0][0].Status == 1 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
