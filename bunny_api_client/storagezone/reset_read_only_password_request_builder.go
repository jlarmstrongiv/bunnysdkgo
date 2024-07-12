package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ResetReadOnlyPasswordRequestBuilder builds and executes requests for operations under \storagezone\resetReadOnlyPassword
type ResetReadOnlyPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResetReadOnlyPasswordRequestBuilderPostQueryParameters [ResetReadOnlyPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetreadonlypassword)
type ResetReadOnlyPasswordRequestBuilderPostQueryParameters struct {
    Id *int64 `uriparametername:"id"`
}
// NewResetReadOnlyPasswordRequestBuilderInternal instantiates a new ResetReadOnlyPasswordRequestBuilder and sets the default values.
func NewResetReadOnlyPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetReadOnlyPasswordRequestBuilder) {
    m := &ResetReadOnlyPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/resetReadOnlyPassword?id={id}", pathParameters),
    }
    return m
}
// NewResetReadOnlyPasswordRequestBuilder instantiates a new ResetReadOnlyPasswordRequestBuilder and sets the default values.
func NewResetReadOnlyPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetReadOnlyPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetReadOnlyPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetReadOnlyPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetreadonlypassword)
func (m *ResetReadOnlyPasswordRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetReadOnlyPasswordRequestBuilderPostQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [ResetReadOnlyPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetreadonlypassword)
// returns a *RequestInformation when successful
func (m *ResetReadOnlyPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetReadOnlyPasswordRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ResetReadOnlyPasswordRequestBuilder when successful
func (m *ResetReadOnlyPasswordRequestBuilder) WithUrl(rawUrl string)(*ResetReadOnlyPasswordRequestBuilder) {
    return NewResetReadOnlyPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
