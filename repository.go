package main

import (
	"context"
	pb "github.com/Buzzology/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

// Mongo implementation
type MongoRepository struct {
	collection *mongo.Collection
}

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
}

type Specification struct {
	Capacity  int32 `json:capacity`
	MaxWeight int32 `json:max_weight`
}

func MarshalVesselCollection(vessels []*pb.Vessel) []*Vessel {
	collection := make([]*Vessel, 0)
	for _, vessel := range vessels {
		collection = append(collection, MarshalVessel(vessel))
	}
	return collection
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalVesselCollection(vessels []*Vessel) []*pb.Vessel {
	collection := make([]*pb.Vessel, 0)
	for _, vessel := range vessels {
		collection = append(collection, UnmarshalVessel(vessel))
	}
	return collection
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}


func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
			bson.D{{
				"$gte",
				spec.Capacity,
			}},
		},
		{
			"maxweight",
			bson.D{{
				"$gte",
				spec.MaxWeight,
			}},
		}}

	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}
