package seeds

import (
	"log"

	"github.com/gift-redemption/internal/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	seedUsers(db)
	seedGifts(db)
	log.Println("seeding completed")
}

func seedUsers(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		return
	}

	users := []model.User{
		{Name: "Admin gift-redemption", Email: "admin@gift-redemption.com", Role: model.RoleAdmin},
		{Name: "John Doe", Email: "john@example.com", Role: model.RoleUser},
	}

	for i := range users {
		if err := users[i].HashPassword("password123"); err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}
	log.Printf("seeded %d users", len(users))
}

func seedGifts(db *gorm.DB) {
	var count int64
	db.Model(&model.Gift{}).Count(&count)
	if count > 0 {
		return
	}

	gifts := []model.Gift{
		{
			Name:         "Samsung Galaxy S9 Midnight Black 4/64 GB",
			Description:  "Ukuran layar 6.2 inci, Dual Edge Super AMOLED 2960x1440, iQuad HD+529 ppi, RAM 6 GB, ROM 64 GB, MicroSD up to 400GB, Android 8.0 Oreo",
			Point:        200000,
			Stock:        10,
			ImageURL:     "https://via.placeholder.com/400x400?text=Galaxy+S9",
			IsNew:        true,
			IsBestSeller: false,
			AvgRating:    4.3,
			TotalReviews: 160,
		},
		{
			Name:         "Apple AirPods Pro 2nd Gen",
			Description:  "Active Noise Cancellation, Adaptive Transparency, Personalized Spatial Audio with dynamic head tracking",
			Point:        350000,
			Stock:        5,
			ImageURL:     "https://via.placeholder.com/400x400?text=AirPods+Pro",
			IsNew:        false,
			IsBestSeller: true,
			AvgRating:    4.7,
			TotalReviews: 320,
		},
		{
			Name:         "Xiaomi Redmi Note 12 Pro",
			Description:  "6.67 inch AMOLED, 200MP Camera, 5000mAh battery, 67W HyperCharge",
			Point:        150000,
			Stock:        15,
			ImageURL:     "https://via.placeholder.com/400x400?text=Redmi+Note+12",
			IsNew:        false,
			IsBestSeller: true,
			AvgRating:    3.8,
			TotalReviews: 95,
		},
		{
			Name:         "JBL Flip 6 Bluetooth Speaker",
			Description:  "Portable waterproof speaker, 12 hours battery life, PartyBoost compatible",
			Point:        80000,
			Stock:        20,
			ImageURL:     "https://via.placeholder.com/400x400?text=JBL+Flip+6",
			IsNew:        false,
			IsBestSeller: false,
			AvgRating:    4.1,
			TotalReviews: 210,
		},
		{
			Name:         "Sony WH-1000XM5 Headphones",
			Description:  "Industry-leading noise canceling, 30-hour battery, multipoint connection",
			Point:        420000,
			Stock:        0, // sold out
			ImageURL:     "https://via.placeholder.com/400x400?text=Sony+WH1000XM5",
			IsNew:        false,
			IsBestSeller: false,
			AvgRating:    4.9,
			TotalReviews: 540,
		},
		{
			Name:         "Logitech MX Master 3S Mouse",
			Description:  "8K DPI sensor, MagSpeed scroll wheel, quiet clicks, USB-C charging",
			Point:        120000,
			Stock:        8,
			ImageURL:     "https://via.placeholder.com/400x400?text=MX+Master+3S",
			IsNew:        true,
			IsBestSeller: false,
			AvgRating:    4.6,
			TotalReviews: 185,
		},
	}

	if err := db.Create(&gifts).Error; err != nil {
		log.Fatalf("failed to seed gifts: %v", err)
	}
	log.Printf("seeded %d gifts", len(gifts))
}
