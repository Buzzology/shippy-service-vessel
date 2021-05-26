package main

import (
	"context"
	pb "github.com/buzzology/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

func (s *handler) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {

	// Save
	if err := s.repository.Create(ctx, MarshalVessel(req)); err != nil {
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

	res.Vessel = UnmarshalVessel(vessel)
	return nil
}