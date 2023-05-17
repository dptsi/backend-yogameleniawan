package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/dptsi/backend-yogameleniawan/06-go-database-mysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"

	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = id

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ?"

	comment := entity.Comment{}

	rows, err := repository.DB.QueryContext(ctx, script, id)

	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.FormatInt(id, 10) + " not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"

	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}

		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
