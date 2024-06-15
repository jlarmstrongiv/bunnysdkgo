package pullzone

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOptimizerRequestBuilder builds and executes requests for operations under \pullzone\{-id}\optimizer
type ItemOptimizerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemOptimizerRequestBuilderInternal instantiates a new ItemOptimizerRequestBuilder and sets the default values.
func NewItemOptimizerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOptimizerRequestBuilder) {
    m := &ItemOptimizerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/optimizer", pathParameters),
    }
    return m
}
// NewItemOptimizerRequestBuilder instantiates a new ItemOptimizerRequestBuilder and sets the default values.
func NewItemOptimizerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOptimizerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOptimizerRequestBuilderInternal(urlParams, requestAdapter)
}
// Statistics the statistics property
// returns a *ItemOptimizerStatisticsRequestBuilder when successful
func (m *ItemOptimizerRequestBuilder) Statistics()(*ItemOptimizerStatisticsRequestBuilder) {
    return NewItemOptimizerStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
