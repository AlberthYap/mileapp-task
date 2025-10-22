package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"task-api/models"
)

// UserRepository - interface for user repository
type UserRepository interface {
  FindByEmail(ctx context.Context, email string) (*models.User, error)
}

// userRepository - implement UserRepository
type userRepository struct {
  db         *mongo.Database
  collection *mongo.Collection
}

// NewUserRepository - constructor
func NewUserRepository(db *mongo.Database) UserRepository {
  return &userRepository{
    db:         db,
    collection: db.Collection("users"),
  }
}

// FindByEmail - find user base on email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
  var user models.User
  
  filter := bson.M{"email": email}
  err := r.collection.FindOne(ctx, filter).Decode(&user)
  
  if err != nil {
    if err == mongo.ErrNoDocuments {
      return nil, errors.New("user not found")
    }
    return nil, err
  }
  
  return &user, nil
}
