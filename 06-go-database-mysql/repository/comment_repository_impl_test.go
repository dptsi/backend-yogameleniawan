package repository

import (
	"context"
	"fmt"
	"testing"

	go_database_mysql "github.com/dptsi/backend-yogameleniawan/06-go-database-mysql"
	"github.com/dptsi/backend-yogameleniawan/06-go-database-mysql/entity"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
