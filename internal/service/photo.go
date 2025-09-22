package service

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"responsible_employee/internal/repository"
	"strings"
	"time"
)

type PhotoService struct {
	repoTask   repository.Task
	repoReport repository.Report
}

func NewPhotoService(repoTask repository.Task, repoReport repository.Report) *PhotoService {
	return &PhotoService{repoTask: repoTask, repoReport: repoReport}
}

func (s *PhotoService) SaveTaskPhoto(taskID string, photo *multipart.FileHeader) (string, error) {
	path, err := s.savePhoto(photo)
	if err != nil {
		return "", err
	}

	err = s.repoTask.AddPhotoToTask(taskID, path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (s *PhotoService) SaveReportPhoto(reportID string, photo *multipart.FileHeader) (string, error) {
	path, err := s.savePhoto(photo)
	if err != nil {
		return "", err
	}

	err = s.repoReport.AddPhotoToReport(reportID, path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (s *PhotoService) savePhoto(file *multipart.FileHeader) (string, error) {
	// Валидация расширения
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	if !allowed[strings.ToLower(ext)] {
		return "", errors.New("unsupported file type")
	}
	timestamp := time.Now().Unix()
	randomSuffix := rand.Intn(1000)
	newFileName := fmt.Sprintf("photo_%d_%03d%s", timestamp, randomSuffix, ext)
	filePath := filepath.Join("uploads", newFileName)

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", fmt.Errorf("unable to create upload directory: %v", err)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to create file: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("unable to save file: %v", err)
	}

	return filePath, nil
}
