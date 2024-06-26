package country

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i582bea14855a9b569962fbd5880f00ec93190b737da9a06b77f7f301a7754121 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/countries/getcountrylist"
)

// CountryRequestBuilder builds and executes requests for operations under \country
type CountryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewCountryRequestBuilderInternal instantiates a new CountryRequestBuilder and sets the default values.
func NewCountryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CountryRequestBuilder) {
    m := &CountryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/country", pathParameters),
    }
    return m
}
// NewCountryRequestBuilder instantiates a new CountryRequestBuilder and sets the default values.
func NewCountryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CountryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCountryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetCountryList API Docs](https://docs.bunny.net/reference/countriespublic_getcountrylist)
// returns a []Countryable when successful
func (m *CountryRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i582bea14855a9b569962fbd5880f00ec93190b737da9a06b77f7f301a7754121.Countryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i582bea14855a9b569962fbd5880f00ec93190b737da9a06b77f7f301a7754121.CreateCountryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i582bea14855a9b569962fbd5880f00ec93190b737da9a06b77f7f301a7754121.Countryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i582bea14855a9b569962fbd5880f00ec93190b737da9a06b77f7f301a7754121.Countryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [GetCountryList API Docs](https://docs.bunny.net/reference/countriespublic_getcountrylist)
// returns a *RequestInformation when successful
func (m *CountryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CountryRequestBuilder when successful
func (m *CountryRequestBuilder) WithUrl(rawUrl string)(*CountryRequestBuilder) {
    return NewCountryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
