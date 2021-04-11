package endpoints

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
	"strings"
)

type localStackResolver struct {
	defaultResolver endpoints.Resolver
}

func getLocalStackUrl() string {
	hostname := os.Getenv("LOCALSTACK_HOSTNAME")

	if hostname == "" {
		hostname = "localhost.localstack.cloud"
	}

	port := os.Getenv("EDGE_PORT")
	if port == "" {
		port = "4566"
	}

	ssl := os.Getenv("USE_SSL")
	disableSSL := true
	if ssl == "1" || strings.ToLower(ssl) == "true" {
		disableSSL = false
	}

	return endpoints.AddScheme(hostname + ":" + port, disableSSL)
}

func (l *localStackResolver) EndpointFor(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
	endpointFor, err := l.defaultResolver.EndpointFor(service, region, opts...)

	if err != nil {
		return endpointFor, err
	}

	endpointFor.URL = getLocalStackUrl()

	return endpointFor, err
}

func NewLocalstackResolver() *localStackResolver {
	resolver := &localStackResolver{
		defaultResolver: endpoints.DefaultResolver(),
	}

	return resolver
}
