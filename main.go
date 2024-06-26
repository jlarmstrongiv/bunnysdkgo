package bunnysdkgo

import (
  auth "github.com/microsoft/kiota-abstractions-go/authentication"
  http "github.com/microsoft/kiota-http-go"

  bunnyApiClient "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client"
  edgeStorageApiClient "github.com/jlarmstrongiv/bunnysdkgo/edge_storage_api_client"
  loggingApiClient "github.com/jlarmstrongiv/bunnysdkgo/logging_api_client"
  streamApiClient "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client"
)

type CreateBunnyApiClientParameters struct {
  AccessKey string
}

func CreateBunnyApiClient(
  options *CreateBunnyApiClientParameters,
) (*bunnyApiClient.BunnyApiClient, error) {
  authProvider, err := auth.NewApiKeyAuthenticationProvider(options.AccessKey, "AccessKey", auth.HEADER_KEYLOCATION)
  if err != nil {
    return nil, err
  }
  httpAdapter, err := http.NewNetHttpRequestAdapter(authProvider)
  if err != nil {
    return nil, err
  }
  return bunnyApiClient.NewBunnyApiClient(httpAdapter), nil
}

type CreateEdgeStorageApiClientParameters struct {
  AccessKey string
}

func CreateEdgeStorageApiClient(
  options *CreateEdgeStorageApiClientParameters,
) (*edgeStorageApiClient.EdgeStorageApiClient, error) {
  authProvider, err := auth.NewApiKeyAuthenticationProvider(options.AccessKey, "AccessKey", auth.HEADER_KEYLOCATION)
  if err != nil {
    return nil, err
  }
  httpAdapter, err := http.NewNetHttpRequestAdapter(authProvider)
  if err != nil {
    return nil, err
  }
  return edgeStorageApiClient.NewEdgeStorageApiClient(httpAdapter), nil
}

type CreateStreamApiClientParameters struct {
  AccessKey string
}

func CreateStreamApiClient(
  options *CreateStreamApiClientParameters,
) (*streamApiClient.StreamApiClient, error) {
  authProvider, err := auth.NewApiKeyAuthenticationProvider(options.AccessKey, "AccessKey", auth.HEADER_KEYLOCATION)
  if err != nil {
    return nil, err
  }
  httpAdapter, err := http.NewNetHttpRequestAdapter(authProvider)
  if err != nil {
    return nil, err
  }
  return streamApiClient.NewStreamApiClient(httpAdapter), nil
}

type CreateLoggingApiClientParameters struct {
  AccessKey string
}

func CreateLoggingApiClient(
  options *CreateStreamApiClientParameters,
) (*loggingApiClient.LoggingApiClient, error) {
  authProvider, err := auth.NewApiKeyAuthenticationProvider(options.AccessKey, "AccessKey", auth.HEADER_KEYLOCATION)
  if err != nil {
    return nil, err
  }
  httpAdapter, err := http.NewNetHttpRequestAdapter(authProvider)
  if err != nil {
    return nil, err
  }
  return loggingApiClient.NewLoggingApiClient(httpAdapter), nil
}
