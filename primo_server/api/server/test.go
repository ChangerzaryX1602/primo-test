package server

import (
	"context"
	"errors"
	"test/internal/usecase"
	"test/utils"

	helpers "github.com/zercle/gofiber-helpers"
)

type testServer struct {
	usecase usecase.TestUsecase
}

func NewTestServer(usecase usecase.TestUsecase) TestServiceServer {
	return &testServer{usecase: usecase}
}
func (s *testServer) mustEmbedUnimplementedTestServiceServer() {}
func (s *testServer) Merge(ctx context.Context, req *SortRequest) (*SortResponse, error) {
	if req.Collection_1 == nil {
		return nil, utils.NewErrorWithSource(errors.New("collection_1 is required"), helpers.WhereAmI())
	}
	if req.Collection_2 == nil {
		return nil, utils.NewErrorWithSource(errors.New("collection_2 is required"), helpers.WhereAmI())
	}
	if req.Collection_3 == nil {
		return nil, utils.NewErrorWithSource(errors.New("collection_3 is required"), helpers.WhereAmI())
	}
	result, err := s.usecase.Merge(req.Collection_1, req.Collection_2, req.Collection_3)
	if err != nil {
		return nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return &SortResponse{SortedCollection: result}, nil
}
