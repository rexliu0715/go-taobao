package models

type Item struct {
	Cover         *string    `json:"cover,omitempty"`
	ItemId        *string    `json:"itemId,omitempty" gorm:"type:varchar(255);uniqueIndex:idx_item_id"`
	Title         *string    `json:"title,omitempty" gorm:"type:varchar(255)"`
	Description   *string    `json:"description,omitempty"`
	Currency      *string    `json:"currency,omitempty" gorm:"type:varchar(100)"`
	Price         *float64   `json:"price,omitempty" gorm:"type:varchar(100)"`
	AuctionImages *[]string  `json:"auctionImages,omitempty" gorm:"type:jsonb;serializer:json"`
	SellerId      *string    `json:"sellerId,omitempty" gorm:"type:varchar(255);index;"`
	SellerNick    *string    `json:"sellerNick,omitempty" gorm:"type:varchar(255)"`
	ShopId        *string    `json:"shopId,omitempty" gorm:"type:varchar(255);index;"`
	ShopName      *string    `json:"shopName,omitempty" gorm:"type:varchar(255)"`
	ShopURL       *string    `json:"shopUrl,omitempty"`
	Variants      *[]Variant `json:"variants,omitempty" gorm:"references:ItemId;constraint:OnUpdate:CASCADE"`
}

func (Item) TableName() string {
	return "taobao_items"
}
