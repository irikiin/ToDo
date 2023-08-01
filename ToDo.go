package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TodoItem struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

type TodoList struct {
	Items []TodoItem `json:"items"`
}

func main() {
	todoList := loadTodoList()

	fmt.Println("ToDoリストアプリへようこそ！")

	for {
		showMenu()
		option := getUserInput()

		switch option {
		case 1:
			addTodoItem(&todoList)
		case 2:
			listTodoItems(todoList)
		case 3:
			markItemAsDone(&todoList)
		case 4:
			deleteTodoItem(&todoList)
		case 5:
			saveTodoList(todoList)
			fmt.Println("アプリを終了します。")
			return
		default:
			fmt.Println("無効な選択です。もう一度選んでください。")
		}
	}
}

func showMenu() {
	fmt.Println("1. タスクを追加する")
	fmt.Println("2. タスク一覧を表示する")
	fmt.Println("3. タスクを完了する")
	fmt.Println("4. タスクを削除する")
	fmt.Println("5. 終了する")
}

func getUserInput() int {
	var option int
	fmt.Print("選択肢を入力してください：")
	fmt.Scanln(&option)
	return option
}

func loadTodoList() TodoList {
	file, err := ioutil.ReadFile("todo_list.json")
	if err != nil {
		return TodoList{}
	}

	var todoList TodoList
	json.Unmarshal(file, &todoList)
	return todoList
}

func saveTodoList(todoList TodoList) {
	data, err := json.MarshalIndent(todoList, "", "    ")
	if err != nil {
		fmt.Println("データの保存に失敗しました。")
		return
	}

	err = ioutil.WriteFile("todo_list.json", data, 0644)
	if err != nil {
		fmt.Println("ファイルの保存に失敗しました。")
		return
	}
}

func addTodoItem(todoList *TodoList) {
	var title string
	fmt.Print("タスク名を入力してください：")
	fmt.Scanln(&title)

	item := TodoItem{Title: title, IsDone: false}
	todoList.Items = append(todoList.Items, item)

	fmt.Println("タスクを追加しました。")
}

func listTodoItems(todoList TodoList) {
	if len(todoList.Items) == 0 {
		fmt.Println("タスクはありません。")
		return
	}

	fmt.Println("ToDoリスト：")
	for i, item := range todoList.Items {
		status := "未完了"
		if item.IsDone {
			status = "完了"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, item.Title)
	}
}

func markItemAsDone(todoList *TodoList) {
	listTodoItems(*todoList)

	var index int
	fmt.Print("完了にするタスクの番号を入力してください：")
	fmt.Scanln(&index)

	if index <= 0 || index > len(todoList.Items) {
		fmt.Println("無効な番号です。")
		return
	}

	todoList.Items[index-1].IsDone = true
	fmt.Println("タスクを完了しました。")
}

func deleteTodoItem(todoList *TodoList) {
	listTodoItems(*todoList)

	var index int
	fmt.Print("削除するタスクの番号を入力してください：")
	fmt.Scanln(&index)

	if index <= 0 || index > len(todoList.Items) {
		fmt.Println("無効な番号です。")
		return
	}

	todoList.Items = append(todoList.Items[:index-1], todoList.Items[index:]...)
	fmt.Println("タスクを削除しました。")
}
