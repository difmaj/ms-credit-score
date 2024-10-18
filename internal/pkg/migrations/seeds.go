package migrations

import (
	"os"
	"path/filepath"

	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"gorm.io/gorm"
)

const seedsDir = "seeds/"

// RunSeeds executes the seeds.
func RunSeeds(db *gorm.DB) error {
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

			if err := db.Exec(string(script)).Error; err != nil {
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
