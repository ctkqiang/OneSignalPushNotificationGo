package database

import (
	"valuefarm_pushnotification_services/internal/structure"

	"github.com/kamva/mgm/v3"
)

type Status string

var (
	StatusSuccess Status = "success"
	StatusPending Status = "pending"
	StatusFailed  Status = "failed"
)

type NotificationResponse struct {
	mgm.DefaultModel `bson:",inline"`

	Status     Status                         `json:"status" bson:"status"`
	Content    *structure.NotificationContent `json:"content" bson:"content"`
	AuditTrail *structure.AuditTrail          `json:"audit_trail" bson:"audit_trail"`
}
