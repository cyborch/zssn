package user

import (
	"testing"
)

func TestRemoveItemFromUser(t *testing.T) {
	// Create a user with some initial items
	user := &User{
		Items: []UserItem{
			{Item: Ammunition},
			{Item: Medication},
			{Item: Food},
			{Item: Water},
		},
	}

	// Test removing an existing item
	existingItem := UserItem{Item: Ammunition}
	result := RemoveItemFromUser(user, existingItem)
	if !result {
		t.Errorf("Expected item %v to be removed, but it was not", existingItem)
	}

	// Check if the item was removed from the user's items
	expectedItems := []UserItem{
		{Item: Medication},
		{Item: Food},
		{Item: Water},
	}
	assertItemsEqual(t, user.Items, expectedItems)

	// Test removing an item that was already removed
	result = RemoveItemFromUser(user, existingItem)
	if result {
		t.Errorf("Expected item %v to not be found, but it was removed", existingItem)
	}

	// Check if the user's items remain unchanged
	assertItemsEqual(t, user.Items, expectedItems)
}

func TestAddItemToUser(t *testing.T) {
	// Create a user with no items
	user := &User{}

	// Add an item to the user
	newItem := UserItem{Item: Ammunition}
	AddItemToUser(user, newItem)

	// Check if the item was added to the user's items
	expectedItems := []UserItem{
		{Item: Ammunition},
	}
	assertItemsEqual(t, user.Items, expectedItems)
}

// Test that the trade value is verified correctly
// by creating a valid trade and tetsing if the verify function
// returns true
func TestValidTrade(t *testing.T) {
	// Create a valid trade
	trade := TradeRequest{
		RequestedItems: []UserItem{
			{Item: Food},
			{Item: Medication},
			{Item: Medication},
		},
		OfferedItems: []UserItem{
			{Item: Ammunition},
			{Item: Medication},
			{Item: Water},
		},
	}

	// Verify the trade is valid
	result := VerifyTradeValue(trade)
	if !result {
		t.Errorf("Expected trade %v to be valid, but it was not", trade)
	}
}

// Test that the trade value is verified correctly
// by creating an invalid trade and tetsing if the verify function
// returns false
func TestInvalidTrade(t *testing.T) {
	// Create an invalid trade
	trade := TradeRequest{
		RequestedItems: []UserItem{
			{Item: Ammunition},
		},
		OfferedItems: []UserItem{
			{Item: Water},
		},
	}

	// Verify the trade is invalid
	result := VerifyTradeValue(trade)
	if result {
		t.Errorf("Expected trade %v to be invalid, but it was valid", trade)
	}
}

// Helper function to check if two slices of items are equal
func assertItemsEqual(t *testing.T, actual, expected []UserItem) {
	if len(actual) != len(expected) {
		t.Errorf("Expected items %v, but got %v", expected, actual)
		return
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected items %v, but got %v", expected, actual)
			return
		}
	}
}

// Test that a non-infected user can trade
func TestVerifyNonInfectedUser(t *testing.T) {
	// Create a user with no flags
	user := &User{
		FlagCount: 0,
	}

	// Verify the user is non-infected
	result := VerifyNonInfectedUser(user)
	if !result {
		t.Errorf("Expected user %v to be non-infected, but it was infected", user)
	}
}

// Test that an infected user cannot trade
func TestVerifyInfectedUser(t *testing.T) {
	// Create a user with 3 flags
	user := &User{
		FlagCount: 3,
	}

	// Verify the user is infected
	result := VerifyNonInfectedUser(user)
	if result {
		t.Errorf("Expected user %v to be infected, but it was non-infected", user)
	}
}
