package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ShieldZoneGetByPullzoneRequestBuilder builds and executes requests for operations under \shield\shield-zone\get-by-pullzone
type ShieldZoneGetByPullzoneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPullZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.shieldZone.getByPullzone.item collection
// returns a *ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder when successful
func (m *ShieldZoneGetByPullzoneRequestBuilder) ByPullZoneId(pullZoneId int32)(*ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["pullZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(pullZoneId), 10)
    return NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewShieldZoneGetByPullzoneRequestBuilderInternal instantiates a new ShieldZoneGetByPullzoneRequestBuilder and sets the default values.
func NewShieldZoneGetByPullzoneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldZoneGetByPullzoneRequestBuilder) {
    m := &ShieldZoneGetByPullzoneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/shield-zone/get-by-pullzone", pathParameters),
    }
    return m
}
// NewShieldZoneGetByPullzoneRequestBuilder instantiates a new ShieldZoneGetByPullzoneRequestBuilder and sets the default values.
func NewShieldZoneGetByPullzoneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldZoneGetByPullzoneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewShieldZoneGetByPullzoneRequestBuilderInternal(urlParams, requestAdapter)
}
