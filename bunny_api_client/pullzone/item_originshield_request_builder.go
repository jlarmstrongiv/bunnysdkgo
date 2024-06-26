package pullzone

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOriginshieldRequestBuilder builds and executes requests for operations under \pullzone\{-id}\originshield
type ItemOriginshieldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemOriginshieldRequestBuilderInternal instantiates a new ItemOriginshieldRequestBuilder and sets the default values.
func NewItemOriginshieldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOriginshieldRequestBuilder) {
    m := &ItemOriginshieldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/originshield", pathParameters),
    }
    return m
}
// NewItemOriginshieldRequestBuilder instantiates a new ItemOriginshieldRequestBuilder and sets the default values.
func NewItemOriginshieldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOriginshieldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOriginshieldRequestBuilderInternal(urlParams, requestAdapter)
}
// Queuestatistics the queuestatistics property
// returns a *ItemOriginshieldQueuestatisticsRequestBuilder when successful
func (m *ItemOriginshieldRequestBuilder) Queuestatistics()(*ItemOriginshieldQueuestatisticsRequestBuilder) {
    return NewItemOriginshieldQueuestatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
