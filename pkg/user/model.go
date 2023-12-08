package user

type ItemType uint

const (
	Ammunition ItemType = 1
	Medication          = 2
	Food                = 3
	Water               = 4
)

// UserItem is a struct that represents an item a user has
type UserItem struct {
	ID      uint     `json:"-" gorm:"primaryKey"`
	Item    ItemType `json:"item" gorm:"not null;type:integer"` // The item type
	User_id uint     `json:"-" gorm:"not null; index:idx_user_id"`
}

// Location is a struct that represents a location in the world
type Location struct {
	Lat float64 `json:"lat" gorm:"not null"` // Latitude
	Lon float64 `json:"lon" gorm:"not null"` // Longitude
}

// User is a struct that represents a user
type User struct {
	ID uint `json:"-" gorm:"primary_key"`
	// Your name
	Name string `json:"name" gorm:"not null"`
	// Your age
	Age int `json:"age" gorm:"not null"`
	// Even in a post-apocalyptic world, you can be whatever you want
	Gender string `json:"gender" gorm:"not null"`
	// Your location in the world
	Location Location `json:"location" gorm:"embedded"`
	/*
		Items you have, can be ammunition (1), medication (2),
		food (3) or water (4), each can appear multiple times
		in case you have more than one
	*/
	Items     []UserItem `json:"items" gorm:"foreignKey:user_id"`
	FlagCount int        `json:"-" gorm:"not null;default:0"`
}

// TradeRequest is a struct that represents a trade request
type TradeRequest struct {
	SenderID       uint       `json:"sender_id"`       // The ID of the user who sent the trade request
	RecepientID    uint       `json:"recepient_id"`    // The ID of the user who to trade with
	RequestedItems []UserItem `json:"requested_items"` // The items the sender wants to receive
	OfferedItems   []UserItem `json:"offered_items"`   // The items the sender offers to give
}

// UserItemAverage is a struct that represents the average number of items per user
type UserItemAverage struct {
	Item    ItemType `json:"item"`
	Average float64  `json:"average"`
}

// LostItemValue is a struct that represents the total value of lost items
// because of infected users
type LostItemValue struct {
	Lost int `json:"lost"`
}
