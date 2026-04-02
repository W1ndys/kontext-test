package service

import (
	"blog/internal/repository"
	"blog/internal/utils"
	"errors"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepo: repository.NewCategoryRepository(),
	}
}

func (s *CategoryService) Create(name, slug string, sort int) (*repository.Category, error) {
	if name == "" {
		return nil, errors.New("分类名称不能为空")
	}
	if slug == "" {
		slug = utils.GenerateSlug(name)
	}

	category := &repository.Category{
		Name: name,
		Slug: slug,
		Sort: sort,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, errors.New("创建分类失败")
	}

	return category, nil
}

func (s *CategoryService) GetByID(id uint) (*repository.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("分类不存在")
	}
	return category, nil
}

func (s *CategoryService) GetBySlug(slug string) (*repository.Category, error) {
	category, err := s.categoryRepo.GetBySlug(slug)
	if err != nil {
		return nil, errors.New("分类不存在")
	}
	return category, nil
}

func (s *CategoryService) Update(id uint, name, slug string, sort int) (*repository.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("分类不存在")
	}

	if name != "" {
		category.Name = name
	}
	if slug != "" {
		category.Slug = slug
	}
	category.Sort = sort

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, errors.New("更新分类失败")
	}

	return category, nil
}

func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}

func (s *CategoryService) List() ([]*repository.Category, error) {
	return s.categoryRepo.List()
}
