package user

// RemoveItemFromUser removes an item from a user
// and returns true if the item was removed
// and false if the item was not found
func RemoveItemFromUser(user *User, item UserItem) bool {
	for i, v := range user.Items {
		if v.Item == item.Item {
			user.Items = append(user.Items[:i], user.Items[i+1:]...)
			return true
		}
	}
	return false
}

// AddItemToUser adds an item to a user
func AddItemToUser(user *User, item UserItem) {
	user.Items = append(user.Items, item)
}

// VerifyTradeValue verifies that the trade offer
// is equal to the trade request
// and returns true if the trade is valid
func VerifyTradeValue(trade TradeRequest) bool {
	tradeOffer := 0
	for _, user_item := range trade.OfferedItems {
		tradeOffer += int(user_item.Item)
	}
	tradeRequest := 0
	for _, user_item := range trade.RequestedItems {
		tradeRequest += int(user_item.Item)
	}

	return tradeOffer == tradeRequest
}

// RemoveTradedItemsFromUser removes the offered items
// from the user and returns true if the items were available
// and false if the items were not available
func RemoveTradedItemsFromUser(user *User, items []UserItem) bool {
	for _, user_item := range items {
		if !RemoveItemFromUser(user, user_item) {
			return false
		}
	}
	return true
}

// AddTradedItemsToUser adds the requested items
// to the user
func AddTradedItemsToUser(user *User, items []UserItem) {
	for _, item := range items {
		AddItemToUser(user, item)
	}
}

// VerifyNonInfectedUser verifies that the user is not infected
// and returns true if the user is not infected
func VerifyNonInfectedUser(user *User) bool {
	return user.FlagCount < 3
}
