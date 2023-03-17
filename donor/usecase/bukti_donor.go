package usecase

import (
	"context"
	"io"
	"mime/multipart"
	"os"
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
	// newFilename := "new_filename.jpg"
	path := "donor/upload/" + file.Filename
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
