package migrations

import (
	"context"
	"os"
	"path/filepath"

	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/jmoiron/sqlx"
)

const seedsDir = "seeds/"

// RunSeeds executes the seeds.
func RunSeeds(ctx context.Context, db *sqlx.DB) error {
	enviromentDir := seedsDir + config.Env.Environment
	if _, err := os.Stat(enviromentDir); os.IsNotExist(err) {
		return nil
	}

	err := filepath.Walk(enviromentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			script, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if _, err := db.ExecContext(ctx, string(script)); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
