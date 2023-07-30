package models

type Variant struct {
	Id       string    `json:"id,omitempty" gorm:"type:varchar(255);uniqueIndex:idx_item_id_variant_id"`
	Image    *string   `json:"image,omitempty"`
	Atr      *[]string `json:"atr,omitempty" gorm:"type:jsonb;serializer:json;"`
	Currency *string   `json:"currency,omitempty" gorm:"type:varchar(100)"`
	Price    *float64  `json:"price,omitempty"`
	ItemId   *string   `json:"itemId,omitempty" gorm:"type:varchar(255);uniqueIndex:idx_item_id_variant_id"`
	Item     *Item     `json:"item,omitempty" gorm:"foreignKey:ItemId;references:ItemId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Variant) TableName() string {
	return "taobao_item_variants"
}
