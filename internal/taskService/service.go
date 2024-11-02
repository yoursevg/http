package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) CreateTask(message Task) (Task, error) {
	return s.repo.CreateTask(message)
}

func (s *TaskService) UpdateTaskByID(id uint, message Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, message)
}

func (s *TaskService) DeleteTaskByID(id int) error {
	return s.repo.DeleteTaskByID(id)
}
