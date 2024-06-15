package bunnysdkgo

import (
    auth "github.com/microsoft/kiota-abstractions-go/authentication"
    http "github.com/microsoft/kiota-http-go"

    bunnyApiClient "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient"
    edgeStorageApiClient "github.com/jlarmstrongiv/bunnysdkgo/generated/EdgeStorageApiClient"
    streamApiClient "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient"
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
