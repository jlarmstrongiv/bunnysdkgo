package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithPathItemRequestBuilder builds and executes requests for operations under \{storageZoneName}\{+path}
type WithPathItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByFileName gets an item from the github.com/jlarmstrongiv/bunnysdkgo/edge_storage_api_client.item.item.item collection
// returns a *ItemWithFileNameItemRequestBuilder when successful
func (m *WithPathItemRequestBuilder) ByFileName(fileName string)(*ItemWithFileNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fileName != "" {
        urlTplParams["fileName"] = fileName
    }
    return NewItemWithFileNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithPathItemRequestBuilderInternal instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewWithPathItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPathItemRequestBuilder) {
    m := &WithPathItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{storageZoneName}/{+path}", pathParameters),
    }
    return m
}
// NewWithPathItemRequestBuilder instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewWithPathItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPathItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithPathItemRequestBuilderInternal(urlParams, requestAdapter)
}
