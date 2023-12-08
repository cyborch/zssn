package api

import (
	"net/http"

	"cyborch.com/apocalypse/pkg/db"
	"cyborch.com/apocalypse/pkg/report"
	"cyborch.com/apocalypse/pkg/user"
	"github.com/gin-gonic/gin"
)

type PercentageResponse struct {
	Percentage float64 `json:"percentage"`
}

type AveragesResponse struct {
	Averages []user.UserItemAverage `json:"averages"`
}

// @Summary report (un)infected percentage
// @Description Report the percentage of (un)infected users
// @Accept json
// @Produce json
// @Param infected query bool true "infected"
// @Response 200 {object} PercentageResponse
// @Router /report/percentage [get]
func ReportFlaggedUserPercentage(c *gin.Context, database *db.PostgresSql) {
	infected := c.Query("infected") == "true"
	c.JSON(http.StatusOK, gin.H{
		"percentage": report.ReportFlaggedUserPercentage(database, infected) * 100.0,
	})
}

// @Summary report averages
// @Description Report the average number resources per user
// @Produce json
// @Response 200 {object} AveragesResponse
// @Router /report/averages [get]
func ReportAverages(c *gin.Context, database *db.PostgresSql) {
	averages, err := db.GetAverages(database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"averages": averages,
	})
}

// @Summary report lost value of items because of infected users
// @Description Report the total value of lost items because of infected users
// @Produce json
// @Response 200 {object} user.LostItemValue
// @Router /report/lost [get]
func ReportLostValue(c *gin.Context, database *db.PostgresSql) {
	lost, err := db.GetLostItems(database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lost": lost.Lost,
	})
}
