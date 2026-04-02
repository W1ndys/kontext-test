package dto

type CreateCommentDTO struct {
	ArticleID uint   `json:"article_id" binding:"required"`
	Nickname  string `json:"nickname" binding:"required,min=1,max=50"`
	Email     string `json:"email"`
	Content   string `json:"content" binding:"required,min=1"`
}

type UpdateCommentStatusDTO struct {
	Status string `json:"status" binding:"required,oneof=pending approved rejected"`
}
