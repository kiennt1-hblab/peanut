package usecase

import (
	"context"
	"mime/multipart"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/pkg/filemanager"
	"peanut/repository"
	"time"
)

type ContentUsecase interface {
	GetContents(ctx context.Context) ([]domain.Content, error)
	CreateContent(ctx context.Context, content domain.CreateContent, filepath string) error
	GgStorage(ctx context.Context, file *multipart.FileHeader) (string, error)
	DeleteGCS(ctx context.Context, name string) error
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func (c contentUsecase) GetContents(ctx context.Context) (contents []domain.Content, err error) {
	contents, err = c.ContentRepo.GetContents()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (c contentUsecase) CreateContent(ctx context.Context, content domain.CreateContent, filePath string) (err error) {
	data := domain.Content{
		Name:        content.Name,
		Thumbnail:   filePath,
		Description: content.Description,
		Tag:         content.Tag,
		Category:    content.Category,
		Resolution:  content.Resolution,
		AspectRatio: content.AspectRatio,
	}
	data.PlayTime, err = time.Parse(config.TimeFormatDefault, content.PlayTime)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	err = c.ContentRepo.CreateContent(data)
	if err != nil {
		return
	}
	return nil
}

func (c contentUsecase) GgStorage(ctx context.Context, file *multipart.FileHeader) (string, error) {
	// open file
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	path, err := filemanager.UploadGgStorage(fileContent, "")
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c contentUsecase) DeleteGCS(ctx context.Context, name string) error {
	err := filemanager.DeleteGCS(name)
	if err != nil {
		return err
	}
	return nil
}

func NewContentUsecase(repo repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: repo,
	}
}
