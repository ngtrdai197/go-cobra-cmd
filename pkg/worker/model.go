package worker

type EmailDeliveryPayload struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
