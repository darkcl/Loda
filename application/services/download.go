package services

import (
	"github.com/darkcl/loda/application/models"
	"github.com/darkcl/loda/application/repositories"
	"github.com/darkcl/loda/lib/downloader"
)

// DownloadProgressService is a service to update progress
type DownloadProgressService interface {

	// UpdateProgress takes a downloader progress and transform in to database record
	UpdateProgress(task *models.DownloadTask, progress downloader.DownloadProgress) error

	// MarkDone mark a task as done
	MarkDone(task *models.DownloadTask) error
}

type downloadProgressService struct {
	repo repositories.DownloadRepository
}

// NewDownloadProgessService create download progress service
func NewDownloadProgessService(repo repositories.DownloadRepository) DownloadProgressService {
	return downloadProgressService{
		repo: repo,
	}
}

func (d downloadProgressService) UpdateProgress(task *models.DownloadTask, progress downloader.DownloadProgress) error {
	task.Progress = models.DownloadProgress{
		Label:          progress.Label,
		ETA:            progress.ETA,
		StartAt:        progress.StartAt,
		EndAt:          progress.EndAt,
		BytesComplete:  progress.BytesComplete,
		BytesPerSecond: progress.BytesPerSecond,
		Progress:       progress.Progress,
	}
	err := d.repo.Update(task)
	return err
}

func (d downloadProgressService) MarkDone(task *models.DownloadTask) error {
	task.IsDone = true
	err := d.repo.Update(task)
	return err
}
