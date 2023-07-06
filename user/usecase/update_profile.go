package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bxcodec/go-clean-arch/user/delivery/http_request"
)

func (us *userUsecase) UpdateProfile(c context.Context, userId int64, req *http_request.BodyUpdateProfile) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	file := req.File
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

	path := us.cfg.PATH_UPLOAD + filename
	pathMeta := us.cfg.PATH_UPLOAD_META + filename
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	req.Path = pathMeta
	errUpdateProfile := us.userRepo.UpdateProfile(ctx, userId, req)
	if errUpdateProfile != nil {
		return errUpdateProfile
	}

	errUpdateUser := us.userRepo.UpdateUser(ctx, userId, req)
	if errUpdateUser != nil {
		return errUpdateUser
	}

	fmt.Println(errUpdateUser)

	return nil
}
