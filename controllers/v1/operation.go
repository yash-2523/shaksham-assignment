package v1

import (
	"shaksham/helpers"
	"shaksham/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateOperation(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var reqBody struct {
		URL       string `json:"url"`
		Timestamp int64  `json:"timestamp"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid Request!",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	now := time.Now().Unix()

	if reqBody.URL == "" || !(reqBody.Timestamp >= now && reqBody.Timestamp <= now+2629746) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid Request!",
			Data:    nil,
			Error:   nil,
		})
	}

	var job models.Job

	job.URL = reqBody.URL
	job.Timestamp = reqBody.Timestamp
	job.Status = models.Pending

	err := db.Create(&job).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: "Error in creating job",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	helpers.ScheduleJob(job.Timestamp, job.ID)

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Job created successfully!",
		Data:    job,
		Error:   nil,
	})

}
