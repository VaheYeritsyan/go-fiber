package models

import "time"

type Order struct {
 ID           uint `json:"id" gorm:"primaryKey"`
 CreatedAt    time.Time
 ProductRefer int     `json:"product_id" gorm:"foreignKey:ProductRefer"`
 Product      Product `gorm:"foreignKey:ProductRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
 UserRefer    int     `json:"user_id" gorm:"foreignKey:UserRefer"`
 User         User    `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
