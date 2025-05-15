package pulldown

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/kurin/blazer/b2"
	"memb2/pkg/exec"
)

func PulldownToMemory(url string) ([]byte, error) {
	var bucketName, fileName string

	if strings.HasPrefix(url, "b2://") {
		path := strings.TrimPrefix(url, "b2://")
		parts := strings.SplitN(path, "/", 2)
		if len(parts) == 2 {
			bucketName, fileName = parts[0], parts[1]
		} else if len(parts) == 1 {
			bucketName, fileName = exec.DefaultBucket, parts[0]
		} else {
			return nil, fmt.Errorf("invalid B2 path: %s", path)
		}
	} else {
		bucketName, fileName = exec.DefaultBucket, url
	}
	ctx := context.Background()
	client, err := b2.NewClient(ctx, exec.AccountID, exec.ApplicationKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create B2 client: %w", err)
	}
	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket %q: %w", bucketName, err)
	}
	reader := bucket.Object(fileName).NewReader(ctx)
	fileData, err := io.ReadAll(reader)
	if err != nil {
		reader.Close()
		return nil, fmt.Errorf("failed to read file %q: %w", fileName, err)
	}
	reader.Close()

	return fileData, nil
}

func PulldownAndExec(url string, args []string) error {
	var filePath string

	if strings.HasPrefix(url, "b2://") {
		filePath = strings.TrimPrefix(url, "b2://")
	} else {
		filePath = url
	}
	ctx := context.Background()
	return exec.ExecFromB2(ctx, filePath, args)
}