package bizmodel

type Todo struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Description       *string `json:"description"`
	Image             *string `json:"image"`
	Status            string  `json:"status"`
	Date              string  `json:"date"`
	CreatedAtDatetime string  `json:"createdAtDatetime"`
}
