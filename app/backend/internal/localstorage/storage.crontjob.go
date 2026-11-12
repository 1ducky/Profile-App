package localstorage

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"profile-template/internal/config"
	"time"
)

type CronJob struct {
	conf config.StorageConfig
	TTL  time.Duration
}

func NewCronJob(conf config.StorageConfig) *CronJob {
	return &CronJob{conf: conf, TTL: 15 * time.Minute}
}

func (c *CronJob) CleanTemp(ctx context.Context) error {
	finalDestination, err := filepath.Abs(filepath.Join(c.conf.StorageRoot, c.conf.StoragePathTemp))
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(finalDestination)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if time.Since(info.ModTime()) > c.TTL {
			_ = os.Remove(path.Join(finalDestination, entry.Name()))
		}
	}
	return nil
}
