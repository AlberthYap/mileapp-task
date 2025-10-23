package repositories

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"task-api/models"
	"task-api/types"
)

// TaskRepository - interface
type TaskRepository interface {
  Create(ctx context.Context, task *models.Task) error
  FindByID(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) (*models.Task, error)
  FindByUserID(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) ([]models.Task, int64, error)
  Update(ctx context.Context, id bson.ObjectID, userID bson.ObjectID, updates bson.M) error
  Delete(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) error
}

// taskRepository - implementation
type taskRepository struct {
  collection *mongo.Collection
}

// NewTaskRepository - constructor
func NewTaskRepository(db *mongo.Database) TaskRepository {
  return &taskRepository{
    collection: db.Collection("tasks"),
  }
}

// Create - create new task
func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
  _, err := r.collection.InsertOne(ctx, task)
  return err
}

// FindByID - find task by ID (with user ownership check)
func (r *taskRepository) FindByID(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) (*models.Task, error) {
  var task models.Task
  
  filter := bson.M{
    "_id":     id,
    "user_id": userID, // Ensure user owns this task
  }
  
  err := r.collection.FindOne(ctx, filter).Decode(&task)
  if err != nil {
    if err == mongo.ErrNoDocuments {
      return nil, errors.New("task not found")
    }
    return nil, err
  }
  
  return &task, nil
}

// FindByUserID - find tasks by user ID with filters
func (r *taskRepository) FindByUserID(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) ([]models.Task, int64, error) {
  // Build filter
  filter := bson.M{"user_id": userID}
  
  if query.Status != "" {
    filter["status"] = query.Status
  }
  
  if query.Priority != "" {
    filter["priority"] = query.Priority
  }
  
  if query.Search != "" {
    filter["$or"] = []bson.M{
      {"title": bson.M{"$regex": query.Search, "$options": "i"}},
      {"description": bson.M{"$regex": query.Search, "$options": "i"}},
    }
  }
  
  // Count total
  total, err := r.collection.CountDocuments(ctx, filter)
  if err != nil {
    return nil, 0, err
  }
  
  // Pagination
  page := 1
  if query.Page > 0 {
    page = query.Page
  }
  
  limit := 10
  if query.Limit > 0 {
    limit = query.Limit
  }
  
  skip := (page - 1) * limit
  
  // Sort
  sortField := "created_at"
  sortOrder := -1
  if query.Sort != "" {
    if query.Sort[0] == '-' {
      sortField = query.Sort[1:]
      sortOrder = -1
    } else {
      sortField = query.Sort
      sortOrder = 1
    }
  }
  
  // Find with options
  opts := options.Find().
    SetSort(bson.D{{Key: sortField, Value: sortOrder}}).
    SetSkip(int64(skip)).
    SetLimit(int64(limit))
  
  cursor, err := r.collection.Find(ctx, filter, opts)
  if err != nil {
    return nil, 0, err
  }
  defer cursor.Close(ctx)
  
  var tasks []models.Task
  if err = cursor.All(ctx, &tasks); err != nil {
    return nil, 0, err
  }
  
  return tasks, total, nil
}

// Update - update task
func (r *taskRepository) Update(ctx context.Context, id bson.ObjectID, userID bson.ObjectID, updates bson.M) error {
  filter := bson.M{
    "_id":     id,
    "user_id": userID,
  }
  
  updates["updated_at"] = time.Now()
  
  update := bson.M{"$set": updates}
  
  result, err := r.collection.UpdateOne(ctx, filter, update)
  if err != nil {
    return err
  }
  
  if result.MatchedCount == 0 {
    return errors.New("task not found")
  }
  
  return nil
}

// Delete - delete task
func (r *taskRepository) Delete(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) error {
  filter := bson.M{
    "_id":     id,
    "user_id": userID,
  }
  
  result, err := r.collection.DeleteOne(ctx, filter)
  if err != nil {
    return err
  }
  
  if result.DeletedCount == 0 {
    return errors.New("task not found")
  }
  
  return nil
}
