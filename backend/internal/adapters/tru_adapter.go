package adapters

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TruAdapter struct {
	apiURL         string
	token          string
	ratingService  services.AnalystRatingsService
	companyService services.CompanyService
	dataSourceID   uuid.UUID
}

func NewTruAdapter(apiURL string, token string, ratingService services.AnalystRatingsService, companyService services.CompanyService, dataSourceID uuid.UUID) RatingAdapter {
	return &TruAdapter{
		apiURL:         apiURL,
		token:          token,
		ratingService:  ratingService,
		companyService: companyService,
		dataSourceID:   dataSourceID,
	}
}

func (truAdapter *TruAdapter) FetchData() ([]models.AnalystRating, error) {
	var allRatings []models.AnalystRating
	nextPage := ""

	for {
		url := truAdapter.buildUrl(nextPage)

		response, err := truAdapter.callAPI(url)
		if err != nil {
			return nil, err
		}

		ratings, newNextPage, err := truAdapter.parseResponse(response)
		if err != nil {
			return nil, err
		}

		allRatings = append(allRatings, ratings...)
		if newNextPage == "" {
			break
		}
		nextPage = newNextPage
	}

	err := truAdapter.ratingService.SaveAnalystRatingsBatch(allRatings)
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

		existingCompany, err := truAdapter.companyService.CreateCompanyByTicker(item.Ticker, item.Company)
		if err != nil {
			return nil, "", err
		}

		targetFrom := utils.ParsePrice(item.TargetFrom)
		targetTo := utils.ParsePrice(item.TargetTo)

		rating := models.AnalystRating{
			TargetFrom:                 targetFrom,
			TargetTo:                   targetTo,
			Action:                     item.Action,
			Brokerage:                  item.Brokerage,
			RatingFrom:                 item.RatingFrom,
			RatingTo:                   item.RatingTo,
			RatedAt:                    parsedTime,
			DataSourceID:               truAdapter.dataSourceID,
			CompanyID:                  existingCompany.ID,
			ActionImpactScore:          utils.CalculateActionImpactScore(item.Action),
			RatingChangeImpact:         utils.CalculateRatingChangeImpact(item.RatingFrom, item.RatingTo),
			TargetAdjustmentPercentage: utils.CalculateTargetAdjustment(targetFrom, targetTo),
		}

		rating.CombinedPredictionIndex = utils.CalculateRawCPI(rating.ActionImpactScore, rating.RatingChangeImpact, rating.TargetAdjustmentPercentage)

		ratings = append(ratings, rating)
	}

	return ratings, response.NextPage, nil
}
