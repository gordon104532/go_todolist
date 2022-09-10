package TodoList

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"

	"main/app/ErrorHandle"
	"os"
)

// 初始化待辦事項存檔
func InitTodoListFile() {
	filename := "todoList.txt"

	if _, err := os.Stat(filename); err == nil {
		// 有檔則讀取

		// open the file
		file, err := os.Open(filename)
		if err != nil {
			ErrorHandle.Error.Printf("%s 開啟失敗 err:%v\n", filename, err)
		}

		fileScanner := bufio.NewScanner(file)

		// read line by line
		var tempStr string
		for fileScanner.Scan() {
			tempStr = tempStr + fileScanner.Text()
		}

		err = json.Unmarshal([]byte(tempStr), &TodoList)
		if err != nil {
			ErrorHandle.Error.Printf("%s 內容解析失敗 err:%v\n", filename, err)
		}

		// handle first encountered error while reading
		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error while reading file: %s", err)
		}

		file.Close()

	} else if errors.Is(err, os.ErrNotExist) {
		// 無檔則建立
		_, err := os.Create(filename)
		if err != nil {
			ErrorHandle.Error.Println("ERROR", "CreateFile: 建立檔案錯誤, "+err.Error())
		}

		ErrorHandle.Info.Println("建立檔案:" + filename)

	} else {
		ErrorHandle.Info.Printf("InitIndexFile else err: %v\n", err)
	}
}

// 寫入待辦事項
func WriteTodoListFile() {
	msgJSON, _ := json.Marshal(TodoList)
	err := os.WriteFile("todoList.txt", msgJSON, 0644)
	if err != nil {
		ErrorHandle.Error.Printf("寫回事項失敗 內容如下:\n %v", TodoList)
	}
}
