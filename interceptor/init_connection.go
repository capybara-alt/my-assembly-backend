package interceptor

import (
	"context"
	"database/sql"
	"os"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func NewInitConnectionInterceptor() connect.UnaryInterceptorFunc {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
			db, err := gorm.Open(postgres.New(postgres.Config{
				DSN:                  core.DB_DSN,
				PreferSimpleProtocol: true,
			}), &gorm.Config{
				Logger: gorm_logger.New(slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo), gorm_logger.Config{
					LogLevel: gorm_logger.Info,
					Colorful: true,
				}),
			})
			if err != nil {
				return nil, err
			}
			conn, err := db.DB()
			if err != nil {
				return nil, err
			}
			defer func(conn *sql.DB) {
				logger.Info("Close DB Connection")
				if err := conn.Close(); err != nil {
					logger.Error("DB Connection close error", "detail", err)
				}
			}(conn)

			err = db.Transaction(func(tx *gorm.DB) error {
				resp, err = next(core.SetTx(ctx, tx), req)
				if err != nil {
					return err
				}
				return nil
			})

			return resp, err
		})
	}
}
