package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"blog/internal/config"

	_ "modernc.org/sqlite"
)

var (
	db     *sql.DB
	once   sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		cfg := config.GlobalConfig
		if cfg == nil {
			cfg = &config.Config{DBPath: "./data/blog.db"}
		}

		var err error
		db, err = sql.Open("sqlite", cfg.DBPath)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)

		if err := initSchema(); err != nil {
			log.Fatalf("Failed to initialize schema: %v", err)
		}

		fmt.Println("Database connection established and schema initialized")
	})
	return db
}

func initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		nickname TEXT,
		avatar TEXT
	);

	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		sort INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		title TEXT NOT NULL,
		slug TEXT UNIQUE,
		content TEXT NOT NULL,
		summary TEXT,
		cover_image TEXT,
		status TEXT DEFAULT 'draft',
		view_count INTEGER DEFAULT 0,
		category_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		article_id INTEGER NOT NULL,
		nickname TEXT NOT NULL,
		email TEXT,
		content TEXT NOT NULL,
		ip TEXT,
		status TEXT DEFAULT 'pending',
		FOREIGN KEY (article_id) REFERENCES articles(id)
	);

	CREATE TABLE IF NOT EXISTS article_tags (
		article_id INTEGER,
		tag_id INTEGER,
		PRIMARY KEY (article_id, tag_id),
		FOREIGN KEY (article_id) REFERENCES articles(id),
		FOREIGN KEY (tag_id) REFERENCES tags(id)
	);
	`

	_, err := db.Exec(schema)
	return err
}

func InitDefaultData() error {
	db := GetDB()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = db.Exec(`
			INSERT INTO categories (name, slug, sort) VALUES
			('技术', 'tech', 1),
			('生活', 'life', 2),
			('随笔', 'essay', 3)
		`)
		if err != nil {
			return fmt.Errorf("failed to create categories: %v", err)
		}
		fmt.Println("Default categories created")
	}

	return nil
}
