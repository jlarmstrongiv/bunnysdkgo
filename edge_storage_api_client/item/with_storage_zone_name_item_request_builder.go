package item

import (
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
