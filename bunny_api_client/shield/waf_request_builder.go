package shield

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WafRequestBuilder builds and executes requests for operations under \shield\waf
type WafRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWafRequestBuilderInternal instantiates a new WafRequestBuilder and sets the default values.
func NewWafRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafRequestBuilder) {
    m := &WafRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/waf", pathParameters),
    }
    return m
}
// NewWafRequestBuilder instantiates a new WafRequestBuilder and sets the default values.
func NewWafRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WafRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWafRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomRule the customRule property
// returns a *WafCustomRuleRequestBuilder when successful
func (m *WafRequestBuilder) CustomRule()(*WafCustomRuleRequestBuilder) {
    return NewWafCustomRuleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomRules the customRules property
// returns a *WafCustomRulesRequestBuilder when successful
func (m *WafRequestBuilder) CustomRules()(*WafCustomRulesRequestBuilder) {
    return NewWafCustomRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EngineConfig the engineConfig property
// returns a *WafEngineConfigRequestBuilder when successful
func (m *WafRequestBuilder) EngineConfig()(*WafEngineConfigRequestBuilder) {
    return NewWafEngineConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enums the enums property
// returns a *WafEnumsRequestBuilder when successful
func (m *WafRequestBuilder) Enums()(*WafEnumsRequestBuilder) {
    return NewWafEnumsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Profiles the profiles property
// returns a *WafProfilesRequestBuilder when successful
func (m *WafRequestBuilder) Profiles()(*WafProfilesRequestBuilder) {
    return NewWafProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rules the rules property
// returns a *WafRulesRequestBuilder when successful
func (m *WafRequestBuilder) Rules()(*WafRulesRequestBuilder) {
    return NewWafRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
