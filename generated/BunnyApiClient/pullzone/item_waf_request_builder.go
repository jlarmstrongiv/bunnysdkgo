package pullzone

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWafRequestBuilder builds and executes requests for operations under \pullzone\{-id}\waf
type ItemWafRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemWafRequestBuilderInternal instantiates a new ItemWafRequestBuilder and sets the default values.
func NewItemWafRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWafRequestBuilder) {
    m := &ItemWafRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/waf", pathParameters),
    }
    return m
}
// NewItemWafRequestBuilder instantiates a new ItemWafRequestBuilder and sets the default values.
func NewItemWafRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWafRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWafRequestBuilderInternal(urlParams, requestAdapter)
}
// Statistics the statistics property
// returns a *ItemWafStatisticsRequestBuilder when successful
func (m *ItemWafRequestBuilder) Statistics()(*ItemWafStatisticsRequestBuilder) {
    return NewItemWafStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
