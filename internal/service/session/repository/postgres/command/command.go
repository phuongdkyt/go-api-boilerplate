package sessioncommand

import (
	"context"
	"time"

	"go.uber.org/fx"
	"gorm.io/gorm/clause"

	"github.com/anhnmt/go-api-boilerplate/gen/gormgen"
	sessionentity "github.com/anhnmt/go-api-boilerplate/internal/model"
)

type Command struct {
	db *gormgen.Query
}

type Params struct {
	fx.In

	DB *gormgen.Query
}

func New(p Params) *Command {
	return &Command{
		db: p.DB,
	}
}

func (c *Command) CreateOnConflict(ctx context.Context, session *sessionentity.Session) error {
	return c.db.WriteDB().Session.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_seen_at", "expires_at", "updated_at"}),
	}).Create(session)
}

func (c *Command) UpdateIsRevoked(ctx context.Context, sessionId string, isRevoked bool, now time.Time) error {
	e := c.db.Session

	_, err := c.db.WriteDB().Session.WithContext(ctx).Where(e.ID.Eq(sessionId)).
		Updates(map[string]interface{}{
			"is_revoked":   isRevoked,
			"last_seen_at": now,
			"updated_at":   now,
		})
	return err
}

func (c *Command) UpdateLastSeenAt(ctx context.Context, sessionId string, now time.Time) error {
	e := c.db.Session

	_, err := c.db.WriteDB().Session.WithContext(ctx).Where(e.ID.Eq(sessionId)).
		Updates(map[string]interface{}{
			"last_seen_at": now,
			"updated_at":   now,
		})
	return err
}

func (c *Command) UpdateRevokedByUserIdWithoutSessionId(ctx context.Context, userId, sessionId string) error {
	return c.db.WriteDB().Session.WithContext(ctx).
		UpdateRevokedByUserIdWithoutSessionId(userId, sessionId)
}
