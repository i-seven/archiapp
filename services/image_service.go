package services

import (
	"backendAf/config"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var ImageService = imageService{}

type imageService struct{}

// UploadProductImage saves a new image and returns metadata (filename + url)
func (s imageService) UploadProductImage(productID string, file *multipart.FileHeader) (map[string]string, error) {
	dir := config.ImageDir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// determine next sequence for this product
	pattern := filepath.Join(dir, fmt.Sprintf("%s_*", productID))
	files, _ := filepath.Glob(pattern)
	imageID := len(files) + 1

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", productID, imageID, ext)
	savePath := filepath.Join(dir, filename)

	if err := saveUploadedFile(file, savePath); err != nil {
		return nil, err
	}

	return map[string]string{
		"filename": filename,
		"url":      "/images/" + filename,
	}, nil
}

// UpdateProductImage replaces any existing file for productID_imageID and writes the new one.
// It removes previously existing files with any extension for that image slot.
func (s imageService) UpdateProductImage(productID, imageID string, file *multipart.FileHeader) (map[string]string, error) {
	dir := config.ImageDir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// remove old files like productID_imageID.*
	oldPattern := filepath.Join(dir, fmt.Sprintf("%s_%s.*", productID, imageID))
	oldFiles, _ := filepath.Glob(oldPattern)
	for _, of := range oldFiles {
		_ = os.Remove(of)
	}

	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s_%s%s", productID, imageID, ext)
	savePath := filepath.Join(dir, newFilename)

	if err := saveUploadedFile(file, savePath); err != nil {
		return nil, err
	}

	return map[string]string{
		"filename": newFilename,
		"url":      "/images/" + newFilename,
	}, nil
}

// ListProductImages returns public URLs for all images for a product.
func (s imageService) ListProductImages(productID string) ([]string, error) {
	dir := config.ImageDir
	pattern := filepath.Join(dir, fmt.Sprintf("%s_*", productID))
	files, _ := filepath.Glob(pattern)

	var urls []string
	for _, f := range files {
		urls = append(urls, "/images/"+filepath.Base(f))
	}
	return urls, nil
}

// DeleteProductImage deletes the image file(s) for a product/imageId slot (any extension).
func (s imageService) DeleteProductImage(productID, imageID string) error {
	dir := config.ImageDir
	pattern := filepath.Join(dir, fmt.Sprintf("%s_%s.*", productID, imageID))
	files, _ := filepath.Glob(pattern)
	if len(files) == 0 {
		return os.ErrNotExist
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

// saveUploadedFile streams the uploaded file to dst path (no large memory buffers)
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// copy stream (efficient)
	if _, err := io.Copy(out, src); err != nil {
		return err
	}
	return nil
}
