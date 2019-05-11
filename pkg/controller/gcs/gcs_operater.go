/*
Copyright 2019 Hiroki Matsumoto.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcs

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// GenerateBacketFullName generate backet fullname for uniqeue name.
func GenerateBacketFullName(name string) string {
	now := time.Now().Format("20060102")
	s1 := make([]rune, 8)
	for i, w := range now {
		j, _ := strconv.ParseInt(string(w), 10, 64)
		s1[i] = letterRunes[j]
	}

	s2 := make([]rune, 6)
	for i := range s2 {
		s2[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return fmt.Sprintf("%s-%s-%s", name, string(s1), string(s2))
}

// NewGcsClient create storage client.
func NewGcsClient(ctx context.Context, credential []byte) (*storage.Client, error) {
	return storage.NewClient(ctx, option.WithCredentialsJSON(credential))
}

// CreateBucket create bucket.
func CreateBucket(ctx context.Context, client *storage.Client, projectID, bucketName string) error {
	bkt := client.Bucket(bucketName)
	return bkt.Create(ctx, projectID, nil)
}

// IfExistsBucket return backet exists.
func IfExistsBucket(ctx context.Context, client *storage.Client, projectID, bucketName string) bool {
	bkt := client.Bucket(bucketName)
	_, err := bkt.Attrs(ctx)
	return (err == nil)
}
