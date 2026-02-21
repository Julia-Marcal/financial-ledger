package model

import "time"

// AuditEntry stores a timestamp and optional user identifier for an event.
type AuditEntry struct {
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	UserID    string    `json:"userId" bson:"userId"`
}

// Audit groups created and updated entries for a document.
type Audit struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	CreatedBy string    `json:"createdBy" bson:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy string    `json:"updatedBy" bson:"updatedBy"`
}
