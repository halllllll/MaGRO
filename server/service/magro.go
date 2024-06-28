package service

import (
	"context"
	"fmt"

	"github.com/halllllll/MaGRO/entity"
)

type MagroSystem struct {
	Repo Infoer
}

func (m *MagroSystem) GetSystemInfo(ctx context.Context) (*entity.System, error) {

	info, err := m.Repo.GetSystemInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to info: %w", err)
	}

	return info, nil
}

type MutateMAGRO struct {
	Repo MaGROMutater
}

func (m *MutateMAGRO) UpdateRole(ctx context.Context, roles *entity.ReqNewRoleAlias) error {
	a := string(roles.AliasAdminName)
	d := string(roles.AliasDirectorName)
	mm := string(roles.AliasManagerName)
	g := string(roles.AliasGuestName)
	if a == "" && d == "" && mm == "" && g == "" {
		return fmt.Errorf("no content")
	}

	err := m.Repo.UpdateRole(ctx, roles)
	if err != nil {
		return err
	}
	return nil
}
