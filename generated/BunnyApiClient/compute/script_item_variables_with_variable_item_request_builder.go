package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/compute"
)

// ScriptItemVariablesWithVariableItemRequestBuilder builds and executes requests for operations under \compute\script\{id}\variables\{variableId}
type ScriptItemVariablesWithVariableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewScriptItemVariablesWithVariableItemRequestBuilderInternal instantiates a new ScriptItemVariablesWithVariableItemRequestBuilder and sets the default values.
func NewScriptItemVariablesWithVariableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesWithVariableItemRequestBuilder) {
    m := &ScriptItemVariablesWithVariableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/variables/{variableId}", pathParameters),
    }
    return m
}
// NewScriptItemVariablesWithVariableItemRequestBuilder instantiates a new ScriptItemVariablesWithVariableItemRequestBuilder and sets the default values.
func NewScriptItemVariablesWithVariableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesWithVariableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemVariablesWithVariableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_deletevariable)
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// Get [GetComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_getvariable)
// returns a EdgeScriptVariableable when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.CreateEdgeScriptVariableFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable), nil
}
// Post [UpdateComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_updatevariable)
// returns a EdgeScriptVariableable when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) Post(ctx context.Context, body i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.CreateEdgeScriptVariableFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable), nil
}
// ToDeleteRequestInformation [DeleteComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_deletevariable)
// returns a *RequestInformation when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation [GetComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_getvariable)
// returns a *RequestInformation when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateComputeScriptVariable API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_updatevariable)
// returns a *RequestInformation when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i38a965aeb9236c992f545da633385ad6c9c2260b710b77a35ada2503fa3a982c.EdgeScriptVariableable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ScriptItemVariablesWithVariableItemRequestBuilder when successful
func (m *ScriptItemVariablesWithVariableItemRequestBuilder) WithUrl(rawUrl string)(*ScriptItemVariablesWithVariableItemRequestBuilder) {
    return NewScriptItemVariablesWithVariableItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
