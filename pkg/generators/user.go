package generators

import (
	gofakeit "github.com/brianvoe/gofakeit/v7"
	uuid "github.com/google/uuid"
	models "templ-demo/internal/domain/models"
)

func GenerateUser(t testingI, args map[string]any) models.User {
	fak := models.User{
		CreatedAt: gofakeit.Date(),
		Email:     gofakeit.Word(),
		Id:        uuid.New(),
		Password:  gofakeit.Word(),
		Status:    models.Active,
	}

	decode(t, args, &fak)

	return fak
}
