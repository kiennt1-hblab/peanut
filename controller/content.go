package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/filemanager"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"
)

type ContentController struct {
	Content usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	return &ContentController{
		Content: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

// ListContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/contents [get]
func (c *ContentController) ListContent(ctx *gin.Context) {
	contents, err := c.Content.GetContents(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    contents,
	})
}

// CreateContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Param			Name		formData	string	true	"string"	minlength(1)	maxlength(30)
//	@Param			Thumbnail	formData	file	true	"file"
//	@Param			Description	formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Param			PlayTime	formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Param			Resolution	formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Param			AspectRatio	formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Param			Tag			formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Param			Category	formData	string	false	"string"	minlength(0)	maxlength(500)
//	@Success		200			{object}	domain.Response
//	@Failure		400			{object}	domain.ErrorResponse
//	@Failure		500			{object}	domain.ErrorResponse
//	@Router			/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	var content domain.CreateContent
	if !bind(ctx, &content) {
		return
	}
	if int(content.Thumbnail.Size) > config.MaxSizeUpload {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file size is too big!",
		})
		return
	}

	extension := filepath.Ext(content.Thumbnail.Filename)
	extensions := []string{".jpeg", ".png", ".jpg"}
	if !filemanager.CheckExtensionAvailable(extension, extensions) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file type not allow",
		})
		return
	}

	err, path := SaveUploadedFileTo(ctx, content.Thumbnail, config.TmpPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.Content.CreateContent(ctx, content, path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
	})
}

// GgStorage godoc
//
//	@Summary		upload file to google storage
//	@Description	google storage
//	@Tags			Google-storage
//	@Accept			json
//	@Produce		json
//	@Param			File	formData	file	true	"file"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/gg-storage [post]
func (c *ContentController) GgStorage(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	if int(file.Size) > config.MaxSizeUpload {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file size is too big!",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	extensions := []string{".jpeg", ".png", ".jpg"}
	if !filemanager.CheckExtensionAvailable(extension, extensions) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file type not allow",
		})
		return
	}

	path, err := c.Content.GgStorage(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response.OK(ctx, path)
}

// Download godoc
//
//	@Summary		download file from Google storage
//	@Description	download file in google storage
//	@Tags			Google-storage
//	@Accept			json
//	@Produce		json
//	@Param			File	formData	file	true	"file"
//	@Success		200		{object}	string
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/gg-storage/{name} [get]
func (c *ContentController) Download(ctx *gin.Context) {
	fileName := ctx.Param("name")
	url := config.PublicUrlGgStorage + "/" + config.BucketUpload + "/" + fileName

	res, err := http.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	reader := res.Body
	defer reader.Close()
	contentLength := res.ContentLength
	contentType := res.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + fileName + `.png"`,
	}

	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

// DeleteGCS godoc
//
//	@Summary		delete file in google storage
//	@Description	delete file in google storage
//	@Tags			Google-storage
//	@Accept			json
//	@Produce		json
//	@Param			File	formData	file	true	"file"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/gg-storage/{name} [delete]
func (c *ContentController) DeleteGCS(ctx *gin.Context) {
	name := ctx.Param("name")
	err := c.Content.DeleteGCS(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
	})
}
