package compute

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScriptItemVariablesRequestBuilder builds and executes requests for operations under \compute\script\{id}\variables
type ScriptItemVariablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Add the add property
// returns a *ScriptItemVariablesAddRequestBuilder when successful
func (m *ScriptItemVariablesRequestBuilder) Add()(*ScriptItemVariablesAddRequestBuilder) {
    return NewScriptItemVariablesAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByVariableId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.compute.script.item.variables.item collection
// returns a *ScriptItemVariablesWithVariableItemRequestBuilder when successful
func (m *ScriptItemVariablesRequestBuilder) ByVariableId(variableId int64)(*ScriptItemVariablesWithVariableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["variableId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(variableId, 10)
    return NewScriptItemVariablesWithVariableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewScriptItemVariablesRequestBuilderInternal instantiates a new ScriptItemVariablesRequestBuilder and sets the default values.
func NewScriptItemVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesRequestBuilder) {
    m := &ScriptItemVariablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/variables", pathParameters),
    }
    return m
}
// NewScriptItemVariablesRequestBuilder instantiates a new ScriptItemVariablesRequestBuilder and sets the default values.
func NewScriptItemVariablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemVariablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemVariablesRequestBuilderInternal(urlParams, requestAdapter)
}
