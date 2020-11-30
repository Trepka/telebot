package database

import (
	"database/sql"
	"fmt"
	"telebot/internal/models"

	_ "github.com/golang-migrate/migrate/v4/source/file" // required for go-migrate via files
	_ "github.com/lib/pq"                                // required for PostgreSQL connection
)

// PostgresWasteStorage incapsulates PostgreSQL storage
type TelebotLanguageStorage struct {
	db *sql.DB
}

// Migrate ups version of DB model
// func (t *TelebotLanguageStorage) Migrate() {
// 	driver, err := postgres.WithInstance(p.db, &postgres.Config{})
// 	if err != nil {
// 		log.Fatal("[MIGRATE] Unable to get driver due to: " + err.Error())
// 	}
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file:///app/migrations",
// 		"postgres", driver)
// 	if err != nil {
// 		log.Fatal("[MIGRATE] Unable to get migrate instance due to: " + err.Error())
// 	}
// 	err = m.Up()
// 	switch err {
// 	case migrate.ErrNoChange:
// 		return
// 	default:
// 		log.Fatal("[MIGRATE] Unable to apply DB migrations due to: " + err.Error())
// 	}
// }

func NewTelebotLanguageStorage(config *models.Config) *TelebotLanguageStorage {
	dbURL := config.DbURL
	if dbURL == "" {
		// dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
		dbURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		fmt.Printf(err.Error())
	}
	if err = db.Ping(); err != nil {
		fmt.Printf(err.Error())
	}

	storage := TelebotLanguageStorage{db: db}
	return &storage
}

func (t *TelebotLanguageStorage) GetAllRows() (MessageList, error) {
	rows, err := t.db.Query("SELECT * FROM telebot_language")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messageList := make(MessageList, 0)
	for rows.Next() {
		var telebotMessage TelebotMessage
		if err := rows.Scan(&telebotMessage.ID, &telebotMessage.RUS, &telebotMessage.EN); err != nil {
			return nil, err
		}

		// fmt.Printf(fmt.Sprintf("%v\n", telebotMessage))
		messageList = append(messageList, telebotMessage)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messageList, nil
}

// func (p *PostgresWasteStorage) GetWasteTypeByID(ctx context.Context, wasteTypeID string) (*models.WasteType, error) {
// 	var wt models.WasteType
// 	err := p.db.QueryRowContext(ctx, `SELECT id, name, description FROM waste_type WHERE id = $1;`, wasteTypeID).Scan(&wt.ID, &wt.Name, &wt.Description)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return nil, ErrNotFound
// 	case nil:
// 		return &wt, nil
// 	default:
// 		return nil, err
// 	}
// }
