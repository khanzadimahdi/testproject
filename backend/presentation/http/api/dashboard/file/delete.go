package file

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/khanzadimahdi/testproject/application/auth"
	deletefile "github.com/khanzadimahdi/testproject/application/dashboard/file/deleteFile"
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/permission"
)

type deleteHandler struct {
	deleteFileUseCase *deletefile.UseCase
	authorizer        domain.Authorizer
}

func NewDeleteHandler(deleteFileUseCase *deletefile.UseCase, a domain.Authorizer) *deleteHandler {
	return &deleteHandler{
		deleteFileUseCase: deleteFileUseCase,
		authorizer:        a,
	}
}

func (h *deleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	userUUID := auth.FromContext(r.Context()).UUID
	if ok, err := h.authorizer.Authorize(userUUID, permission.FilesCreate); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	UUID := httprouter.ParamsFromContext(r.Context()).ByName("uuid")

	err := h.deleteFileUseCase.DeleteFile(deletefile.Request{
		FileUUID: UUID,
	})

	switch true {
	case errors.Is(err, domain.ErrNotExists):
		rw.WriteHeader(http.StatusNotFound)
	case err != nil:
		rw.WriteHeader(http.StatusInternalServerError)
	default:
		rw.WriteHeader(http.StatusOK)
	}
}
