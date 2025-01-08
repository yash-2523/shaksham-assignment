package models

import "time"

type statusType string

const (
	Pending   statusType = "pending"
	Completed statusType = "completed"
	Expired   statusType = "expired"
	Failed    statusType = "failed"
)

type Job struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	URL       string     `json:"url"`
	Status    statusType `json:"status" gorm:"column:status;index:idx_timestamp_status"`
	Timestamp int64      `json:"timestamp" gorm:"index:idx_timestamp_status"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
