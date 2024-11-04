package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WafCustomRulesRequestBuilder builds and executes requests for operations under \shield\waf\custom-rules
type WafCustomRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.waf.customRules.item collection
// returns a *WafCustomRulesWithShieldZoneItemRequestBuilder when successful
func (m *WafCustomRulesRequestBuilder) ByShieldZoneId(shieldZoneId int64)(*WafCustomRulesWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(shieldZoneId, 10)
    return NewWafCustomRulesWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWafCustomRulesRequestBuilderInternal instantiates a new WafCustomRulesRequestBuilder and sets the default values.
func NewWafCustomRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafCustomRulesRequestBuilder) {
    m := &WafCustomRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/waf/custom-rules", pathParameters),
    }
    return m
}
// NewWafCustomRulesRequestBuilder instantiates a new WafCustomRulesRequestBuilder and sets the default values.
func NewWafCustomRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafCustomRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWafCustomRulesRequestBuilderInternal(urlParams, requestAdapter)
}
