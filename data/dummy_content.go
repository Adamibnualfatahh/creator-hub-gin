package data

import (
	"creator-hub-gin/models"
	"time"
)

var Contents = []models.Content{
	{
		ID:          1,
		IdeaTitle:   "Konten 1",
		Category:    "Edukasi",
		Platform:    "Instagram",
		Status:      "Idea",
		PublishDate: time.Now(),
	},
	{
		ID:          2,
		IdeaTitle:   "Konten 2",
		Category:    "Hiburan",
		Platform:    "Facebook",
		Status:      "Scheduled",
		PublishDate: time.Now().Add(24 * time.Hour),
	},
	{
		ID:          3,
		IdeaTitle:   "Konten 3",
		Category:    "Motivasi",
		Platform:    "Twitter",
		Status:      "Posted",
		PublishDate: time.Now().Add(-48 * time.Hour),
		EngagementData: models.Engagement{
			Likes:    300,
			Comments: 150,
			Shares:   75,
		},
	},
}
