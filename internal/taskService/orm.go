package taskService

import "gorm.io/gorm"

type Task struct {
	//Эмбеддинг стандартной модели GORM
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}
