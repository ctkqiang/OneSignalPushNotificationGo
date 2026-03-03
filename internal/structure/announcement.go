package structure

import "time"

type AnnouncementType string

var (
	ANNOUNCEMENT_TYPE_HOLIDAY = AnnouncementType("HOLIDAY")
	ANNOUNCEMENT_TYPE_EVENT   = AnnouncementType("EVENT")
)

type Announcement struct {
	ID        string           `bson:"_id,omitempty" json:"id"`
	Type      AnnouncementType `bson:"type" json:"type"`
	Message   string           `bson:"message" json:"message"`
	Priority  Priority         `bson:"priority" json:"priority"`
	CreatedAt time.Time        `bson:"created_at" json:"created_at"`
	StartedAt  time.Time        `bson:"started_at" json:"started_at"`
	ExpiresAt time.Time        `bson:"expires_at" json:"expires_at"`
}