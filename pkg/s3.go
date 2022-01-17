package pkg

import (
	"fmt"
	"net/url"
)

type S3Config struct {
	Region *string
	Bucket *string
	Endpoint *string
	ForcePathStyle *bool
}

func GetS3BaseUrl(conf *S3Config) string {
	var baseUrl string

	if conf.Endpoint == nil || len(*conf.Endpoint) == 0 {
		baseUrl = fmt.Sprintf("https://s3.%s.amazonaws.com", *conf.Region)
	} else {
		baseUrl = *conf.Endpoint
	}

	if *conf.ForcePathStyle {
		baseUrl = fmt.Sprintf("%s/%s", baseUrl, *conf.Bucket)
	} else {
		u, err := url.Parse(baseUrl)

		if err != nil {
			panic(err)
		}

		u.Host = fmt.Sprintf("%s.%s", *conf.Bucket, u.Host)
		baseUrl = u.String()
	}

	return baseUrl
}

func ResolveS3URL(baseUrl string, key string) string {
	return fmt.Sprintf("%s/%s", baseUrl, key)
}