package db

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -o ./mocks/ -s ".go"

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client - описывает БД клиент
type Client interface {
	DB() DB
	Close() error
}

// Query - структура, описывающая SQL запрос и тип запроса для логирования действий
type Query struct {
	Name     string
	QueryRaw string
}

// Tx - описывает методы для работы с транзакциями
type Tx interface {
	Begin(ctx context.Context) (pgx.Tx, error)

	BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error)

	Commit(ctx context.Context) error

	Rollback(ctx context.Context) error

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects

	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)

	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)

	Conn() *pgx.Conn
}

// Row - описывает метод получения данных при запросе в БД из одной строки
type Row interface {
	Scan(dest ...interface{}) error
}

// Handler - функция-обертка над транзакциями
type Handler func(ctx context.Context) error

// TxManager - интерфейс, описывающий методы дял работы с транзакциями разного уровня
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Transactor - интерфейс, описывающий методы дял работы с функциями транзакций pgx
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (Tx, error)
}

// SQLExecer - описывает интерфейсы методов для создания запросов к БД
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer - интерфейс, описывающий методы-обертки надм методами QueryExecer с рекурсивным сканированием моделей данных
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer - интерфейс, описывающий методы для создания методов-оберток над стандартными функциями pgx
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger - описывает метод Ping для теста подключения к БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB - интерфейс, описывающий методы для работы с БД
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
