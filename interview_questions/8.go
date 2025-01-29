package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	StoryTitle string `json:"storytitle"`
}

type PaginatedResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Article `json:"data"`
}

func getArticleTitles(author string) {
	baseURL := "https://jsonmock.hackerrank.com/api/articles"
	page := 1
	var title []string

	for {
		url := fmt.Sprintf("%s?author=%s&page=%d", baseURL, author, page)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		}
		defer resp.Body.Close()

		var response PaginatedResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Println("error decoding the response body")
		}

		for _, article := range response.Data {
			if article.Title != "" {
				title = append(title, article.Title)
			} else if article.StoryTitle != "" {
				title = append(title, article.StoryTitle)
			}
		}

		if page >= response.TotalPages {
			break
		}
		page++

	}

	fmt.Println(title)

}

// func main() {
// 	getArticleTitles("olalonde")
// }
