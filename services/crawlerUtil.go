package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetFoodMap() map[string]string {
	foodMap := make(map[string]string)

	url := "https://heyform.net/f/f9gtvNqd"
	res, err := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	// Extract JSON from HTML
	jsonData := extractJSONFromHTML(string(body))

	// Parse the JSON data
	var formData formData
	err = json.Unmarshal([]byte(jsonData), &formData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return foodMap
	}

	// Process all fields and their choices
	for _, field := range formData.Props.PageProps.Form.Fields {
		// Skip fields without choices
		if len(field.Properties.Choices) == 0 {
			continue
		}

		// Add all choices to the map
		for _, choice := range field.Properties.Choices {
			foodMap[choice.Label] = choice.ID
		}
	}
	return foodMap
}

type choice struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// field represents a category of food items
type field struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Properties struct {
		Choices []choice `json:"choices,omitempty"`
	} `json:"properties"`
}

// form represents the main form structure
type form struct {
	Fields []field `json:"fields"`
}

// formData represents the top-level structure
type formData struct {
	Props struct {
		PageProps struct {
			Form form `json:"form"`
		} `json:"pageProps"`
	} `json:"props"`
}

func extractJSONFromHTML(htmlContent string) string {
	start := strings.Index(htmlContent, `<script id="__NEXT_DATA__" type="application/json">`)
	if start == -1 {
		return ""
	}
	start += len(`<script id="__NEXT_DATA__" type="application/json">`)

	end := strings.Index(htmlContent[start:], "</script>")
	if end == -1 {
		return ""
	}

	return htmlContent[start : start+end]
}
