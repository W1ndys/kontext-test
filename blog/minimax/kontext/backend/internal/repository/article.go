package repository

import (
	"database/sql"
	"time"
)

type Article struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Summary    string    `json:"summary"`
	CoverImage string    `json:"cover_image"`
	Status     string    `json:"status"`
	ViewCount  int       `json:"view_count"`
	CategoryID uint      `json:"category_id"`
	Category   *Category `json:"category,omitempty"`
	Tags       []*Tag   `json:"tags,omitempty"`
}

func ScanArticle(row *sql.Row) (*Article, error) {
	a := &Article{}
	err := row.Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt, &a.Title, &a.Slug, &a.Content, &a.Summary, &a.CoverImage, &a.Status, &a.ViewCount, &a.CategoryID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func ScanArticleRows(rows *sql.Rows) ([]*Article, error) {
	articles := []*Article{}
	for rows.Next() {
		a := &Article{}
		err := rows.Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt, &a.Title, &a.Slug, &a.Content, &a.Summary, &a.CoverImage, &a.Status, &a.ViewCount, &a.CategoryID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{db: GetDB()}
}

func (r *ArticleRepository) Create(article *Article) error {
	result, err := r.db.Exec(
		"INSERT INTO articles (title, slug, content, summary, cover_image, status, view_count, category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		article.Title, article.Slug, article.Content, article.Summary, article.CoverImage, article.Status, article.ViewCount, article.CategoryID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	article.ID = uint(id)
	return nil
}

func (r *ArticleRepository) GetByID(id uint) (*Article, error) {
	return ScanArticle(r.db.QueryRow(
		"SELECT id, created_at, updated_at, title, slug, content, summary, cover_image, status, view_count, category_id FROM articles WHERE id = ?",
		id,
	))
}

func (r *ArticleRepository) GetBySlug(slug string) (*Article, error) {
	return ScanArticle(r.db.QueryRow(
		"SELECT id, created_at, updated_at, title, slug, content, summary, cover_image, status, view_count, category_id FROM articles WHERE slug = ?",
		slug,
	))
}

func (r *ArticleRepository) Update(article *Article) error {
	_, err := r.db.Exec(
		"UPDATE articles SET title = ?, slug = ?, content = ?, summary = ?, cover_image = ?, status = ?, view_count = ?, category_id = ? WHERE id = ?",
		article.Title, article.Slug, article.Content, article.Summary, article.CoverImage, article.Status, article.ViewCount, article.CategoryID, article.ID,
	)
	return err
}

func (r *ArticleRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

func (r *ArticleRepository) List(page, pageSize int, categoryID, tagID uint, status string) ([]*Article, int64, error) {
	countQuery := "SELECT COUNT(*) FROM articles WHERE 1=1"
	query := "SELECT id, created_at, updated_at, title, slug, content, summary, cover_image, status, view_count, category_id FROM articles WHERE 1=1"
	args := []interface{}{}

	if status != "" {
		countQuery += " AND status = ?"
		query += " AND status = ?"
		args = append(args, status)
	}
	if categoryID > 0 {
		countQuery += " AND category_id = ?"
		query += " AND category_id = ?"
		args = append(args, categoryID)
	}
	if tagID > 0 {
		countQuery += " AND id IN (SELECT article_id FROM article_tags WHERE tag_id = ?)"
		query += " AND id IN (SELECT article_id FROM article_tags WHERE tag_id = ?)"
		args = append(args, tagID)
	}

	var total int64
	r.db.QueryRow(countQuery, args...).Scan(&total)

	offset := (page - 1) * pageSize
	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	articles, err := ScanArticleRows(rows)
	if err != nil {
		return nil, 0, err
	}

	// Load categories and tags
	for _, a := range articles {
		if a.CategoryID > 0 {
			catRepo := NewCategoryRepository()
			a.Category, _ = catRepo.GetByID(a.CategoryID)
		}
		a.Tags = r.GetArticleTags(a.ID)
	}

	return articles, total, nil
}

func (r *ArticleRepository) ListAll() ([]*Article, error) {
	rows, err := r.db.Query(
		"SELECT id, created_at, updated_at, title, slug, content, summary, cover_image, status, view_count, category_id FROM articles ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanArticleRows(rows)
}

func (r *ArticleRepository) GetArticleTags(articleID uint) []*Tag {
	rows, err := r.db.Query(`
		SELECT t.id, t.created_at, t.updated_at, t.name, t.slug
		FROM tags t
		INNER JOIN article_tags at ON t.id = at.tag_id
		WHERE at.article_id = ?
	`, articleID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var tags []*Tag
	for rows.Next() {
		t := &Tag{}
		rows.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Name, &t.Slug)
		tags = append(tags, t)
	}
	return tags
}

func (r *ArticleRepository) UpdateTags(articleID uint, tagIDs []uint) error {
	_, err := r.db.Exec("DELETE FROM article_tags WHERE article_id = ?", articleID)
	if err != nil {
		return err
	}
	for _, tagID := range tagIDs {
		_, err = r.db.Exec("INSERT INTO article_tags (article_id, tag_id) VALUES (?, ?)", articleID, tagID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ArticleRepository) IncrementViewCount(id uint) error {
	_, err := r.db.Exec("UPDATE articles SET view_count = view_count + 1 WHERE id = ?", id)
	return err
}
