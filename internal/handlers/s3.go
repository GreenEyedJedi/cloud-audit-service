package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Handler struct {
	Client *s3.Client
}

func (h *S3Handler) ListBuckets(writer http.ResponseWriter, request *http.Request) {
	out, err := h.Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		http.Error(writer, fmt.Sprintf("Failed to list buckets: %v", err), http.StatusInternalServerError)
		return
	}

	for _, bucket := range out.Buckets {
		fmt.Fprintf(writer, "Bucket: %s\n", *bucket.Name)
	}
}
