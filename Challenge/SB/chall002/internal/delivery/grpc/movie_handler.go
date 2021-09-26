package grpc

import (
	context "context"

	"github.com/syafdia/sb/chal002/internal/domain/usecase"
)

type movieHandler struct {
	findOneMovieUseCase       usecase.FindOneMovieUseCase
	findMultipleMoviesUseCase usecase.FindMultipleMoviesUseCase
}

func NewMovieHandler(
	findOneMovieUseCase usecase.FindOneMovieUseCase,
	findMultipleMoviesUseCase usecase.FindMultipleMoviesUseCase,
) MovieServiceServer {
	return &movieHandler{
		findOneMovieUseCase:       findOneMovieUseCase,
		findMultipleMoviesUseCase: findMultipleMoviesUseCase,
	}
}

func (m *movieHandler) FindOne(ctx context.Context, in *MovieFindOneRequest) (*MovieFindOneResponse, error) {
	detail, err := m.findOneMovieUseCase.Execute(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &MovieFindOneResponse{Payload: &detail}, nil
}
func (m *movieHandler) FindMultiple(ctx context.Context, in *MovieFindMultipleRequest) (*MovieFindMultipleResponse, error) {
	summaries, err := m.findMultipleMoviesUseCase.Execute(ctx, in.SearchWord, int(in.Pagination))
	if err != nil {
		return nil, err
	}

	response := &MovieFindMultipleResponse{}
	for _, summary := range summaries {
		response.Payload = append(response.Payload, &summary)
	}

	return response, nil
}
