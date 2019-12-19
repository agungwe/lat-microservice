package model

type ProductData struct {
	Data Product `json:"data"`
}

type Product struct {
	ProductName string `json:"product_name"`
	SKU         string `json:"sku"`
	Qty         int32  `json:"qty"`
}
