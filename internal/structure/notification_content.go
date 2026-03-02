package structure

type NotificationContent struct { 
	Title string `json:"title"`
	Message string `json:"message"`
	ImageUrl string `json:"image_url"`
	Channel string `json:"channel"`
	Locale Locale `json:"locale"`
}
