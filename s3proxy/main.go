package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/samber/lo"
)

func getEnv(key string, fallback ...string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		if len(fallback) > 0 {
			return fallback[0]
		}
		panic(fmt.Errorf("var not found: %s", key))
	}
	return value
}

type Proxy struct {
	S3Client *s3.Client
	Bucket   string
}

func (p Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	obj, err := p.S3Client.GetObject(r.Context(), &s3.GetObjectInput{
		Bucket: &p.Bucket,
		Key:    &r.URL.Path,
	})
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		panic(err)
	}
	defer obj.Body.Close()

	w.Header().Add("Content-Length", fmt.Sprintf("%d", *obj.ContentLength))
	io.Copy(w, obj.Body)
}

func main() {
	bucket := getEnv("BUCKET")
	endpoint := getEnv("ENDPOINT")
	forcePathStype := getEnv("FORCE_PATH_STYLE") != ""

	cfg := lo.Must(config.LoadDefaultConfig(context.Background()))

	if endpoint != "" {
		cfg.BaseEndpoint = &endpoint
	}

	proxy := Proxy{
		Bucket: bucket,
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = forcePathStype
		}),
	}
	http.HandleFunc("/*", proxy.Handle)

	log.Fatal(http.ListenAndServe(":9292", nil))
}
