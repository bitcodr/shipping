package main

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/amiralii/shipping/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := s.tokenService.Encode(user)
	if err != nil {
		return nil
	}
	res.Token = token
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Generates a hashed version of our password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	req.Password = string(hashPass)
	if err := s.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
