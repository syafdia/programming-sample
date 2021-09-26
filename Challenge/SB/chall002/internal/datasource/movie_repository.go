package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

const movieAPIBaseURL = "http://www.omdbapi.com/?apikey=faf7e5bb"

var httpStatusWithError = map[int]error{
	http.StatusUnauthorized: entity.ErrUnauthorized,
}

type movieRepo struct {
	httpClient *http.Client
}

func NewMovieRepo(httpClient *http.Client) repo.MovieRepository {
	return &movieRepo{httpClient: httpClient}
}

func (m *movieRepo) FindOneDetail(ctx context.Context, id string) (entity.MovieDetail, error) {
	url := fmt.Sprintf("%s&i=%s", movieAPIBaseURL, id)
	var detail entity.MovieDetail

	err := m.getAndDecode(ctx, url, &detail)
	if err != nil {
		return entity.MovieDetail{}, err
	}

	detail.Id = id

	return detail, nil
}

func (m *movieRepo) FindMultipleSummaries(ctx context.Context, searchWord string, pagination int) ([]entity.MovieSummary, error) {
	url := fmt.Sprintf("%s&s=%s&page=%d", movieAPIBaseURL, searchWord, pagination)
	var jsonBody struct {
		Search []struct {
			Title  string
			Year   string
			ImdbID string `json:"imdbID"`
			Type   string
			Poster string
		}
	}

	err := m.getAndDecode(ctx, url, &jsonBody)
	if err != nil {
		return []entity.MovieSummary{}, err
	}

	summaries := []entity.MovieSummary{}
	for _, s := range jsonBody.Search {
		summaries = append(summaries, entity.MovieSummary{
			Id:     s.ImdbID,
			Title:  s.Title,
			Year:   s.Year,
			Type:   s.Type,
			Poster: s.Poster,
		})
	}

	return summaries, nil
}

func (m *movieRepo) getAndDecode(ctx context.Context, url string, target interface{}) error {
	resp, err := m.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("movieRepo: failed on fetch JSON body, err: %w", err)
	}

	defer resp.Body.Close()

	switch {
	case resp.StatusCode == http.StatusUnauthorized:
		return entity.ErrUnauthorized
	case resp.StatusCode >= 500 && resp.StatusCode <= 599:
		return entity.ErrBadGateway
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
