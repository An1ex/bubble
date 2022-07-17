package db

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Create(todo Todo) error {
	err := DB.Create(&todo).Error
	return err
}

func ReadAll() (*[]Todo, error) {
	var todoList []Todo
	if err := DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return &todoList, nil
}

func Update(id string) error {
	var todo Todo
	err := DB.Where("id = ?", id).First(&todo).Update("status", !todo.Status).Error
	return err
}

func Delete(id string) error {
	var todo Todo
	err := DB.Where("id = ?", id).First(&todo).Delete(todo).Error
	return err
}
