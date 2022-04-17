package transaction

import user "rexencorpstartup/User"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}