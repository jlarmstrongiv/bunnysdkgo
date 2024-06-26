package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia9743b6c58ebf0bf73bd3104e3b6f2c76e377208acef9fd8fa96912da9313a67 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/streamvideolibrary/getlanguages"
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
func (m *LanguagesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]ia9743b6c58ebf0bf73bd3104e3b6f2c76e377208acef9fd8fa96912da9313a67.Languageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ia9743b6c58ebf0bf73bd3104e3b6f2c76e377208acef9fd8fa96912da9313a67.CreateLanguageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ia9743b6c58ebf0bf73bd3104e3b6f2c76e377208acef9fd8fa96912da9313a67.Languageable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ia9743b6c58ebf0bf73bd3104e3b6f2c76e377208acef9fd8fa96912da9313a67.Languageable)
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
