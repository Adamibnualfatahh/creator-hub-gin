package models

import "time"

type Content struct {
	ID             int        `json:"id"`
	IdeaTitle      string     `json:"idea_title"`
	Category       string     `json:"category"`
	Platform       string     `json:"platform"`
	Status         string     `json:"status"` // "Idea", "Scheduled", "Posted"
	PublishDate    time.Time  `json:"publish_date"`
	EngagementData Engagement `json:"engagement"`
}

func (c Content) TotalEngagement() int {
	return c.EngagementData.Likes + c.EngagementData.Comments + c.EngagementData.Shares
}
