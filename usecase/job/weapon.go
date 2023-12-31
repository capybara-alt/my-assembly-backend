package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type WeaponJob struct {
	db_repo      repository.Weapon
	externalRepo []repository.ExternalWeapon
	convertor    convert.IConvertor[model.Weapon]
	logger       *slog.Logger
}

func NewWeaponJob(
	db_repo repository.Weapon,
	externalRepo []repository.ExternalWeapon,
	convertor convert.IConvertor[model.Weapon],
	logger *slog.Logger) ICrawlJobUsecase {
	return &WeaponJob{
		db_repo:      db_repo,
		externalRepo: externalRepo,
		convertor:    convertor,
		logger:       logger,
	}
}

func (c *WeaponJob) Execute(ctx context.Context) {
	models := []model.Weapon{}

	for _, repo := range c.externalRepo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "detail", err)
		}
		c.logger.Info("Crawl successful")
		weaponList, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, weaponList...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.db_repo.UpsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
