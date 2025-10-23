package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Task - database model
type Task struct {
  ID          bson.ObjectID  `bson:"_id,omitempty"`
  UserID      bson.ObjectID  `bson:"user_id"`
  Title       string         `bson:"title"`
  Description string         `bson:"description"`
  Status      string         `bson:"status"`      // pending, in_progress, completed
  Priority    string         `bson:"priority"`    // low, medium, high
  DueDate     *time.Time     `bson:"due_date,omitempty"`
  Tags        []string       `bson:"tags"`
  CreatedAt   time.Time      `bson:"created_at"`
  UpdatedAt   time.Time      `bson:"updated_at"`
  CompletedAt *time.Time     `bson:"completed_at,omitempty"`
}
