package dto

import "time"

// ChartReviewEntity represents the entity for a chart review
type ChartReviewEntity struct {
	PatientID string    `bson:"patientId" json:"patientId"`
	Charts    Chart     `bson:"charts" json:"charts"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type Chart struct {
	ChartType string `bson:"chartType" json:"chartType"`
	Top       string `bson:"top" json:"top"`
	Left      string `bson:"left" json:"left"`
}
