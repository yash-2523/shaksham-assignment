package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shaksham/config"
	"shaksham/models"
	"time"
)

func makeRequest(id int) (bool, error) {
	log.Print("Request called", id)
	db := config.Db

	var job models.Job

	err := db.First(&job, id).Error
	if err != nil {
		log.Print(err)
		return false, err
	}

	if job.Status != models.Pending {
		log.Print(err)
		return false, fmt.Errorf("request already")
	}

	samplePayload := map[string]interface{}{
		"message": "hello wordl",
		"id":      id,
	}

	jsonPayload, err := json.Marshal(samplePayload)
	if err != nil {
		log.Print(err)
		return false, fmt.Errorf("unable to marshal payload")
	}

	req, err := http.NewRequest("GET", job.URL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Print(err)
		return false, fmt.Errorf("%v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return false, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		job.Status = models.Completed

		err = db.Where("id = ?", job.ID).Updates(models.Job{Status: models.Completed}).Error
		if err != nil {
			log.Print(err)
			return false, fmt.Errorf("error saving in db %v", err)
		}
	} else {
		log.Print("request failed")
		job.Status = models.Failed

		err = db.Where("id = ?", job.ID).Updates(models.Job{Status: models.Failed}).Error
		if err != nil {
			log.Print(err)
			return false, fmt.Errorf("error saving in db %v", err)
		}
		return false, fmt.Errorf("request failed %v", err)
	}

	return true, nil

}

func ScheduleJob(timestamp int64, id int) {
	reqTime := time.Unix(timestamp, 0)

	delay := time.Until(reqTime)

	if delay > 0 {
		go func() {
			<-time.After(delay)

			makeRequest(id)
		}()
	}
}
