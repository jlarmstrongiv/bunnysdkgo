package abusecase

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/abusecases"
)

// AbusecaseItemRequestBuilder builds and executes requests for operations under \abusecase\{id}
type AbusecaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Check the check property
// returns a *ItemCheckRequestBuilder when successful
func (m *AbusecaseItemRequestBuilder) Check()(*ItemCheckRequestBuilder) {
    return NewItemCheckRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAbusecaseItemRequestBuilderInternal instantiates a new AbusecaseItemRequestBuilder and sets the default values.
func NewAbusecaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AbusecaseItemRequestBuilder) {
    m := &AbusecaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/abusecase/{id}", pathParameters),
    }
    return m
}
// NewAbusecaseItemRequestBuilder instantiates a new AbusecaseItemRequestBuilder and sets the default values.
func NewAbusecaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AbusecaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAbusecaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetAbuseCase API Docs](https://docs.bunny.net/reference/abusecasepublic_getabusecase2)
// returns a AbuseCaseable when successful
func (m *AbusecaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250.AbuseCaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250.CreateAbuseCaseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250.AbuseCaseable), nil
}
// Resolve the resolve property
// returns a *ItemResolveRequestBuilder when successful
func (m *AbusecaseItemRequestBuilder) Resolve()(*ItemResolveRequestBuilder) {
    return NewItemResolveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation [GetAbuseCase API Docs](https://docs.bunny.net/reference/abusecasepublic_getabusecase2)
// returns a *RequestInformation when successful
func (m *AbusecaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AbusecaseItemRequestBuilder when successful
func (m *AbusecaseItemRequestBuilder) WithUrl(rawUrl string)(*AbusecaseItemRequestBuilder) {
    return NewAbusecaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
