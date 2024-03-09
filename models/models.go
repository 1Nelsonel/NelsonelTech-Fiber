package models

import (
	"fmt"
    "gorm.io/gorm"
	"github.com/gosimple/slug"
)

type ContactForm struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Subject string `json:"subject"`
	Body     string `json:"body"`
}

type Service struct {
	gorm.Model
	Name          string         `json:"name"`
	Slug          string         `json:"slug" gorm:"unique_index"`
	Price         float64        `json:"price"`
	Description   string         `json:"description"`	
	Status        bool           `json:"status" gorm:"default:true"`
	ServiceImage  string         `json:"service_image"`
}



// BeforeCreate hook to generate a unique slug before creating a new service
func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
    generatedSlug := slug.Make(s.Name)
    
    // Check if the generated slug is unique
    var count int64
    if err := tx.Model(&Service{}).Where("Slug = ?", generatedSlug).Count(&count).Error; err != nil {
        return err
    }
    
    // If the slug is not unique, append a counter to make it unique
    if count > 0 {
        generatedSlug = fmt.Sprintf("%s-%d", generatedSlug, count)
    }
    
    s.Slug = generatedSlug
    return nil
}
