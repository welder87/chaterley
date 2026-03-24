package main

import (
	"chaterley/internal/app/chat/entities"
	interRepo "chaterley/internal/app/chat/entities/repositories"
	"chaterley/internal/app/core"
	"chaterley/internal/app/db"
	infraRepo "chaterley/internal/infrastructure/chat/entities/repositories"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Инициалиазация коннекта к БД sqllite
	database := db.GetDBCon()
	// Закрытие коннекта при остановке приложения
	defer database.Close()

	// Создание репозитория для сущности Message
	var repository interRepo.Repository[entities.Message] = infraRepo.NewMessageRepository(database)

	// Генерация вспомогательных данных для доменной модели Message
	authorID := core.NewEntityID()
	content := "test text"
	contentType := "text"

	// Создание доменной модели Message
	message := entities.NewMessage(
		authorID,
		content,
		contentType,
	)

	ctx := context.Background()

	// Сохранение сущности Message
	errSave := repository.Save(ctx, message)
	if errSave != nil {
		panic("Error to save message entity.")
	}

	// Получение сущности Message
	messageDTO := message.ToSnapshot()
	messageID, _ := uuid.Parse(messageDTO.ID)
	messageEntity, errGet := repository.Get(ctx, core.NewExistsEntityID(messageID))
	if errGet != nil {
		panic("Errot to get message entity.")
	}
	fmt.Println(messageEntity.ToSnapshot().ID)

	// Удаление сущности Message
	errRemove := repository.Remove(ctx, message)
	if errRemove != nil {
		panic("Error to delete message entity.")
	}

	fmt.Println("Success!")
}
