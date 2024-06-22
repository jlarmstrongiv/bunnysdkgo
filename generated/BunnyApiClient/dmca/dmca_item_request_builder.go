package dmca

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/abusecases"
)

// DmcaItemRequestBuilder builds and executes requests for operations under \dmca\{id}
type DmcaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewDmcaItemRequestBuilderInternal instantiates a new DmcaItemRequestBuilder and sets the default values.
func NewDmcaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DmcaItemRequestBuilder) {
    m := &DmcaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dmca/{id}", pathParameters),
    }
    return m
}
// NewDmcaItemRequestBuilder instantiates a new DmcaItemRequestBuilder and sets the default values.
func NewDmcaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DmcaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDmcaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetDmcaCase API Docs](https://docs.bunny.net/reference/abusecasepublic_getabusecase)
// returns a AbuseCaseable when successful
func (m *DmcaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(icdffe60cae25ff52fa12b923bc6a7419c592a8234065e261cea39e4662e1e250.AbuseCaseable, error) {
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
func (m *DmcaItemRequestBuilder) Resolve()(*ItemResolveRequestBuilder) {
    return NewItemResolveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation [GetDmcaCase API Docs](https://docs.bunny.net/reference/abusecasepublic_getabusecase)
// returns a *RequestInformation when successful
func (m *DmcaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DmcaItemRequestBuilder when successful
func (m *DmcaItemRequestBuilder) WithUrl(rawUrl string)(*DmcaItemRequestBuilder) {
    return NewDmcaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
