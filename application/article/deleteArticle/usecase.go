package deletearticle

import (
	"github.com/khanzadimahdi/testproject.git/domain/article"
)

type UseCase struct {
	articlesRepository article.Repository
}

func NewUseCase(articlesRepository article.Repository) *UseCase {
	return &UseCase{
		articlesRepository: articlesRepository,
	}
}

func (uc *UseCase) DeleteArticle(request Request) error {
	return uc.articlesRepository.Delete(request.ArticleUUID)
}
