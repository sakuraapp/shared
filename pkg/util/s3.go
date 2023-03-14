package util

import (
	"fmt"
	"gopkg.in/guregu/null.v4"
	"net/url"
)

type S3Config struct {
	Region         null.String
	Bucket         string
	Endpoint       null.String
	ForcePathStyle bool
}

func GetS3BaseUrl(conf *S3Config) string {
	if conf.Region.IsZero() && conf.Endpoint.IsZero() {
		panic("Invalid input")
	}

	var baseUrl string

	if len(conf.Endpoint.ValueOrZero()) == 0 {
		baseUrl = fmt.Sprintf("https://s3.%s.amazonaws.com", conf.Region.ValueOrZero())
	} else {
		baseUrl = conf.Endpoint.String
	}

	if conf.ForcePathStyle {
		baseUrl = fmt.Sprintf("%s/%s", baseUrl, conf.Bucket)
	} else {
		u, err := url.Parse(baseUrl)

		if err != nil {
			panic(err)
		}

		u.Host = fmt.Sprintf("%s.%s", conf.Bucket, u.Host)
		baseUrl = u.String()
	}

	return baseUrl
}

func ResolveS3URL(baseUrl string, key string) string {
	return fmt.Sprintf("%s/%s", baseUrl, key)
}
