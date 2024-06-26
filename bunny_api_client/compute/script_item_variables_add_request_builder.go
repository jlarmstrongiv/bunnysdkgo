package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/compute"
)

// ScriptItemVariablesAddRequestBuilder builds and executes requests for operations under \compute\script\{id}\variables\add
type ScriptItemVariablesAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewScriptItemVariablesAddRequestBuilderInternal instantiates a new ScriptItemVariablesAddRequestBuilder and sets the default values.
func NewScriptItemVariablesAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesAddRequestBuilder) {
    m := &ScriptItemVariablesAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/variables/add", pathParameters),
    }
    return m
}
// NewScriptItemVariablesAddRequestBuilder instantiates a new ScriptItemVariablesAddRequestBuilder and sets the default values.
func NewScriptItemVariablesAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemVariablesAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_addvariable)
// returns a EdgeScriptVariableable when successful
func (m *ScriptItemVariablesAddRequestBuilder) Post(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.EdgeScriptVariableable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.EdgeScriptVariableable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.CreateEdgeScriptVariableFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.EdgeScriptVariableable), nil
}
// ToPostRequestInformation [AddComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_addvariable)
// returns a *RequestInformation when successful
func (m *ScriptItemVariablesAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.EdgeScriptVariableable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptItemVariablesAddRequestBuilder when successful
func (m *ScriptItemVariablesAddRequestBuilder) WithUrl(rawUrl string)(*ScriptItemVariablesAddRequestBuilder) {
    return NewScriptItemVariablesAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
