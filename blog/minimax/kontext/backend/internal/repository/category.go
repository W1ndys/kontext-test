package repository

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Sort      int       `json:"sort"`
}

func ScanCategory(row *sql.Row) (*Category, error) {
	c := &Category{}
	err := row.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.Name, &c.Slug, &c.Sort)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ScanCategoryRows(rows *sql.Rows) ([]*Category, error) {
	var categories []*Category
	for rows.Next() {
		c := &Category{}
		err := rows.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.Name, &c.Slug, &c.Sort)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: GetDB()}
}

func (r *CategoryRepository) Create(category *Category) error {
	result, err := r.db.Exec(
		"INSERT INTO categories (name, slug, sort) VALUES (?, ?, ?)",
		category.Name, category.Slug, category.Sort,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	category.ID = uint(id)
	return nil
}

func (r *CategoryRepository) GetByID(id uint) (*Category, error) {
	return ScanCategory(r.db.QueryRow(
		"SELECT id, created_at, updated_at, name, slug, sort FROM categories WHERE id = ?",
		id,
	))
}

func (r *CategoryRepository) GetBySlug(slug string) (*Category, error) {
	return ScanCategory(r.db.QueryRow(
		"SELECT id, created_at, updated_at, name, slug, sort FROM categories WHERE slug = ?",
		slug,
	))
}

func (r *CategoryRepository) Update(category *Category) error {
	_, err := r.db.Exec(
		"UPDATE categories SET name = ?, slug = ?, sort = ? WHERE id = ?",
		category.Name, category.Slug, category.Sort, category.ID,
	)
	return err
}

func (r *CategoryRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}

func (r *CategoryRepository) List() ([]*Category, error) {
	rows, err := r.db.Query("SELECT id, created_at, updated_at, name, slug, sort FROM categories ORDER BY sort ASC, id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanCategoryRows(rows)
}
