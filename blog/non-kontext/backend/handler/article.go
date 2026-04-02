package handler

import (
	"blog-backend/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type paginatedResponse struct {
	Data       []model.Article `json:"data"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}

// GetArticles handles GET /api/articles?page=1&size=10&category=&tag=
func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if size < 1 || size > 100 {
		size = 10
	}

	category := r.URL.Query().Get("category")
	tag := r.URL.Query().Get("tag")

	query := model.DB.Model(&model.Article{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	var total int64
	query.Count(&total)

	var articles []model.Article
	offset := (page - 1) * size
	query.Order("created_at DESC").Offset(offset).Limit(size).Find(&articles)

	totalPages := int(total) / size
	if int(total)%size != 0 {
		totalPages++
	}

	json.NewEncoder(w).Encode(paginatedResponse{
		Data:       articles,
		Total:      total,
		Page:       page,
		PageSize:   size,
		TotalPages: totalPages,
	})
}

// GetArticle handles GET /api/articles/{id}
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "article not found"})
		return
	}

	json.NewEncoder(w).Encode(article)
}

// GetCategories handles GET /api/categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories []string
	model.DB.Model(&model.Article{}).Distinct("category").Where("category != ''").Pluck("category", &categories)

	json.NewEncoder(w).Encode(categories)
}

// GetTags handles GET /api/tags
func GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tagStrings []string
	model.DB.Model(&model.Article{}).Where("tags != ''").Pluck("tags", &tagStrings)

	tagSet := make(map[string]bool)
	for _, ts := range tagStrings {
		for _, t := range strings.Split(ts, ",") {
			t = strings.TrimSpace(t)
			if t != "" {
				tagSet[t] = true
			}
		}
	}

	tags := make([]string, 0, len(tagSet))
	for t := range tagSet {
		tags = append(tags, t)
	}

	json.NewEncoder(w).Encode(tags)
}

// CreateArticle handles POST /api/admin/articles
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if err := model.DB.Create(&article).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to create article"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

// UpdateArticle handles PUT /api/admin/articles/{id}
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "article not found"})
		return
	}

	var input model.Article
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	model.DB.Model(&article).Updates(model.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Tags:     input.Tags,
		Summary:  input.Summary,
		Cover:    input.Cover,
		Status:   input.Status,
	})

	json.NewEncoder(w).Encode(article)
}

// DeleteArticle handles DELETE /api/admin/articles/{id}
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if err := model.DB.Delete(&model.Article{}, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "article not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "article deleted"})
}
