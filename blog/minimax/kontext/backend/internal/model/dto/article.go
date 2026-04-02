package dto

type CreateArticleDTO struct {
	Title      string   `json:"title" binding:"required,min=1,max=200"`
	Slug       string   `json:"slug"`
	Content    string   `json:"content" binding:"required"`
	Summary    string   `json:"summary"`
	CoverImage string   `json:"cover_image"`
	Status     string   `json:"status" binding:"omitempty,oneof=draft published"`
	CategoryID uint     `json:"category_id"`
	TagIDs     []uint   `json:"tag_ids"`
}

type UpdateArticleDTO struct {
	Title      string   `json:"title" binding:"omitempty,min=1,max=200"`
	Slug       string   `json:"slug"`
	Content    string   `json:"content"`
	Summary    string   `json:"summary"`
	CoverImage string   `json:"cover_image"`
	Status     string   `json:"status" binding:"omitempty,oneof=draft published"`
	CategoryID uint     `json:"category_id"`
	TagIDs     []uint   `json:"tag_ids"`
}

type ArticleListQuery struct {
	Page       int    `form:"page" binding:"omitempty,min=1"`
	PageSize   int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	CategoryID uint   `form:"category_id"`
	TagID      uint   `form:"tag_id"`
	Status     string `form:"status"`
}
