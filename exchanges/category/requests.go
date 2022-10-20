package category

type StoreCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type IdRequest struct {
	ID uint `json:"id"`
}
type UpdatedRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          uint   `json:"id"`
}
