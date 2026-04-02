package service

import (
	"blog/internal/model/dto"
	"blog/internal/repository"
	"blog/internal/utils"
	"errors"
)

type ArticleService struct {
	articleRepo *repository.ArticleRepository
	categoryRepo *repository.CategoryRepository
	tagRepo     *repository.TagRepository
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRepo:  repository.NewArticleRepository(),
		categoryRepo: repository.NewCategoryRepository(),
		tagRepo:      repository.NewTagRepository(),
	}
}

func (s *ArticleService) List(page, pageSize int, categoryID, tagID uint, status string) ([]*repository.Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.articleRepo.List(page, pageSize, categoryID, tagID, status)
}

func (s *ArticleService) GetByID(id uint) (*repository.Article, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("文章不存在")
	}
	s.articleRepo.IncrementViewCount(id)
	return article, nil
}

func (s *ArticleService) GetBySlug(slug string) (*repository.Article, error) {
	article, err := s.articleRepo.GetBySlug(slug)
	if err != nil {
		return nil, errors.New("文章不存在")
	}
	s.articleRepo.IncrementViewCount(article.ID)
	return article, nil
}

func (s *ArticleService) Create(req *dto.CreateArticleDTO) (*repository.Article, error) {
	if req.Slug == "" {
		req.Slug = utils.GenerateSlug(req.Title)
	}

	if req.Status == "" {
		req.Status = "draft"
	}

	article := &repository.Article{
		Title:      req.Title,
		Slug:       req.Slug,
		Content:    req.Content,
		Summary:    req.Summary,
		CoverImage: req.CoverImage,
		Status:     req.Status,
		CategoryID: req.CategoryID,
	}

	if err := s.articleRepo.Create(article); err != nil {
		return nil, errors.New("创建文章失败")
	}

	if len(req.TagIDs) > 0 {
		if err := s.articleRepo.UpdateTags(article.ID, req.TagIDs); err != nil {
			return nil, errors.New("更新标签失败")
		}
	}

	return s.articleRepo.GetByID(article.ID)
}

func (s *ArticleService) Update(id uint, req *dto.UpdateArticleDTO) (*repository.Article, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("文章不存在")
	}

	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Slug != "" {
		article.Slug = req.Slug
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.CoverImage != "" {
		article.CoverImage = req.CoverImage
	}
	if req.Status != "" {
		article.Status = req.Status
	}
	if req.CategoryID > 0 {
		article.CategoryID = req.CategoryID
	}

	if err := s.articleRepo.Update(article); err != nil {
		return nil, errors.New("更新文章失败")
	}

	if req.TagIDs != nil {
		if err := s.articleRepo.UpdateTags(article.ID, req.TagIDs); err != nil {
			return nil, errors.New("更新标签失败")
		}
	}

	return s.articleRepo.GetByID(article.ID)
}

func (s *ArticleService) Delete(id uint) error {
	return s.articleRepo.Delete(id)
}

func (s *ArticleService) ListAll() ([]*repository.Article, error) {
	return s.articleRepo.ListAll()
}
