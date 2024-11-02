package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(message Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех заданий в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task и ошибку
	UpdateTaskByID(id uint, message Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// CreateTask - Создает новую запись в бд
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// GetAllTasks - Получение всего списка сообщений
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// UpdateTaskByID - Ищем запись в БД по id и обновляем
func (r *taskRepository) UpdateTaskByID(id uint, newTask Task) (Task, error) {
	var task Task
	task.ID = id
	err := r.db.First(&task).Error
	if err != nil {
		return Task{}, err
	}
	task.Text = newTask.Text
	if newTask.IsDone != task.IsDone {
		task.IsDone = newTask.IsDone
	}
	err = r.db.Save(&task).Error
	if err != nil {
		return Task{}, err
	}

	return task, err
}

// DeleteTaskByID - Ищем запись в БД по id и удаляем её
func (r *taskRepository) DeleteTaskByID(id int) error {
	var taskToDelete Task
	taskToDelete.ID = uint(id)
	err := r.db.First(&taskToDelete, id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&taskToDelete).Error
	return err
}
