package dto

type UserItem struct {
	Username string `json:"username"`
}

type ItemCreateRequest struct {
	ItemName string `json:"item_name" binding:"required"`
	Description string `json:"description"`
}

type ItemListResponse struct {
	User UserItem `json:"owner"`
	ItemName string `json:"item_name"`
}

type ItemDetailResponse struct {
	User UserItem `json:"owner"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
}

type ItemUpdateRequest struct {
	ItemName string `json:"item_name"`
	Description string `json:"description"`
}

type ItemDeleteRequest struct {
	ID uint `json:"id"`
}
