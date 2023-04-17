package service

import (
	"fmt"

	"use/internal/model"
	"use/internal/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUser(id int64) (*model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) ListUsers(page, perPage uint64) ([]model.User, error) {
	return s.userRepository.FindAll(page, perPage)
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepository.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}
	if err != nil && err.Error() != "user not found" {
		return nil, fmt.Errorf("failed to check user existence: %v", err)
	}

	// Save user to database
	user, err = s.userRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}
	return user, err
}

func (s *UserService) UpdateUser(id int64, user *model.User) (*model.User, error) {
	// Check if user exists
	existingUser, err := s.userRepository.FindByID(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user with ID %d not found", user.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %v", err)
	}

	if existingUser.Email == user.Email {
		user.Email = ""
	}

	if existingUser.UserName == user.UserName {
		user.UserName = ""
	}

	if existingUser.FirstName == user.FirstName {
		user.FirstName = ""
	}

	if existingUser.LastName == user.LastName {
		user.LastName = ""
	}

	if existingUser.UserStatus == user.UserStatus {
		user.UserStatus = ""
	}

	if existingUser.Department == user.Department {
		user.Department = ""
	}

	user.ID = id
	// Update user in database
	user, err = s.userRepository.Update(user)
	return user, err
}

func (s *UserService) DeleteUser(userID int64) error {
	// Check if user exists
	existingUser, err := s.userRepository.FindByID(userID)
	if existingUser == nil {
		return fmt.Errorf("user with ID %d not found", userID)
	}
	if err != nil {
		return fmt.Errorf("failed to check user existence: %v", err)
	}

	// Delete user from database
	err = s.userRepository.Delete(userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
