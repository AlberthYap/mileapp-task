package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"task-api/models"
)

// setupTestDB creates an in-memory MongoDB instance for testing
func setupTestDB(t *testing.T) *mongo.Database {
  // Configure memongo with timeout
  opts := &memongo.Options{
    StartupTimeout: 30 * time.Second,
    MongoVersion:   "6.0.4",
  }

  // Start in-memory MongoDB server with options
  mongoServer, err := memongo.StartWithOptions(opts)
  if err != nil {
    t.Skipf("Skipping test: MongoDB not available - %v", err)
    return nil
  }

  // Cleanup after test completes
  t.Cleanup(func() {
    mongoServer.Stop()
  })

  // Connect to in-memory MongoDB
  _, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  client, err := mongo.Connect(options.Client().ApplyURI(mongoServer.URI()))
  if err != nil {
    t.Fatalf("Failed to connect to in-memory MongoDB: %v", err)
  }

  // Return test database
  return client.Database("test_db")
}

func TestUserRepository_FindByEmail(t *testing.T) {
  t.Run("should find user by email", func(t *testing.T) {
    // Setup
    db := setupTestDB(t)
    if db == nil {
      return // Skip if MongoDB not available
    }
    repo := NewUserRepository(db)
    ctx := context.Background()

    // Insert test user
    testUser := models.User{
      ID:        bson.NewObjectID(),
      Email:     "test@test.com",
      Name:      "Test User",
      Password:  "hashedpassword123",
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }

    _, err := db.Collection("users").InsertOne(ctx, testUser)
    assert.NoError(t, err)

    // Execute
    result, err := repo.FindByEmail(ctx, "test@test.com")

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "test@test.com", result.Email)
    assert.Equal(t, "Test User", result.Name)
  })

  t.Run("should return error when user not found", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewUserRepository(db)
    ctx := context.Background()

    result, err := repo.FindByEmail(ctx, "notfound@test.com")

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "user not found", err.Error())
  })

  t.Run("should be case sensitive for email", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewUserRepository(db)
    ctx := context.Background()

    testUser := models.User{
      ID:       bson.NewObjectID(),
      Email:    "Test@Example.com",
      Name:     "Test User",
      Password: "password",
    }
    db.Collection("users").InsertOne(ctx, testUser)

    result, err := repo.FindByEmail(ctx, "test@example.com")

    assert.Error(t, err)
    assert.Nil(t, result)
  })
}
