package shield

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/shield"
)

// WafRulesRequestBuilder builds and executes requests for operations under \shield\waf\rules
type WafRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWafRulesRequestBuilderInternal instantiates a new WafRulesRequestBuilder and sets the default values.
func NewWafRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafRulesRequestBuilder) {
    m := &WafRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/waf/rules", pathParameters),
    }
    return m
}
// NewWafRulesRequestBuilder instantiates a new WafRulesRequestBuilder and sets the default values.
func NewWafRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWafRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []WafRuleMainGroupModelable when successful
// returns a UnauthorizedResult error when the service returns a 401 status code
func (m *WafRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.WafRuleMainGroupModelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateUnauthorizedResultFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateWafRuleMainGroupModelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.WafRuleMainGroupModelable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.WafRuleMainGroupModelable)
        }
    }
    return val, nil
}
// returns a *RequestInformation when successful
func (m *WafRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WafRulesRequestBuilder when successful
func (m *WafRulesRequestBuilder) WithUrl(rawUrl string)(*WafRulesRequestBuilder) {
    return NewWafRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
