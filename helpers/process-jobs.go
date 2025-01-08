package helpers

import (
	"fmt"
	"log"
	"shaksham/config"
	"shaksham/models"
	"time"
)

func ProcessOldJobs() (bool, error) {
	db := config.Db
	currentTime := time.Now().Unix()

	err := db.Where("timestamp < ? AND status = ?", currentTime, models.Pending).Updates(models.Job{Status: models.Expired}).Error

	if err != nil {
		log.Print(err)
		return false, fmt.Errorf("error updating expired jobs %v", err)
	}

	currentTime = time.Now().Unix()
	var jobs []models.Job
	err = db.Where("timestamp > ? AND status = ?", currentTime, models.Pending).Find(&jobs).Error
	if err != nil {
		log.Print(err)
		return false, fmt.Errorf("error getting jobs %v", err)
	}

	for _, job := range jobs {
		ScheduleJob(job.Timestamp, job.ID)
	}

	return true, nil
}
