package main

import (
	pb "github.com/amiralii/shipping/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shipping"
	vesselCollection = "vessels"
)

type Repository interface {
	Find(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) Find(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxWeight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) Close() {
	repo.session.Clone()
}

func (repo *VesselRepository) collection() *mgo.Collation {
	return repo.session.DB(dbName).C(vesselCollection)
}
