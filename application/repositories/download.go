package repositories

import (
	storm "github.com/asdine/storm/v3"
	"github.com/darkcl/loda/application/models"
)

// DownloadRepository describe all database operation
type DownloadRepository interface {
	Create(destination string, taskType string) (models.DownloadTask, error)
	Update(task models.DownloadTask) error
	FindOne(id int) (models.DownloadTask, error)
	Destroy(id int) error
}

type downloadRepository struct {
	DownloadRepository
	db *storm.DB
}

// NewDownloadRepository create download repository with storm database
func NewDownloadRepository(db *storm.DB) DownloadRepository {
	return &downloadRepository{
		db: db,
	}
}

func (d downloadRepository) Create(destination string, taskType string) (models.DownloadTask, error) {
	task := models.DownloadTask{
		Destination: destination,
		TaskType:    taskType,
		IsDone:      false,
	}

	err := d.db.Save(task)
	return task, err
}

func (d downloadRepository) Update(task models.DownloadTask) error {
	err := d.db.Save(task)
	return err
}

func (d downloadRepository) FindOne(id int) (models.DownloadTask, error) {
	var task models.DownloadTask
	err := d.db.One("ID", id, &task)
	return task, err
}

func (d downloadRepository) Destroy(id int) error {
	var task models.DownloadTask
	err := d.db.One("ID", id, &task)
	err = d.db.DeleteStruct(&task)
	return err
}
