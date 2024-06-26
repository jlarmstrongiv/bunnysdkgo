package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder builds and executes requests for operations under \storagezone\resetReadOnlyPassword
type ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResetreadonlypasswordResetReadOnlyPasswordRequestBuilderPostQueryParameters [ResetReadOnlyPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetreadonlypassword)
type ResetreadonlypasswordResetReadOnlyPasswordRequestBuilderPostQueryParameters struct {
    Id *int64 `uriparametername:"id"`
}
// NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilderInternal instantiates a new ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder and sets the default values.
func NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) {
    m := &ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/resetReadOnlyPassword?id={id}", pathParameters),
    }
    return m
}
// NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilder instantiates a new ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder and sets the default values.
func NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetReadOnlyPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetreadonlypassword)
func (m *ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetreadonlypasswordResetReadOnlyPasswordRequestBuilderPostQueryParameters])(error) {
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
func (m *ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetreadonlypasswordResetReadOnlyPasswordRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder when successful
func (m *ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) WithUrl(rawUrl string)(*ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) {
    return NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
