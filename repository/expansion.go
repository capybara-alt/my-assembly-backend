package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Expansion interface {
	InsertBatch(context.Context, []model.Expansion) error
}