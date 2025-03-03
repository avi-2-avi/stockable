package adapters

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TruAdapter struct {
	apiURL       string
	token        string
	service      *services.AnalystRatingsService
	dataSourceID uint
}

func NewTruAdapter(apiURL string, token string, service *services.AnalystRatingsService, dataSourceID uint) RatingAdapter {
	return &TruAdapter{
		apiURL:       apiURL,
		token:        token,
		service:      service,
		dataSourceID: dataSourceID,
	}
}

func (truAdapter *TruAdapter) FetchData() ([]models.AnalystRating, error) {
	var allRatings []models.AnalystRating
	nextPage := ""
	// Temporal counter
	// count := 0

	for {
		url := truAdapter.buildUrl(nextPage)
		println("URL: ", url)

		response, err := truAdapter.callAPI(url)
		if err != nil {
			return nil, err
		}

		ratings, newNextPage, err := truAdapter.parseResponse(response)
		println("Ratings: ", ratings)
		if err != nil {
			return nil, err
		}

		allRatings = append(allRatings, ratings...)
		if newNextPage == "" {
			break
		}
		nextPage = newNextPage

		// Testing if the loop is working
		// count++
		// if count > 10 {
		// 	break
		// }
	}

	err := truAdapter.service.SaveAnalystRatingsBatch(allRatings)
	if err != nil {
		return nil, err
	}

	return allRatings, nil
}

func (truAdapter *TruAdapter) buildUrl(nextPage string) string {
	if nextPage == "" {
		return truAdapter.apiURL
	}
	return fmt.Sprintf("%s?next_page=%s", truAdapter.apiURL, nextPage)
}

func (truAdapter *TruAdapter) callAPI(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+truAdapter.token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ratings: %w", err)
	}

	return response, nil
}

func (truAdapter *TruAdapter) parseResponse(httpResponse *http.Response) ([]models.AnalystRating, string, error) {
	defer httpResponse.Body.Close()

	var response models.TruAPIResponse

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, "", fmt.Errorf("failed to decode response: %w", err)
	}

	var ratings []models.AnalystRating
	for _, item := range response.Items {
		parsedTime, _ := time.Parse(time.RFC3339, item.Time)
		ratings = append(ratings, models.AnalystRating{
			Ticker:       item.Ticker,
			TargetFrom:   utils.ParsePrice(item.TargetFrom),
			TargetTo:     utils.ParsePrice(item.TargetTo),
			Company:      item.Company,
			Action:       item.Action,
			Brokerage:    item.Brokerage,
			RatingFrom:   item.RatingFrom,
			RatingTo:     item.RatingTo,
			RatedAt:      parsedTime,
			DataSourceID: truAdapter.dataSourceID,
		})
	}

	return ratings, response.NextPage, nil
}
