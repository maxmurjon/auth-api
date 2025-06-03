package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadFile(file multipart.File, fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("creating upload directory: %w", err)
	}

	sanitizedFileName := filepath.Base(fileHeader.Filename)
	uniqueFileName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(sanitizedFileName)
	filePath := filepath.Join(uploadDir, uniqueFileName)

	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return "", fmt.Errorf("writing file: %w", err)
	}

	return filePath, nil
}

func ValidateProductID(productIdStr string) (int, error) {
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return 0, fmt.Errorf("invalid product_id: %w", err)
	}
	return productId, nil
}
