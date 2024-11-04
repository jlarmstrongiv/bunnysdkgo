package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithStorageZoneNameSlashRequestBuilder builds and executes requests for operations under \{storageZoneName}\
type WithStorageZoneNameSlashRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithStorageZoneNameSlashRequestBuilderInternal instantiates a new WithStorageZoneNameSlashRequestBuilder and sets the default values.
func NewWithStorageZoneNameSlashRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, storageZoneName *string)(*WithStorageZoneNameSlashRequestBuilder) {
    m := &WithStorageZoneNameSlashRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{storageZoneName}/", pathParameters),
    }
    if storageZoneName != nil {
        m.BaseRequestBuilder.PathParameters["storageZoneName"] = *storageZoneName
    }
    return m
}
// NewWithStorageZoneNameSlashRequestBuilder instantiates a new WithStorageZoneNameSlashRequestBuilder and sets the default values.
func NewWithStorageZoneNameSlashRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithStorageZoneNameSlashRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithStorageZoneNameSlashRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Post [DownloadZip API Docs](https://toshy.github.io/BunnyNet-PHP/edge-storage-api/#download-zip)
// returns a []byte when successful
func (m *WithStorageZoneNameSlashRequestBuilder) Post(ctx context.Context, body WithStorageZoneNameSlashPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]byte, error) {
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
func (m *WithStorageZoneNameSlashRequestBuilder) ToPostRequestInformation(ctx context.Context, body WithStorageZoneNameSlashPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WithStorageZoneNameSlashRequestBuilder when successful
func (m *WithStorageZoneNameSlashRequestBuilder) WithUrl(rawUrl string)(*WithStorageZoneNameSlashRequestBuilder) {
    return NewWithStorageZoneNameSlashRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
