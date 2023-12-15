package syncer

import (
	"context"

	"github.com/TrevorEdris/retropie-utils/pkg/errors"
	"github.com/TrevorEdris/syncer/pkg/config"
)

type (
	Syncer interface {
		Sync(ctx context.Context) error
	}

	syncer struct {
		cfg SyncConfig
	}

	SyncConfig struct{
        config.Config
    }

	Schedule struct{}
)

func NewSyncer(cfg SyncConfig) (Syncer, error) {
	return &syncer{cfg}, nil
}

func (s *syncer) Sync(ctx context.Context) error {
	return errors.NotImplementedError
}
