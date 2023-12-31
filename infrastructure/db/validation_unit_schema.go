package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type validationUnitSchema struct{}

func NewValidationUnitSchema() repository.ValidationUnitSchema {
	return &validationUnitSchema{}
}

func (v *validationUnitSchema) GetValidationSchema(ctx context.Context, unit_type model.PrimaryUnitType) (model.ValidationUnitSchemaList, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	list := model.NewValidationUnitSchemaList()
	db.Where("unit_type = ?", unit_type).Find(&list)

	return list, nil
}
