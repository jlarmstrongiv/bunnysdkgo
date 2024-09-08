package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7a32171f58a3aa231d71f156ea2e966263a8d50ee4200a6cf87edd492456911b "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/search/globalsearch"
)

// SearchRequestBuilder builds and executes requests for operations under \search
type SearchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SearchRequestBuilderGetQueryParameters [GlobalSearch API Docs](https://docs.bunny.net/reference/searchpublic_globalsearch)
type SearchRequestBuilderGetQueryParameters struct {
    // The number of results skipped in the search query
    From *int32 `uriparametername:"from"`
    // The input query for the search request
    Search *string `uriparametername:"search"`
    // The size of the result set
    Size *int32 `uriparametername:"size"`
}
// NewSearchRequestBuilderInternal instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    m := &SearchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search?search={search}{&from,size}", pathParameters),
    }
    return m
}
// NewSearchRequestBuilder instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GlobalSearch API Docs](https://docs.bunny.net/reference/searchpublic_globalsearch)
// returns a SearchResultsable when successful
func (m *SearchRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[SearchRequestBuilderGetQueryParameters])(i7a32171f58a3aa231d71f156ea2e966263a8d50ee4200a6cf87edd492456911b.SearchResultsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i7a32171f58a3aa231d71f156ea2e966263a8d50ee4200a6cf87edd492456911b.CreateSearchResultsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i7a32171f58a3aa231d71f156ea2e966263a8d50ee4200a6cf87edd492456911b.SearchResultsable), nil
}
// ToGetRequestInformation [GlobalSearch API Docs](https://docs.bunny.net/reference/searchpublic_globalsearch)
// returns a *RequestInformation when successful
func (m *SearchRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[SearchRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SearchRequestBuilder when successful
func (m *SearchRequestBuilder) WithUrl(rawUrl string)(*SearchRequestBuilder) {
    return NewSearchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
