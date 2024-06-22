package abusecase

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AbusecaseRequestBuilder builds and executes requests for operations under \abusecase
type AbusecaseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AbusecaseRequestBuilderGetQueryParameters [ListAbuseCases API Docs](https://docs.bunny.net/reference/abusecasepublic_index)
type AbusecaseRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.abusecase.item collection
// returns a *AbusecaseItemRequestBuilder when successful
func (m *AbusecaseRequestBuilder) ById(id int64)(*AbusecaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewAbusecaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAbusecaseRequestBuilderInternal instantiates a new AbusecaseRequestBuilder and sets the default values.
func NewAbusecaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AbusecaseRequestBuilder) {
    m := &AbusecaseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/abusecase?page={page}&perPage={perPage}", pathParameters),
    }
    return m
}
// NewAbusecaseRequestBuilder instantiates a new AbusecaseRequestBuilder and sets the default values.
func NewAbusecaseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AbusecaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAbusecaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListAbuseCases API Docs](https://docs.bunny.net/reference/abusecasepublic_index)
// returns a AbusecaseGetResponseable when successful
func (m *AbusecaseRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[AbusecaseRequestBuilderGetQueryParameters])(AbusecaseGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAbusecaseGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AbusecaseGetResponseable), nil
}
// ToGetRequestInformation [ListAbuseCases API Docs](https://docs.bunny.net/reference/abusecasepublic_index)
// returns a *RequestInformation when successful
func (m *AbusecaseRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[AbusecaseRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AbusecaseRequestBuilder when successful
func (m *AbusecaseRequestBuilder) WithUrl(rawUrl string)(*AbusecaseRequestBuilder) {
    return NewAbusecaseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
