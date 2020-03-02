package service

type catalogItem struct {
	ListingID       int    `json:"listing_id"`
	ProductID       string `json:"product_id"`
	Description     string `json:"description"`
	Price           uint32 `json:"price"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}

type fulfillmentStatus struct {
	ProductID       string `json:"product_id"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}
