package service

import (
	"blog/internal/repository"
	"blog/internal/utils"
	"errors"
)

type TagService struct {
	tagRepo *repository.TagRepository
}

func NewTagService() *TagService {
	return &TagService{
		tagRepo: repository.NewTagRepository(),
	}
}

func (s *TagService) Create(name, slug string) (*repository.Tag, error) {
	if name == "" {
		return nil, errors.New("标签名称不能为空")
	}
	if slug == "" {
		slug = utils.GenerateSlug(name)
	}

	tag := &repository.Tag{
		Name: name,
		Slug: slug,
	}

	if err := s.tagRepo.Create(tag); err != nil {
		return nil, errors.New("创建标签失败")
	}

	return tag, nil
}

func (s *TagService) GetByID(id uint) (*repository.Tag, error) {
	tag, err := s.tagRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("标签不存在")
	}
	return tag, nil
}

func (s *TagService) GetBySlug(slug string) (*repository.Tag, error) {
	tag, err := s.tagRepo.GetBySlug(slug)
	if err != nil {
		return nil, errors.New("标签不存在")
	}
	return tag, nil
}

func (s *TagService) Update(id uint, name, slug string) (*repository.Tag, error) {
	tag, err := s.tagRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("标签不存在")
	}

	if name != "" {
		tag.Name = name
	}
	if slug != "" {
		tag.Slug = slug
	}

	if err := s.tagRepo.Update(tag); err != nil {
		return nil, errors.New("更新标签失败")
	}

	return tag, nil
}

func (s *TagService) Delete(id uint) error {
	return s.tagRepo.Delete(id)
}

func (s *TagService) List() ([]*repository.Tag, error) {
	return s.tagRepo.List()
}
