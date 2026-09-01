package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type PollRepository struct {
	db *pgxpool.Pool
}

func NewPollRepository(db *pgxpool.Pool) *PollRepository {
	return &PollRepository{
		db: db,
	}
}

func (r *PollRepository) CreatePoll(
	ctx model.Context,
)