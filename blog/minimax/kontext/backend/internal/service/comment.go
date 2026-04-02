package service

import (
	"blog/internal/model/dto"
	"blog/internal/repository"
	"errors"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

func NewCommentService() *CommentService {
	return &CommentService{
		commentRepo: repository.NewCommentRepository(),
	}
}

func (s *CommentService) Create(req *dto.CreateCommentDTO, ip string) (*repository.Comment, error) {
	comment := &repository.Comment{
		ArticleID: req.ArticleID,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Content:   req.Content,
		IP:        ip,
		Status:    "pending",
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, errors.New("创建评论失败")
	}

	return comment, nil
}

func (s *CommentService) GetByID(id uint) (*repository.Comment, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("评论不存在")
	}
	return comment, nil
}

func (s *CommentService) UpdateStatus(id uint, status string) (*repository.Comment, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("评论不存在")
	}

	comment.Status = status
	if err := s.commentRepo.Update(comment); err != nil {
		return nil, errors.New("更新评论状态失败")
	}

	return comment, nil
}

func (s *CommentService) Delete(id uint) error {
	return s.commentRepo.Delete(id)
}

func (s *CommentService) ListByArticleID(articleID uint) ([]*repository.Comment, error) {
	return s.commentRepo.ListByArticleID(articleID, "approved")
}

func (s *CommentService) List(page, pageSize int, status string) ([]*repository.Comment, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.commentRepo.List(page, pageSize, status)
}
