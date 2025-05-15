package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/amenzhinsky/go-memexec"
	"github.com/kurin/blazer/b2"
)


const (
	AccountID       = "0057825b98499560000000003"
	ApplicationKey  = "K005uiYEpOPTirXAEF8zpxwL1uCJh0g"
	DefaultBucket   = "pulldown"
)


func ExecFromB2(ctx context.Context, filePath string, args []string) error {
	var bucketName, fileName string
	
	parts := strings.SplitN(filePath, "/", 2)
	if len(parts) == 2 {
		bucketName, fileName = parts[0], parts[1]
	} else {
		bucketName, fileName = DefaultBucket, filePath
	}

	client, err := b2.NewClient(ctx, AccountID, ApplicationKey)
	if err != nil {
		return fmt.Errorf("failed to create B2 client: %w", err)
	}

	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to get bucket %q: %w", bucketName, err)
	}

	reader := bucket.Object(fileName).NewReader(ctx)
	
	fileData, err := io.ReadAll(reader)
	if err != nil {
		reader.Close()
		return fmt.Errorf("failed to read file %q: %w", fileName, err)
	}
	reader.Close()

	exe, err := memexec.New(fileData)
	if err != nil {
		return fmt.Errorf("failed to create executable from memory: %w", err)
	}
	defer exe.Close()

	cmd := exe.Command(args...)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
