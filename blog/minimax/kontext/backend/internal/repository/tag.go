package repository

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}

func ScanTag(row *sql.Row) (*Tag, error) {
	t := &Tag{}
	err := row.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Name, &t.Slug)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func ScanTagRows(rows *sql.Rows) ([]*Tag, error) {
	var tags []*Tag
	for rows.Next() {
		t := &Tag{}
		err := rows.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Name, &t.Slug)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{db: GetDB()}
}

func (r *TagRepository) Create(tag *Tag) error {
	result, err := r.db.Exec(
		"INSERT INTO tags (name, slug) VALUES (?, ?)",
		tag.Name, tag.Slug,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	tag.ID = uint(id)
	return nil
}

func (r *TagRepository) GetByID(id uint) (*Tag, error) {
	return ScanTag(r.db.QueryRow(
		"SELECT id, created_at, updated_at, name, slug FROM tags WHERE id = ?",
		id,
	))
}

func (r *TagRepository) GetBySlug(slug string) (*Tag, error) {
	return ScanTag(r.db.QueryRow(
		"SELECT id, created_at, updated_at, name, slug FROM tags WHERE slug = ?",
		slug,
	))
}

func (r *TagRepository) Update(tag *Tag) error {
	_, err := r.db.Exec(
		"UPDATE tags SET name = ?, slug = ? WHERE id = ?",
		tag.Name, tag.Slug, tag.ID,
	)
	return err
}

func (r *TagRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM tags WHERE id = ?", id)
	return err
}

func (r *TagRepository) List() ([]*Tag, error) {
	rows, err := r.db.Query("SELECT id, created_at, updated_at, name, slug FROM tags ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanTagRows(rows)
}
