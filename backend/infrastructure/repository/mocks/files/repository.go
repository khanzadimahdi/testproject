package files

import (
	"github.com/khanzadimahdi/testproject/domain/file"
	"github.com/stretchr/testify/mock"
)

type MockFilesRepository struct {
	mock.Mock
}

var _ file.Repository = &MockFilesRepository{}

func (r *MockFilesRepository) GetAll(offset uint, limit uint) ([]file.File, error) {
	args := r.Called(offset, limit)

	if f, ok := args.Get(0).([]file.File); ok {
		return f, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *MockFilesRepository) GetOne(UUID string) (file.File, error) {
	args := r.Called(UUID)

	return args.Get(0).(file.File), args.Error(1)
}

func (r *MockFilesRepository) Save(f *file.File) (string, error) {
	args := r.Called(f)

	return args.String(0), args.Error(1)
}

func (r *MockFilesRepository) Delete(UUID string) error {
	args := r.Called(UUID)

	return args.Error(0)
}

func (r *MockFilesRepository) Count() (uint, error) {
	args := r.Called()

	return args.Get(0).(uint), args.Error(1)
}