package handlers

import (
	"os"
	"encoding/json"
	"github.com/whalerapi/templ-go/models"
)



func getPosts() ([]Post, error) {
    jsonData, err := os.ReadFile("allposts.json")
    if err != nil {
        return nil, err
    }

    var posts []models.Post
    if err := json.Unmarshal(jsonData, &posts); err != nil {
        return nil, err
    }
    return posts, nil
}
