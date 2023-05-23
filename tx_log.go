package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Event struct {
	Sequence  uint64 // record ID
	EventType EventType
	key       string
	value     string
}

const (
	_                     = iota // ignore 0
	EventDelete EventType = iota // iota = 1
	EventPut                     // iota = 2
)

type EventType byte

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key, value string)
	Err() <-chan error

	ReadEvent() (<-chan Event, <-chan error)
	Run()
}

type PostgresTransactionLogger struct {
	events chan<- Event
	errors <-chan error
	db     *sql.DB
}

func (l *PostgresTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, key: key, value: value}
}

func (l *PostgresTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, key: key}
}

func (l *PostgresTransactionLogger) Err() <-chan error {
	return l.errors
}

func (l *PostgresTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		query := `INSERT INTO transactions
		(event_type, key, value)
		VALUES ($1, $2, $3)`
		for e := range events {
			_, err := l.db.Exec(query, e.EventType, e.key, e.value)

			if err != nil {
				errors <- err

			}
		}

	}()
}

func (l *PostgresTransactionLogger) ReadEvent() (<-chan Event, <-chan error) {

	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		defer close(outEvent)
		defer close(outError)

		query := `SELECT sequence, event_type, key, value FROM transactions ORDER BY sequence`
		rows, err := l.db.Query(query)

		if err != nil {
			outError <- fmt.Errorf("sql error :%w", err)
			return
		}

		defer rows.Close()

		e := Event{}

		for rows.Next() {
			err = rows.Scan(&e.Sequence, &e.EventType, &e.key, &e.value)

			if err != nil {
				outError <- fmt.Errorf("error reading rows :%w", err)
			}
			outEvent <- e
		}
		err = rows.Err()
		if err != nil {
			outError <- fmt.Errorf("transaction read failure :%w", err)
		}

	}()

	return outEvent, outError
}

func NewPostgresTransactionLogger(config PostgresDB) (TransactionLogger, error) {
	connString := fmt.Sprintf("host=%s dbname=%s user=%s password%s", config.host, config.dbName, config.user, config.password)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("fail open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("fail connection: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}

	return logger, nil
}

type PostgresDB struct {
	dbName   string
	host     string
	user     string
	password string
}
