package main

import (
	"context"

	pb "github.com/amiralii/shipping/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailabel(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	vessel, err := repo.Find(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func (s *service) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	// repo := s.GetRepo()
	// defer repo.Close()
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
