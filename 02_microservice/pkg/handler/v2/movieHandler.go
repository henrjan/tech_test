package handler

import (
	"context"

	pb "github.com/henrjan/microservice/internal/proto"
	"github.com/henrjan/microservice/pkg/service"
)

type MovieHandler struct {
	movieSrv  service.MovieSrv
	accessSrv service.AccessSrv
	pb.UnimplementedSearchMovieServer
}

func NewMovieHandler(movieSrv service.MovieSrv, accessSrv service.AccessSrv) *MovieHandler {
	return &MovieHandler{
		movieSrv:  movieSrv,
		accessSrv: accessSrv,
	}
}

func (hand *MovieHandler) GetMovie(ctx context.Context, req *pb.SearchRequest) (res *pb.SearchResponse, err error) {
	res = &pb.SearchResponse{
		Result: make([]*pb.Movie, 0),
	}

	query := make(map[string]interface{})
	query["search_word"] = req.SearchWord
	query["page"] = req.Page

	result, e := hand.movieSrv.GetMovie(query)
	if e != nil {
		res.Error = e.Error()
		return
	}

	for _, v := range result {
		movie := &pb.Movie{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbId,
			Type:   v.Type,
			Poster: v.Poster,
		}
		res.Result = append(res.Result, movie)
	}

	return
}
