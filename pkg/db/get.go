package db

import (
	"errors"

	"cyborch.com/apocalypse/pkg/user"
)

// GetUser returns a user from the database
// given a user id
// and an error if any occurs during the query
func GetUser(database *PostgresSql, user_id uint) (*user.User, error) {
	var users []user.User

	if err := database.
		Client.
		Model(&user.User{}).
		Preload("Items").
		Where("id = ?", user_id).
		Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("record not found")
	}

	return &users[0], nil
}

// GetAverages returns the average number of items per user
// for each type of item
// and an error if any occurs during the query
func GetAverages(database *PostgresSql) ([]user.UserItemAverage, error) {
	var averages []user.UserItemAverage

	if err := database.Client.Raw(`
		select 
			item,
			count(1)::decimal / (select count(1) from users)::decimal as average
		from user_items
		join users on user_items.user_id = users.id
		where users.flag_count < 3
		group by item
	`).Scan(&averages).Error; err != nil {
		return nil, err
	}

	return averages, nil
}

func GetLostItems(database *PostgresSql) (*user.LostItemValue, error) {
	var lost user.LostItemValue

	if err := database.Client.Raw(`
		select 
			coalesce(sum(user_items.item), 0) as lost
		from user_items
		join users on users.id = user_items.user_id
		where users.flag_count >= 3
	`).First(&lost).Error; err != nil {
		return nil, err
	}

	return &lost, nil
}
