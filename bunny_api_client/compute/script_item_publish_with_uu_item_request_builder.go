package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScriptItemPublishWithUuItemRequestBuilder builds and executes requests for operations under \compute\script\{id}\publish\{uuid}
type ScriptItemPublishWithUuItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewScriptItemPublishWithUuItemRequestBuilderInternal instantiates a new ScriptItemPublishWithUuItemRequestBuilder and sets the default values.
func NewScriptItemPublishWithUuItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemPublishWithUuItemRequestBuilder) {
    m := &ScriptItemPublishWithUuItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/publish/{uuid}", pathParameters),
    }
    return m
}
// NewScriptItemPublishWithUuItemRequestBuilder instantiates a new ScriptItemPublishWithUuItemRequestBuilder and sets the default values.
func NewScriptItemPublishWithUuItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemPublishWithUuItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemPublishWithUuItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [PublishComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_publish)
func (m *ScriptItemPublishWithUuItemRequestBuilder) Post(ctx context.Context, body ScriptItemPublishItemWithUuPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [PublishComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_publish)
// returns a *RequestInformation when successful
func (m *ScriptItemPublishWithUuItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ScriptItemPublishItemWithUuPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptItemPublishWithUuItemRequestBuilder when successful
func (m *ScriptItemPublishWithUuItemRequestBuilder) WithUrl(rawUrl string)(*ScriptItemPublishWithUuItemRequestBuilder) {
    return NewScriptItemPublishWithUuItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
