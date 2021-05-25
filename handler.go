package main

import (
	"context"
	pb "github.com/buzzology/go-microservices-tutorial/shippy-service-vessel/proto/vessel"
	"github.com/pkg/errors"
)

type handler struct {
	repository
}

func (s *handler) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {

	// Save
	if err := repository.Create(s.repository, ctx); err != nil {
		return err
	}

	res.Created = true
	res.Vessel = req
	return nil
}


func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Save
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req));
	if err != nil {
		return err
	}

	res.Vessel = UnmarshalVessel(&vessel)
	return nil
}