package region

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iff4e327dd37daa2c2c511990085fcd1500c6b7fa22aa5c0ba45b4f7e7ce570b6 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/region/regionlist"
)

// RegionRequestBuilder builds and executes requests for operations under \region
type RegionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewRegionRequestBuilderInternal instantiates a new RegionRequestBuilder and sets the default values.
func NewRegionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegionRequestBuilder) {
    m := &RegionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/region", pathParameters),
    }
    return m
}
// NewRegionRequestBuilder instantiates a new RegionRequestBuilder and sets the default values.
func NewRegionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [RegionList API Docs](https://docs.bunny.net/reference/regionpublic_index)
// returns a []Regionable when successful
func (m *RegionRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]iff4e327dd37daa2c2c511990085fcd1500c6b7fa22aa5c0ba45b4f7e7ce570b6.Regionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, iff4e327dd37daa2c2c511990085fcd1500c6b7fa22aa5c0ba45b4f7e7ce570b6.CreateRegionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]iff4e327dd37daa2c2c511990085fcd1500c6b7fa22aa5c0ba45b4f7e7ce570b6.Regionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(iff4e327dd37daa2c2c511990085fcd1500c6b7fa22aa5c0ba45b4f7e7ce570b6.Regionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [RegionList API Docs](https://docs.bunny.net/reference/regionpublic_index)
// returns a *RequestInformation when successful
func (m *RegionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RegionRequestBuilder when successful
func (m *RegionRequestBuilder) WithUrl(rawUrl string)(*RegionRequestBuilder) {
    return NewRegionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
