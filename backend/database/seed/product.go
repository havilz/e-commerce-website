package seed

import (
	"fmt"
	"log"

	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("[seed] products already exist, skipping...")
		return
	}

	products := []models.Product{
		{
			Name:        "Wireless Headphones",
			Description: "Premium noise-cancelling wireless headphones with 30h battery life",
			Price:       999000,
			Stock:       50,
			ImageURL:    "https://picsum.photos/seed/1/400/400",
			Category:    "Electronics",
		},
		{
			Name:        "Mechanical Keyboard",
			Description: "TKL mechanical keyboard with Cherry MX switches and RGB backlight",
			Price:       750000,
			Stock:       30,
			ImageURL:    "https://picsum.photos/seed/2/400/400",
			Category:    "Electronics",
		},
		{
			Name:        "Running Shoes",
			Description: "Lightweight running shoes with cushioned sole",
			Price:       650000,
			Stock:       100,
			ImageURL:    "https://picsum.photos/seed/3/400/400",
			Category:    "Fashion",
		},
		{
			Name:        "Stainless Water Bottle",
			Description: "500ml double-wall insulated stainless steel water bottle",
			Price:       120000,
			Stock:       200,
			ImageURL:    "https://picsum.photos/seed/4/400/400",
			Category:    "Lifestyle",
		},
		{
			Name:        "Backpack 30L",
			Description: "Waterproof laptop backpack 30L with USB charging port",
			Price:       450000,
			Stock:       75,
			ImageURL:    "https://picsum.photos/seed/5/400/400",
			Category:    "Fashion",
		},
	}

	names := []string{
		"Gaming Mouse", "Smart Watch", "Bluetooth Speaker", "Laptop Stand",
		"USB-C Hub", "Portable SSD", "Webcam HD", "Wireless Charger",
		"Gaming Headset", "Power Bank", "Tablet Stand", "Smart Bulb",
		"Mini Projector", "Security Camera", "Air Fryer", "Rice Cooker",
		"Office Chair", "Study Desk", "Bookshelf", "Wall Clock",
		"Yoga Mat", "Dumbbell Set", "Resistance Band", "Basketball",
		"Football", "Tennis Racket", "Camping Tent", "Cycling Helmet",
		"Hoodie Premium", "Casual Sneakers", "Denim Jacket", "Cargo Pants",
		"Cotton T-Shirt", "Leather Belt", "Baseball Cap", "Sports Socks",
		"Coffee Mug", "Desk Lamp", "Notebook Journal", "Storage Box",
		"Face Wash", "Moisturizer", "Body Lotion", "Lip Balm",
		"Hair Dryer", "Perfume Classic", "Shampoo Herbal", "Sunscreen SPF50",
		"Gaming Chair", "Gaming Monitor", "Streaming Microphone", "VR Headset",
		"Go Programming Book", "Clean Code", "System Design Guide", "Docker Guide",
		"Kubernetes Handbook", "Cloud Computing Basics", "Design Patterns",
		"Pet Bed", "Pet Carrier", "Dog Leash", "Cat Toy",
		"Electric Kettle", "Vacuum Cleaner", "Floor Lamp", "Kitchen Knife Set",
		"Cookware Set", "Lunch Box", "Phone Stand", "Portable Fan",
		"Air Humidifier", "Digital Thermometer", "Smart Plug", "Monitor Arm",
		"Wireless Earbuds", "Mechanical Numpad", "RGB Mouse Pad", "Capture Card",
		"Graphic Tablet", "External HDD", "USB Microphone", "LED Strip Light",
		"Fitness Tracker", "Jump Rope", "Hiking Pole", "Travel Pillow",
		"Scented Candle", "Alarm Clock", "Whiteboard", "Desk Organizer",
		"Smart Scale", "Massage Gun", "Protein Shaker", "Water Filter",
		"Portable Blender",
	}

	categories := []string{
		"Electronics",
		"Fashion",
		"Lifestyle",
		"Home",
		"Sports",
		"Books",
		"Beauty",
		"Gaming",
	}

	for i, name := range names {
		products = append(products, models.Product{
			Name:        name,
			Description: fmt.Sprintf("%s - high quality product for everyday use", name),
			Price:       float64(100000 + ((i + 1) * 25000)),
			Stock:       20 + (i % 100),
			ImageURL:    fmt.Sprintf("https://picsum.photos/seed/product%d/400/400", i+6),
			Category:    categories[i%len(categories)],
		})
	}

	if err := db.Create(&products).Error; err != nil {
		log.Println("[seed] failed:", err)
		return
	}

	log.Printf("[seed] %d sample products created", len(products))
}