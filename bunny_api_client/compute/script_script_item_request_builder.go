package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/compute"
)

// ScriptScriptItemRequestBuilder builds and executes requests for operations under \compute\script\{id}
type ScriptScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Code the code property
// returns a *ScriptItemCodeRequestBuilder when successful
func (m *ScriptScriptItemRequestBuilder) Code()(*ScriptItemCodeRequestBuilder) {
    return NewScriptItemCodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewScriptScriptItemRequestBuilderInternal instantiates a new ScriptScriptItemRequestBuilder and sets the default values.
func NewScriptScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptScriptItemRequestBuilder) {
    m := &ScriptScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}", pathParameters),
    }
    return m
}
// NewScriptScriptItemRequestBuilder instantiates a new ScriptScriptItemRequestBuilder and sets the default values.
func NewScriptScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_delete)
func (m *ScriptScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get [GetComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_index2)
// returns a Scriptable when successful
func (m *ScriptScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.Scriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.CreateScriptFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.Scriptable), nil
}
// Post [UpdateComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_update)
// returns a Scriptable when successful
func (m *ScriptScriptItemRequestBuilder) Post(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.ScriptCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.Scriptable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.CreateScriptFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.Scriptable), nil
}
// Publish the publish property
// returns a *ScriptItemPublishRequestBuilder when successful
func (m *ScriptScriptItemRequestBuilder) Publish()(*ScriptItemPublishRequestBuilder) {
    return NewScriptItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Releases the releases property
// returns a *ScriptItemReleasesRequestBuilder when successful
func (m *ScriptScriptItemRequestBuilder) Releases()(*ScriptItemReleasesRequestBuilder) {
    return NewScriptItemReleasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeleteComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_delete)
// returns a *RequestInformation when successful
func (m *ScriptScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation [GetComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_index2)
// returns a *RequestInformation when successful
func (m *ScriptScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_update)
// returns a *RequestInformation when successful
func (m *ScriptScriptItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.ScriptCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Variables the variables property
// returns a *ScriptItemVariablesRequestBuilder when successful
func (m *ScriptScriptItemRequestBuilder) Variables()(*ScriptItemVariablesRequestBuilder) {
    return NewScriptItemVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptScriptItemRequestBuilder when successful
func (m *ScriptScriptItemRequestBuilder) WithUrl(rawUrl string)(*ScriptScriptItemRequestBuilder) {
    return NewScriptScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
