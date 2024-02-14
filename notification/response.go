package notification

type NotificationResponse struct {
	Name         string              `json:"name"`
	Symbol       string              `json:"symbol"`
	CurrentPrice float64             `json:"price"`
	TargetPrice  float64             `json:"target_price"`
	Time         string              `json:"time"`
	Methods      []string            `json:"methods"`
	GreaterThan  bool                `json:"greater_than"`
	MethodValue  MethodValueResponse `json:"method_value"`
	Notified     bool                `json:"notified"`
	NotifiedAt   string              `json:"notified_at"`
}

type MethodValueResponse struct {
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
	Twitter  string `json:"twitter"`
}
