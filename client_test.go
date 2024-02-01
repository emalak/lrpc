package lrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"testing"
)

func mustJSONString(v any) string {
	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// Make sure to start feed and storage services before testing

func TestGetLandmark(t *testing.T) {
	// Test on correct data
	const (
		landmarkId = "8e20a909-1a4c-43a3-b1d2-2f4505ec19fe"
		userId     = "d00ba81c-4fd9-4e01-a383-97e25f926020"
	)
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}
	landmark, err := client.GetLandmark(ctx, landmarkId, userId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(mustJSONString(landmark))
	// Test on invalid userId
	_, err = client.GetLandmark(ctx, landmarkId, "invalidId")
	fmt.Println(err)
	// Test on invalid landmarkId
	_, err = client.GetLandmark(ctx, "invalidId", userId)
	fmt.Println(err)
	// Test on non-existent landmark
	_, err = client.GetLandmark(ctx, gofakeit.UUID(), userId)
	fmt.Println(err)
	// Test on non-existent user
	_, err = client.GetLandmark(ctx, landmarkId, gofakeit.UUID())
	fmt.Println(err)
}

func TestAddLandmark(t *testing.T) {
	// Set up context and client
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test adding a valid landmark
	validLandmarkId := gofakeit.UUID() // Generate a random UUID for a new landmark
	err = client.AddLandmark(ctx, validLandmarkId)
	if err != nil {
		t.Errorf("Failed to add valid landmark: %v", err)
	}

	// Test adding a landmark with an invalid ID (depending on the format your system expects)
	err = client.AddLandmark(ctx, "invalidId")
	if err == nil {
		t.Error("Expected an error when adding a landmark with an invalid ID, but got none")
	}

	// Test adding a landmark with an empty ID
	err = client.AddLandmark(ctx, "")
	if err == nil {
		t.Error("Expected an error when adding a landmark with an empty ID, but got none")
	}
}

func TestLikeLandmark(t *testing.T) {
	// Setup for tests
	const (
		validUserId     = "0cde699d-a1f2-4f3f-9fa3-0544226c678f" // replace with an actual valid ID
		validLandmarkId = "baa83618-159c-4015-bb22-e97f95190e25" // replace with an actual valid ID
	)
	ctx := context.Background()

	// Initialize client
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test liking a landmark with correct data
	err = client.LikeLandmark(ctx, validUserId, validLandmarkId)
	if err != nil {
		t.Errorf("Failed to like landmark with valid IDs, error: %v", err)
	}

	// Test liking a landmark with an invalid userId
	err = client.LikeLandmark(ctx, "invalidUserId", validLandmarkId)
	if err == nil {
		t.Error("Expected an error when liking a landmark with an invalid userId, but got nil")
	}

	// Test liking a landmark with an invalid landmarkId
	err = client.LikeLandmark(ctx, validUserId, "invalidLandmarkId")
	if err == nil {
		t.Error("Expected an error when liking a landmark with an invalid landmarkId, but got nil")
	}

	// Test liking a non-existent landmark
	err = client.LikeLandmark(ctx, validUserId, gofakeit.UUID())
	if err == nil {
		t.Error("Expected an error when liking a non-existent landmark, but got nil")
	}

	// Test liking a landmark with a non-existent user
	err = client.LikeLandmark(ctx, gofakeit.UUID(), validLandmarkId)
	if err == nil {
		t.Error("Expected an error when liking a landmark with a non-existent user, but got nil")
	}
}

func TestDislikeLandmark(t *testing.T) {
	// Test on correct data
	const (
		landmarkId = "8e20a909-1a4c-43a3-b1d2-2f4505ec19fe"
		userId     = "d00ba81c-4fd9-4e01-a383-97e25f926020"
	)
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test disliking a landmark with correct userId and landmarkId
	err = client.DislikeLandmark(ctx, userId, landmarkId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test disliking a landmark with an invalid userId
	err = client.DislikeLandmark(ctx, "invalidUserId", landmarkId)
	if err == nil {
		t.Error("Expected an error for invalid userId, got none")
	}

	// Test disliking a landmark with an invalid landmarkId
	err = client.DislikeLandmark(ctx, userId, "invalidLandmarkId")
	if err == nil {
		t.Error("Expected an error for invalid landmarkId, got none")
	}

	// Test disliking a non-existent landmark
	err = client.DislikeLandmark(ctx, userId, gofakeit.UUID())
	if err == nil {
		t.Error("Expected an error for non-existent landmark, got none")
	}

	// Test disliking a landmark by a non-existent user
	err = client.DislikeLandmark(ctx, gofakeit.UUID(), landmarkId)
	if err == nil {
		t.Error("Expected an error for non-existent user, got none")
	}
}

func TestGetLikes(t *testing.T) {
	// Define a context for the tests
	ctx := context.Background()

	// Initialize the client with test settings
	client, err := New(ctx, Settings{
		// Assuming FeedOpts and StorageOpts are required for the client to function
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err) // If the client initialization fails, stop the test
	}

	// Test case with a valid landmark ID
	const validLandmarkId = "8e20a909-1a4c-43a3-b1d2-2f4505ec19fe"
	likes, err := client.GetLikes(ctx, validLandmarkId)
	if err != nil {
		t.Errorf("Expected no error for valid landmark ID, got %v", err)
	}
	fmt.Printf("Likes for valid landmark: %d\n", likes)

	// Test case with an invalid landmark ID
	_, err = client.GetLikes(ctx, "invalidLandmarkId")
	if err == nil {
		t.Error("Expected an error for invalid landmark ID, got none")
	}
	fmt.Println("Error for invalid landmark ID:", err)

	// Test case with a non-existent landmark ID (using gofakeit or similar to generate a UUID)
	_, err = client.GetLikes(ctx, gofakeit.UUID())
	if err == nil {
		t.Error("Expected an error for non-existent landmark ID, got none")
	}
	fmt.Println("Error for non-existent landmark ID:", err)
}

func TestViewLandmark(t *testing.T) {
	// Test on correct data
	const (
		landmarkId = "8e20a909-1a4c-43a3-b1d2-2f4505ec19fe"
		userId     = "d00ba81c-4fd9-4e01-a383-97e25f926020"
	)
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = client.ViewLandmark(ctx, userId, landmarkId)
	if err != nil {
		t.Fatalf("Expected no error for valid user and landmark IDs, got %v", err)
	}

	// Test on invalid userId
	err = client.ViewLandmark(ctx, "invalidId", landmarkId)
	if err == nil {
		t.Fatal("Expected an error for invalid userId, got nil")
	}

	// Test on invalid landmarkId
	err = client.ViewLandmark(ctx, userId, "invalidId")
	if err == nil {
		t.Fatal("Expected an error for invalid landmarkId, got nil")
	}

	// Test on non-existent landmark
	err = client.ViewLandmark(ctx, userId, gofakeit.UUID())
	if err == nil {
		t.Fatal("Expected an error for non-existent landmark, got nil")
	}

	// Test on non-existent user
	err = client.ViewLandmark(ctx, gofakeit.UUID(), landmarkId)
	if err == nil {
		t.Fatal("Expected an error for non-existent user, got nil")
	}
}

func TestRecommendLandmarks(t *testing.T) {
	const (
		userId = "d00ba81c-4fd9-4e01-a383-97e25f926020" // Assuming a valid user ID
		amount = 5                                      // Example amount of landmarks to recommend
	)
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test with correct data
	landmarks, err := client.RecommendLandmarks(ctx, userId, amount)
	if err != nil {
		t.Errorf("Unexpected error for RecommendLandmarks with valid data: %v", err)
	}
	if len(landmarks) > amount {
		t.Errorf("Expected at most %d landmarks, got %d", amount, len(landmarks))
	}
	fmt.Printf("Recommended landmarks: %v\n", landmarks)

	// Test with invalid userId
	_, err = client.RecommendLandmarks(ctx, "invalidId", amount)
	if err == nil {
		t.Error("Expected error for RecommendLandmarks with invalid userId, got nil")
	} else {
		fmt.Println("Expected error received:", err)
	}
}

func TestGetRandomFeed(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test with a valid amount
	amount := 5 // Assuming 5 is a reasonable amount for a random feed
	feeds, err := client.GetRandomFeed(ctx, amount)
	if err != nil {
		t.Fatal(err)
	}
	if len(feeds) != amount {
		t.Fatalf("Expected %d feeds, got %d", amount, len(feeds))
	}
	fmt.Println("Successfully retrieved random feed:", mustJSONString(feeds))

	// Test with a negative amount
	_, err = client.GetRandomFeed(ctx, -1)
	if err == nil {
		t.Fatal("Expected an error for negative amount, but got none")
	}
	fmt.Println("Correctly received error for negative amount:", err)

	// Test with amount of 0
	feeds, err = client.GetRandomFeed(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(feeds) != 0 {
		t.Fatalf("Expected 0 feeds, got %d", len(feeds))
	}
	fmt.Println("Successfully handled request for 0 feeds")

	// Test with an excessively large amount
	largeAmount := 10000 // Assuming 10000 is an excessively large amount
	feeds, err = client.GetRandomFeed(ctx, largeAmount)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddUser(t *testing.T) {
	// Context setup
	ctx := context.Background()

	// Create client with test settings
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test with a valid userId
	validUserId := gofakeit.UUID() // Consider using a UUID generation function for dynamic testing
	err = client.AddUser(ctx, validUserId)
	if err != nil {
		t.Fatalf("Failed to add a valid user: %v", err)
	}

	// Optionally, retrieve the user to verify addition, if such a method exists in the client
	// _, err = client.GetUser(ctx, validUserId)
	// if err != nil {
	//     t.Fatalf("Failed to retrieve the added user: %v", err)
	// }

	// Test with an invalid userId
	err = client.AddUser(ctx, "invalidUserId")
	if err == nil {
		t.Error("Expected an error when adding a user with an invalid ID, but did not get one")
	}

	// Test adding a user that already exists to see how the system handles duplicates
	err = client.AddUser(ctx, validUserId)
	if err == nil {
		// Depending on the system's design, this might be okay (idempotency) or might be an error (unique constraint)
		// Adjust the test based on expected behavior
		t.Error("Expected an error when re-adding an existing user, but did not get one")
	}
}

func TestCreateComment(t *testing.T) {
	// Set up context and client
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test data
	const (
		parentId = "712d6060-284b-400b-805d-cc118cd41c5d"
		authorId = "30de160d-e48c-4beb-94c3-b5c6eb8075e6"
	)
	text := gofakeit.HipsterSentence(10)
	validAttachments := []string{gofakeit.UUID(), gofakeit.UUID()}
	invalidAttachments := []string{"invalidUUID", "anotherInvalidUUID"}

	// Test on correct data
	err = client.CreateComment(ctx, parentId, authorId, text, validAttachments, 3)
	if err != nil {
		t.Fatal(err)
	}

	// Test on invalid parentId
	err = client.CreateComment(ctx, "invalidParentId", authorId, text, validAttachments, 3)
	if err == nil {
		t.Error("Expected error for invalid parentId, got nil")
	}

	// Test on invalid authorId
	err = client.CreateComment(ctx, parentId, "invalidAuthorId", text, validAttachments, 3)
	if err == nil {
		t.Error("Expected error for invalid authorId, got nil")
	}

	// Test on invalid attachments
	err = client.CreateComment(ctx, parentId, authorId, text, invalidAttachments, 3)
	if err == nil {
		t.Error("Expected error for invalid attachments, got nil")
	}

	// Test on non-existent parentId
	err = client.CreateComment(ctx, gofakeit.UUID(), authorId, text, validAttachments, 3)
	if err == nil {
		t.Error("Expected error for non-existent parentId, got nil")
	}

	// Test on non-existent authorId
	err = client.CreateComment(ctx, parentId, gofakeit.UUID(), text, validAttachments, 3)
	if err == nil {
		t.Error("Expected error for non-existent authorId, got nil")
	}
}

func TestGetComments(t *testing.T) {
	// Test on correct data
	const (
		landmarkId = "712d6060-284b-400b-805d-cc118cd41c5d"
	)
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	comments, err := client.GetComments(ctx, landmarkId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(mustJSONString(comments))

	// Test on invalid landmarkId
	_, err = client.GetComments(ctx, "invalidId")
	if err == nil {
		t.Error("Expected an error for invalid landmark ID, got none")
	}
	fmt.Println(err)

	// Test on non-existent landmark
	_, err = client.GetComments(ctx, gofakeit.UUID())
	if err == nil {
		t.Error("Expected an error for non-existent landmark ID, got none")
	}
	fmt.Println(err)
}

func TestGetProfileComments(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name      string
		userId    string
		limit     int
		expectErr bool
	}{
		{
			name:      "Valid user and limit",
			userId:    "2e9ee190-217c-4f92-aea4-b8086501fbb2", // replace with an actual valid userId
			limit:     10,
			expectErr: false,
		},
		{
			name:      "Invalid userId",
			userId:    "invalidUserId",
			limit:     10,
			expectErr: true,
		},
		{
			name:      "Zero limit",
			userId:    "2e9ee190-217c-4f92-aea4-b8086501fbb2",
			limit:     0,
			expectErr: true, // Assuming fetching with limit 0 is considered an error
		},
		{
			name:      "Negative limit",
			userId:    "validUserId", // replace with the same actual valid userId
			limit:     -1,
			expectErr: true, // Assuming fetching with a negative limit is considered an error
		},
	}

	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			comments, err := client.GetProfileComments(ctx, tc.userId, tc.limit)
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got one: %v", err)
				}
				if len(comments) == 0 {
					t.Errorf("expected comments but got none")
				}
				// Further validations on the comments could be done here if necessary
			}
		})
	}
}

