package account

import (
	"context"

	activitymodel "github.com/Retrospective53/myGram/module/models"
)

type IAccountActivityRepo interface {
	CreateActivity(ctx context.Context, acc activitymodel.AccountActivity) (created activitymodel.AccountActivity, err error)
}