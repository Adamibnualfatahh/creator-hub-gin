package services

import (
	"creator-hub-gin/data"
	"creator-hub-gin/models"
	"time"
)

func GetAllContents() []models.Content {
	return data.Contents
}

func GetContentByID(id int) *models.Content {
	for i, c := range data.Contents {
		if c.ID == id {
			return &data.Contents[i]
		}
	}
	return nil
}

func CreateContent(newContent models.Content) models.Content {
	newID := 1
	if len(data.Contents) > 0 {
		newID = data.Contents[len(data.Contents)-1].ID + 1
	}
	newContent.ID = newID
	if newContent.PublishDate.IsZero() {
		newContent.PublishDate = time.Now()
	}
	data.Contents = append(data.Contents, newContent)
	return newContent
}

func UpdateContent(id int, updated models.Content) *models.Content {
	for i, c := range data.Contents {
		if c.ID == id {
			updated.ID = id
			data.Contents[i] = updated
			return &data.Contents[i]
		}
	}
	return nil
}

func DeleteContent(id int) bool {
	for i, c := range data.Contents {
		if c.ID == id {
			data.Contents = append(data.Contents[:i], data.Contents[i+1:]...)
			return true
		}
	}
	return false
}
