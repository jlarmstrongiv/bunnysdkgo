package pullzone

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSafehopRequestBuilder builds and executes requests for operations under \pullzone\{-id}\safehop
type ItemSafehopRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSafehopRequestBuilderInternal instantiates a new ItemSafehopRequestBuilder and sets the default values.
func NewItemSafehopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSafehopRequestBuilder) {
    m := &ItemSafehopRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/safehop", pathParameters),
    }
    return m
}
// NewItemSafehopRequestBuilder instantiates a new ItemSafehopRequestBuilder and sets the default values.
func NewItemSafehopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSafehopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSafehopRequestBuilderInternal(urlParams, requestAdapter)
}
// Statistics the statistics property
// returns a *ItemSafehopStatisticsRequestBuilder when successful
func (m *ItemSafehopRequestBuilder) Statistics()(*ItemSafehopStatisticsRequestBuilder) {
    return NewItemSafehopStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
