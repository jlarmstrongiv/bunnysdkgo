package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithStorageZoneNameItemRequestBuilder builds and executes requests for operations under \{storageZoneName}
type WithStorageZoneNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPath gets an item from the github.com/jlarmstrongiv/bunnysdkgo/edge_storage_api_client.item.item collection
// returns a *WithPathItemRequestBuilder when successful
func (m *WithStorageZoneNameItemRequestBuilder) ByPath(path string)(*WithPathItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if path != "" {
        urlTplParams["path"] = path
    }
    return NewWithPathItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithStorageZoneNameItemRequestBuilderInternal instantiates a new WithStorageZoneNameItemRequestBuilder and sets the default values.
func NewWithStorageZoneNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithStorageZoneNameItemRequestBuilder) {
    m := &WithStorageZoneNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{storageZoneName}", pathParameters),
    }
    return m
}
// NewWithStorageZoneNameItemRequestBuilder instantiates a new WithStorageZoneNameItemRequestBuilder and sets the default values.
func NewWithStorageZoneNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithStorageZoneNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithStorageZoneNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [DownloadZip API Docs](https://toshy.github.io/BunnyNet-PHP/edge-storage-api/#download-zip)
// returns a []byte when successful
func (m *WithStorageZoneNameItemRequestBuilder) Post(ctx context.Context, body WithStorageZoneNamePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation [DownloadZip API Docs](https://toshy.github.io/BunnyNet-PHP/edge-storage-api/#download-zip)
// returns a *RequestInformation when successful
func (m *WithStorageZoneNameItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body WithStorageZoneNamePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithStorageZoneNameItemRequestBuilder when successful
func (m *WithStorageZoneNameItemRequestBuilder) WithUrl(rawUrl string)(*WithStorageZoneNameItemRequestBuilder) {
    return NewWithStorageZoneNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
