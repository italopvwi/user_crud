package repositories

import (
	"context"
	"database/sql"
	"user_crud/pkg/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, birthday) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Birthday)
	return err
}

func (r *postgresUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, first_name, last_name, birthday FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Birthday)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postgresUserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, first_name, last_name, birthday FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Birthday)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *postgresUserRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, birthday = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Birthday, user.ID)
	return err
}

func (r *postgresUserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
