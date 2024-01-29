# Landmark API client

## Usage

Creating a client

```
ctx := context.Background()

// To use include service methods, specify its settings
client, err := New(ctx, Settings{
		FeedOpts:    nil,
		StorageOpts: &StorageOptions{Address: "localhost:8080"},
	})
```

Calling a method

```
landmark, err := client.GetLandmark(ctx, landmarkId, userId)
```
