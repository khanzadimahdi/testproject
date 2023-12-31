package getarticles

import (
	"github.com/khanzadimahdi/testproject.git/domain/article"
)

const limit = 10

type UseCase struct {
	articlesRepository article.Repository
}

func NewUseCase(articlesRepository article.Repository) *UseCase {
	return &UseCase{
		articlesRepository: articlesRepository,
	}
}

func (uc *UseCase) GetArticles(request *Request) (*GetArticlesResponse, error) {
	totalArticles, err := uc.articlesRepository.Count()
	if err != nil {
		return nil, err
	}

	currentPage := request.Page
	var offset uint = 0
	if currentPage > 0 {
		offset = (currentPage - 1) * limit
	}

	totalPages := totalArticles / limit

	if (totalPages * limit) != totalArticles {
		totalPages++
	}

	a, err := uc.articlesRepository.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	return NewGetArticlesReponse(a, totalPages, currentPage), nil
}
