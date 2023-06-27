package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (dus *donorUsecase) UploadBukti(c context.Context, id int64, file *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	fileExt := filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

	path := dus.cfg.PATH_UPLOAD + filename
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	errR := dus.donorRepo.UploadBukti(ctx, id, path)
	if errR != nil {
		return errR
	}

	return nil
}
