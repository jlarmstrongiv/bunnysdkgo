package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i9a71e3eaad3e64efd0cf5823f0a3ec5bc2da29761efba49565e409be5c628a63 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/streamvideolibrary/getlanguages"
)

// LanguagesRequestBuilder builds and executes requests for operations under \videolibrary\languages
type LanguagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLanguagesRequestBuilderInternal instantiates a new LanguagesRequestBuilder and sets the default values.
func NewLanguagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LanguagesRequestBuilder) {
    m := &LanguagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/languages", pathParameters),
    }
    return m
}
// NewLanguagesRequestBuilder instantiates a new LanguagesRequestBuilder and sets the default values.
func NewLanguagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LanguagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLanguagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetLanguages API Docs](https://docs.bunny.net/reference/videolibrarypublic_index3)
// returns a []Languageable when successful
func (m *LanguagesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i9a71e3eaad3e64efd0cf5823f0a3ec5bc2da29761efba49565e409be5c628a63.Languageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i9a71e3eaad3e64efd0cf5823f0a3ec5bc2da29761efba49565e409be5c628a63.CreateLanguageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i9a71e3eaad3e64efd0cf5823f0a3ec5bc2da29761efba49565e409be5c628a63.Languageable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i9a71e3eaad3e64efd0cf5823f0a3ec5bc2da29761efba49565e409be5c628a63.Languageable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [GetLanguages API Docs](https://docs.bunny.net/reference/videolibrarypublic_index3)
// returns a *RequestInformation when successful
func (m *LanguagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LanguagesRequestBuilder when successful
func (m *LanguagesRequestBuilder) WithUrl(rawUrl string)(*LanguagesRequestBuilder) {
    return NewLanguagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
