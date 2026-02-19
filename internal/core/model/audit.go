package model

import "time"

// AuditEntry stores a timestamp and optional user identifier for an event.
type AuditEntry struct {
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	UserID    string    `json:"userId" bson:"userId"`
}

// Audit groups created and updated entries for a document.
type Audit struct {
	Created AuditEntry `json:"created" bson:"created"`
	Updated AuditEntry `json:"updated,omitempty" bson:"updated,omitempty"`
}
