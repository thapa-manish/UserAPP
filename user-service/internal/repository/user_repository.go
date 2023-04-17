package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/Masterminds/squirrel"

	"use/internal/model"
)

type IUserRepository interface {
	FindAll(uint64, uint64) ([]model.User, error)
	FindByID(int64) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	Save(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(int64) error
}

type UserRepository struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
	id           int64
	lock         sync.Mutex
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *UserRepository) FindAll(page, perPage uint64) ([]model.User, error) {
	if perPage == 0 {
		perPage = 10
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * perPage
	query, _, _ := r.queryBuilder.Select("*").From("users").Limit(perPage).Offset(offset).ToSql()
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.UserStatus, &user.Department); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rows: %v", err)
	}
	return users, nil
}

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	query, arg, _ := r.queryBuilder.Select("*").From("users").Where(squirrel.Eq{"id": id}).ToSql()
	row := r.db.QueryRow(query, arg...)

	var user model.User
	if err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.UserStatus, &user.Department); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query, args, _ := r.queryBuilder.Select("*").From("users").Where(squirrel.Eq{"email": email}).ToSql()
	row := r.db.QueryRow(query, args...)

	var user model.User
	if err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.UserStatus, &user.Department); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}
	return &user, nil
}

func (r *UserRepository) Save(user *model.User) (*model.User, error) {
	user.ID = r.getNewId()
	query, args, _ := r.queryBuilder.Insert("users").
		Columns("id", "user_name", "email", "first_name", "last_name", "user_status", "department").
		Values(user.ID, user.UserName, user.Email, user.FirstName, user.LastName, user.UserStatus, user.Department).ToSql()
	_, err := r.db.Exec(query, args...)
	fmt.Printf("\nquery:%+v args: %+v\n", query, args)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}
	return user, err
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	queryBuilder := r.queryBuilder.Update("users")

	if user.Email != "" {
		queryBuilder = queryBuilder.Set("email", user.Email)
	}

	if user.UserName != "" {
		queryBuilder = queryBuilder.Set("user_name", user.UserName)
	}

	if user.FirstName != "" {
		queryBuilder = queryBuilder.Set("first_name", user.FirstName)
	}

	if user.LastName != "" {
		queryBuilder = queryBuilder.Set("last_name", user.LastName)
	}

	if user.UserStatus != "" {
		queryBuilder = queryBuilder.Set("user_status", user.UserStatus)
	}

	if user.Department != "" {
		queryBuilder = queryBuilder.Set("department", user.Email)
	}

	query, args, err := queryBuilder.Where(squirrel.Eq{"id": user.ID}).ToSql()
	fmt.Println("query: ", query, args)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(id int64) error {
	query, args, err := r.queryBuilder.Delete("users").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) getNewId() int64 {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.id == 0 {
		query, _, _ := r.queryBuilder.Select("id").From("users").OrderBy("id", "desc").Limit(1).ToSql()
		res := r.db.QueryRow(query)
		if err := res.Scan(&r.id); err != nil {
			r.id = 0
		}
	}
	r.id = r.id + 1
	return r.id
}
