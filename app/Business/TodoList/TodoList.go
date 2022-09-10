package TodoList

var TodoList map[int64][]Todo = make(map[int64][]Todo)

type Todo struct {
	Content string `json:"content"`
	Status  int64  `json:"status"`
	CreatAt int64  `json:"creat_at"`
	DueDate int64  `json:"due_date"`
}

var TodoStatus map[int64]string = map[int64]string{
	1: "執行中",
	2: "完成",
	3: "刪除",
	4: "封存",
}

func init() {

}

// 依狀態列出待辦事項
func ListByStatus(userID, status int64) (list []Todo) {
	for i := range TodoList[userID] {
		if TodoList[userID][i].Status == status {
			list = append(list, TodoList[userID][i])
		}
	}
	return
}

// 新增待辦事項
func AddTodo(userID int64, addedTodo Todo) {
	TodoList[userID] = append(TodoList[userID], addedTodo)
	WriteTodoListFile()
}

// 刪除待辦事項
func DeleteTodo(userID int64, deleteCreatAt int64) {
	for i := range TodoList[userID] {
		if TodoList[userID][i].CreatAt == deleteCreatAt {
			TodoList[userID][i].Status = 3
			break
		}
	}
}

// 編輯待辦事項
func EditTodo(userID int64, editedTodo Todo) {
	for i := range TodoList[userID] {
		if TodoList[userID][i].CreatAt == editedTodo.CreatAt {
			TodoList[userID][i] = editedTodo
			break
		}
	}
}

func SelectTodo(userID, creatAt int64) {

}
