package abusecase

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i41e32aef8b8bf7891740cc7b79701dd14df932502bd55cfb248d7543f74f02fc "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/abusecases"
)

// ItemCheckRequestBuilder builds and executes requests for operations under \abusecase\{id}\check
type ItemCheckRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCheckRequestBuilderInternal instantiates a new ItemCheckRequestBuilder and sets the default values.
func NewItemCheckRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckRequestBuilder) {
    m := &ItemCheckRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/abusecase/{id}/check", pathParameters),
    }
    return m
}
// NewItemCheckRequestBuilder instantiates a new ItemCheckRequestBuilder and sets the default values.
func NewItemCheckRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCheckRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [CheckAbuseCase API Docs](https://docs.bunny.net/reference/abusecasepublic_checkabusecase)
// returns a AbuseCaseable when successful
func (m *ItemCheckRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i41e32aef8b8bf7891740cc7b79701dd14df932502bd55cfb248d7543f74f02fc.AbuseCaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i41e32aef8b8bf7891740cc7b79701dd14df932502bd55cfb248d7543f74f02fc.CreateAbuseCaseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i41e32aef8b8bf7891740cc7b79701dd14df932502bd55cfb248d7543f74f02fc.AbuseCaseable), nil
}
// ToPostRequestInformation [CheckAbuseCase API Docs](https://docs.bunny.net/reference/abusecasepublic_checkabusecase)
// returns a *RequestInformation when successful
func (m *ItemCheckRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCheckRequestBuilder when successful
func (m *ItemCheckRequestBuilder) WithUrl(rawUrl string)(*ItemCheckRequestBuilder) {
    return NewItemCheckRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
