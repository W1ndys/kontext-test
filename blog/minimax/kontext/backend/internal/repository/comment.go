package repository

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ArticleID uint      `json:"article_id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	IP        string    `json:"ip"`
	Status    string    `json:"status"`
}

func ScanComment(row *sql.Row) (*Comment, error) {
	c := &Comment{}
	err := row.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.ArticleID, &c.Nickname, &c.Email, &c.Content, &c.IP, &c.Status)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ScanCommentRows(rows *sql.Rows) ([]*Comment, error) {
	var comments []*Comment
	for rows.Next() {
		c := &Comment{}
		err := rows.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.ArticleID, &c.Nickname, &c.Email, &c.Content, &c.IP, &c.Status)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{db: GetDB()}
}

func (r *CommentRepository) Create(comment *Comment) error {
	result, err := r.db.Exec(
		"INSERT INTO comments (article_id, nickname, email, content, ip, status) VALUES (?, ?, ?, ?, ?, ?)",
		comment.ArticleID, comment.Nickname, comment.Email, comment.Content, comment.IP, comment.Status,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	comment.ID = uint(id)
	return nil
}

func (r *CommentRepository) GetByID(id uint) (*Comment, error) {
	return ScanComment(r.db.QueryRow(
		"SELECT id, created_at, updated_at, article_id, nickname, email, content, ip, status FROM comments WHERE id = ?",
		id,
	))
}

func (r *CommentRepository) Update(comment *Comment) error {
	_, err := r.db.Exec(
		"UPDATE comments SET nickname = ?, email = ?, content = ?, ip = ?, status = ? WHERE id = ?",
		comment.Nickname, comment.Email, comment.Content, comment.IP, comment.Status, comment.ID,
	)
	return err
}

func (r *CommentRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM comments WHERE id = ?", id)
	return err
}

func (r *CommentRepository) ListByArticleID(articleID uint, status string) ([]*Comment, error) {
	query := "SELECT id, created_at, updated_at, article_id, nickname, email, content, ip, status FROM comments WHERE article_id = ?"
	args := []interface{}{articleID}

	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	} else {
		query += " AND status = 'approved'"
	}

	rows, err := r.db.Query(query+" ORDER BY created_at DESC", args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanCommentRows(rows)
}

func (r *CommentRepository) List(page, pageSize int, status string) ([]*Comment, int64, error) {
	countQuery := "SELECT COUNT(*) FROM comments WHERE 1=1"
	query := "SELECT id, created_at, updated_at, article_id, nickname, email, content, ip, status FROM comments WHERE 1=1"
	args := []interface{}{}

	if status != "" {
		countQuery += " AND status = ?"
		query += " AND status = ?"
		args = append(args, status)
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
	comments, err := ScanCommentRows(rows)
	return comments, total, err
}
