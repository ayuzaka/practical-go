package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// トランザクションを制御するための構造体
type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

type PgTable struct {
	SchemaName string
	TableName  string
}

var _ pgx.Logger = (*logger)(nil)

type logger struct{}

var db *sql.DB

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if msg == "Query" {
		log.Printf("SQL:\n%v\nARGS:%v\n", data["sql"], data["args"])
	}
}

func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}

	return tx.Commit()
}

func (s *Service) UpdateProduct(ctx context.Context, productID string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "..."); err != nil {
			return err
		}

		if _, err := s.tx.ExecContext(ctx, "..."); err != nil {
			return err
		}

		return nil
	}

	return s.tx.Transaction(ctx, updateFunc)
}

func fetchUsers() {
	ctx := context.TODO()
	rows, err := db.QueryContext(ctx, `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        time.Time
		)

		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", err)
		}

		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}

	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users: %v", err)
	}
}

func FetchUser(ctx context.Context, userID string) (*User, error) {
	row := db.QueryRowContext(ctx, `SELECT user_id, user_name FROM users WHERE user_id = $1;`, userID)
	user, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return user, nil
}

func scanUser(row *sql.Row) (*User, error) {
	var u User
	err := row.Scan(&u.UserID, &u.UserName)
	if err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}

	return &u, nil
}

func FetchTables(ctx context.Context) {
	config, err := pgx.ParseConfig("user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("parse config: %v\n", err)
	}
	config.Logger = &logger{}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	rows, err := conn.Query(ctx, sql, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pgtables []PgTable
	for rows.Next() {
		var s, t string
		if err := rows.Scan(&s, &t); err != nil {
			log.Fatal(err)
		}
		pgtables = append(pgtables, PgTable{SchemaName: s, TableName: t})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	db, _ = sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
}

func main() {
	// ctx := context.Background()
}
