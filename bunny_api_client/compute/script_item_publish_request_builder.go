package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScriptItemPublishRequestBuilder builds and executes requests for operations under \compute\script\{id}\publish
type ScriptItemPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ScriptItemPublishRequestBuilderPostQueryParameters [PublishComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_publish)
type ScriptItemPublishRequestBuilderPostQueryParameters struct {
    // The UUID of the script release that will be published
    Uuid *string `uriparametername:"uuid"`
}
// ByUuid gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.compute.script.item.publish.item collection
// returns a *ScriptItemPublishWithUuItemRequestBuilder when successful
func (m *ScriptItemPublishRequestBuilder) ByUuid(uuid string)(*ScriptItemPublishWithUuItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if uuid != "" {
        urlTplParams["uuid"] = uuid
    }
    return NewScriptItemPublishWithUuItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewScriptItemPublishRequestBuilderInternal instantiates a new ScriptItemPublishRequestBuilder and sets the default values.
func NewScriptItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemPublishRequestBuilder) {
    m := &ScriptItemPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/publish?uuid={uuid}", pathParameters),
    }
    return m
}
// NewScriptItemPublishRequestBuilder instantiates a new ScriptItemPublishRequestBuilder and sets the default values.
func NewScriptItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [PublishComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_publish)
func (m *ScriptItemPublishRequestBuilder) Post(ctx context.Context, body ScriptItemPublishPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptItemPublishRequestBuilderPostQueryParameters])(error) {
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
func (m *ScriptItemPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, body ScriptItemPublishPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptItemPublishRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptItemPublishRequestBuilder when successful
func (m *ScriptItemPublishRequestBuilder) WithUrl(rawUrl string)(*ScriptItemPublishRequestBuilder) {
    return NewScriptItemPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
