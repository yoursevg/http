package messagesService

import (
	"gorm.io/gorm"
)

type MessageRepository interface {
	// CreateMessage - Передаем в функцию message типа Message из orm.go
	// возвращаем созданный Message и ошибку
	CreateMessage(message Message) (Message, error)
	// GetAllMessages - Возвращаем массив из всех писем в БД и ошибку
	GetAllMessages() ([]Message, error)
	// UpdateMessageByID - Передаем id и Message, возвращаем обновленный Message и ошибку
	UpdateMessageByID(id int, message Message) (Message, error)
	// DeleteMessageByID - Передаем id для удаления, возвращаем только ошибку
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// CreateMessage (r *messageRepository) привязывает данную функцию к нашему репозиторию
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// GetAllMessages - Получение всего списка сообщений
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

// UpdateMessageByID - Ищем запись в БД по id и обновляем поле message
func (r *messageRepository) UpdateMessageByID(id int, newMessage Message) (Message, error) {
	var msg Message
	msg.ID = uint(id)
	err := r.db.First(&msg).Error
	if err != nil {
		return Message{}, err
	}
	msg.Text = newMessage.Text
	err = r.db.Save(&msg).Error
	if err != nil {
		return Message{}, err
	}

	return msg, err
}

// DeleteMessageByID - Ищем запись в БД по id и удаляем её
func (r *messageRepository) DeleteMessageByID(id int) error {
	var messageToDelete Message
	messageToDelete.ID = uint(id)
	err := r.db.First(&messageToDelete, id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&messageToDelete).Error
	return err
}
