package report

import (
	"cyborch.com/apocalypse/pkg/db"
	"cyborch.com/apocalypse/pkg/user"
)

// ReportFlaggedUserPercentage returns the percentage of users
// that are flagged as infected or not
// given a database connection
// and a boolean indicating whether to report the percentage of infected users
// or the percentage of uninfected users
// and an error if any occurs during the query
func ReportFlaggedUserPercentage(database *db.PostgresSql, infected bool) float64 {
	var users []user.User
	database.Client.Find(&users)
	var flaggedUsers []user.User
	for _, u := range users {
		if infected {
			if u.FlagCount >= 3 {
				flaggedUsers = append(flaggedUsers, u)
			}
		} else {
			if u.FlagCount < 3 {
				flaggedUsers = append(flaggedUsers, u)
			}
		}
	}
	return float64(len(flaggedUsers)) / float64(len(users))
}
