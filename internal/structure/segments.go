package structure

type Segments struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	AppId              string  `json:"app_id"`
	ReadOnly           bool    `json:"read_only"`
	IsActive           bool    `json:"is_active"`
	Source             string  `json:"source"`
	SegmentStatus      *string `json:"segment_status"`
	LoadingStartedAt   *string `json:"loading_started_at"`
	LoadingCompletedAt *string `json:"loading_completed_at"`
}
