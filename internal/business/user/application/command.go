package application

import (
	"context"

	"github.com/go-jimu/components/sloghelper"
	"github.com/jacexh/linear-app-intergation/internal/business/user/domain"
	"golang.org/x/exp/slog"
)

type CommandChangePasswordHandler struct {
	repo domain.Repository
}

func NewCommandChangePasswordHandler(repo domain.Repository) *CommandChangePasswordHandler {
	return &CommandChangePasswordHandler{
		repo: repo,
	}
}

func (h *CommandChangePasswordHandler) Handle(ctx context.Context, logger *slog.Logger, command *CommandChangePassword) error {
	entity, err := h.repo.Get(ctx, command.ID)
	if err != nil {
		logger.ErrorCtx(ctx, "failed to get user password", sloghelper.Error(err))
		return err
	}
	if err = entity.ChangePassword(command.OldPassword, command.NewPassword); err != nil {
		logger.ErrorCtx(ctx, "failed to change password", sloghelper.Error(err))
		return err
	}
	if err = h.repo.Save(ctx, entity); err != nil {
		logger.ErrorCtx(ctx, "failed to save new password", sloghelper.Error(err))
		return err
	}
	logger.InfoCtx(ctx, "password is changed")
	return nil
}
