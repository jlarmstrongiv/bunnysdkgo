package compute

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/compute"
)

// ScriptRequestBuilder builds and executes requests for operations under \compute\script
type ScriptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ScriptRequestBuilderGetQueryParameters [ListComputeScripts API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_index)
type ScriptRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
    // The search term that will be used to filter the results
    Search *string `uriparametername:"search"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.compute.script.item collection
// returns a *ScriptScriptItemRequestBuilder when successful
func (m *ScriptRequestBuilder) ById(id int64)(*ScriptScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewScriptScriptItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewScriptRequestBuilderInternal instantiates a new ScriptRequestBuilder and sets the default values.
func NewScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptRequestBuilder) {
    m := &ScriptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script?page={page}&perPage={perPage}&search={search}", pathParameters),
    }
    return m
}
// NewScriptRequestBuilder instantiates a new ScriptRequestBuilder and sets the default values.
func NewScriptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListComputeScripts API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_index)
// returns a ScriptGetResponseable when successful
func (m *ScriptRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptRequestBuilderGetQueryParameters])(ScriptGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScriptGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScriptGetResponseable), nil
}
// Post [AddComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_addscript)
// returns a Scriptable when successful
func (m *ScriptRequestBuilder) Post(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.ScriptCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.Scriptable, error) {
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
// ToGetRequestInformation [ListComputeScripts API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_index)
// returns a *RequestInformation when successful
func (m *ScriptRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddComputeScript API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_addscript)
// returns a *RequestInformation when successful
func (m *ScriptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3e53788908e7269a4f8084b0fcc3c88f921087c8f03fe97697eff66be83bcfa.ScriptCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/compute/script", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptRequestBuilder when successful
func (m *ScriptRequestBuilder) WithUrl(rawUrl string)(*ScriptRequestBuilder) {
    return NewScriptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
