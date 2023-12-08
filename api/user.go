package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"cyborch.com/apocalypse/pkg/db"
	"cyborch.com/apocalypse/pkg/user"
)

type Response struct {
	Message string `json:"message"`
}

type RegistrationResponse struct {
	Message string `json:"message"`
	ID      uint   `json:"id"`
}

type FlagRequest struct {
	SenderID uint `json:"sender_id"`
}

// @Summary register
// @Description Register a user, where items is an array of integers, each representing an item
// @Accept json
// @Produce json
// @Param user body user.User true "user"
// @Success 200 {object} RegistrationResponse
// @Router /user/register [post]
func Register(c *gin.Context, database *db.PostgresSql) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Client.Create(&u)

	c.JSON(http.StatusOK, RegistrationResponse{
		Message: "user registered",
		ID:      u.ID,
	})
}

// @Summary update location
// @Description Update the location of a user
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param location body user.Location true "location"
// @Success 200 {object} Response
// @Router /user/{id}/location [put]
func UpdateLocation(c *gin.Context, database *db.PostgresSql) {
	user_id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	u, err := db.GetUser(database, uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&u.Location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.Client.Model(&u).Save(&u)

	c.JSON(http.StatusOK, Response{
		Message: "location updated",
	})
}

// @Summary Flag another user
// @Description Flag a user as infected
// @Param id path int true "user id"
// @Param flag body FlagRequest true "flag request"
// @Success 200 {object} Response
// @Router /user/{id}/flag [post]
func Flag(c *gin.Context, database *db.PostgresSql) {
	user_id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	u, err := db.GetUser(database, uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var flag FlagRequest
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if flag.SenderID == u.ID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you cannot flag yourself"})
		return
	}

	database.Client.Model(&u).Update("flag_count", u.FlagCount+1)

	c.JSON(http.StatusOK, Response{
		Message: "user flagged",
	})
}

// @Summary Trade with another user
// @Description Trade items with another user
// @Param trade body user.TradeRequest true "trade request"
// @Success 200 {object} Response
// @Router /user/trade [post]
func Trade(c *gin.Context, database *db.PostgresSql) {
	var trade user.TradeRequest
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if trade.RecepientID == trade.SenderID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you cannot trade with yourself"})
		return
	}

	if len(trade.RequestedItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you must request at least one item"})
		return
	}

	if len(trade.OfferedItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you must offer at least one item"})
		return
	}

	// verify that the trade is valid
	if !user.VerifyTradeValue(trade) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you must offer the same value of items you request"})
		return
	}

	sender, err := db.GetUser(database, trade.SenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sender not found"})
		return
	}
	recepient, err := db.GetUser(database, trade.RecepientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "recepient not found"})
		return
	}

	if !user.VerifyNonInfectedUser(sender) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you cannot trade while infected"})
		return
	}
	if !user.VerifyNonInfectedUser(recepient) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you cannot trade with an infected user"})
		return
	}

	// remove the offered items from the sender
	// if the sender does not have the items, return an error
	// and do not proceed with the trade
	if !user.RemoveTradedItemsFromUser(sender, trade.OfferedItems) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you do not have the items you are offering"})
		return
	}
	// remove the requested items from the recepient
	// if the recepient does not have the items, return an error
	// and do not proceed with the trade
	if !user.RemoveTradedItemsFromUser(recepient, trade.RequestedItems) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the recepient does not have the items you are requesting"})
		return
	}

	user.AddTradedItemsToUser(recepient, trade.OfferedItems)
	user.AddTradedItemsToUser(sender, trade.RequestedItems)

	// save the changes to the database
	database.Client.Save(&sender)
	database.Client.Save(&recepient)

	c.JSON(http.StatusOK, Response{
		Message: "trade successful",
	})
}