func TestGetFeed(t *testing.T) {
	// Setup context and client
	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    &FeedOptions{Address: "localhost:8081"},
		StorageOpts: nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test variables
	const (
		validUserId = "2e9ee190-217c-4f92-aea4-b8086501fbb2" // Example valid user ID
		validAmount = 10                                     // Example valid amount
	)

	// Test on correct data
	feed, err := client.GetFeed(ctx, validUserId, validAmount)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(feed)

	// Test on invalid userId
	_, err = client.GetFeed(ctx, "invalidId", validAmount)
	if err == nil {
		t.Fatal("Expected an error for invalid userId, got nil")
	}
	fmt.Println(err)

	// Test on negative amount
	_, err = client.GetFeed(ctx, validUserId, -1)
	if err == nil {
		t.Fatal("Expected an error for negative amount, got nil")
	}
	fmt.Println(err)

	// Test on zero amount
	_, err = client.GetFeed(ctx, validUserId, 0)
	if err != nil {
		t.Fatal(err)
	}

	// Test on excessively large amount
	// Assuming there's a maximum limit, adjust the value accordingly
	_, err = client.GetFeed(ctx, validUserId, 1000)
	if err == nil {
		t.Fatal("Expected an error for excessively large amount, got nil")
	}
	fmt.Println(err)

}

func TestGetFavouriteLandmarks(t *testing.T) {
	const (
		validUserId   = "2e9ee190-217c-4f92-aea4-b8086501fbb2"
		invalidUserId = "invalidId"
	)

	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	ids, err := client.GetFavouriteLandmarks(ctx, validUserId)
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) < 1 {
		t.Fatal("empty ids")
	}
	fmt.Println(ids)

	_, err = client.GetFavouriteLandmarks(ctx, invalidUserId)
	if err == nil {
		t.Fatal("Expected an error for invalid id, got nil")
	}
}

func TestGetLikesAmount(t *testing.T) {
	const (
		validUserId   = "2e9ee190-217c-4f92-aea4-b8086501fbb2"
		invalidUserId = "invalidId"
	)

	ctx := context.Background()
	client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.GetFavouriteLandmarks(ctx, validUserId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	_, err = client.GetFavouriteLandmarks(ctx, invalidUserId)
	if err == nil {
		t.Fatal("Expected an error for invalid id, got nil")
	}
}
