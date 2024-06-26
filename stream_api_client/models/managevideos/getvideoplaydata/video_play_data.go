package getvideoplaydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos"
)

type VideoPlayData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allowEarlyPlay property
    allowEarlyPlay *bool
    // The captionsBackground property
    captionsBackground *string
    // The captionsFontColor property
    captionsFontColor *string
    // The captionsFontSize property
    captionsFontSize *int32
    // The captionsPath property
    captionsPath *string
    // The controls property
    controls *string
    // The drmVersion property
    drmVersion *int32
    // The enableDRM property
    enableDRM *bool
    // The enableMP4Fallback property
    enableMP4Fallback *bool
    // The fallbackUrl property
    fallbackUrl *string
    // The fontFamily property
    fontFamily *string
    // The originalUrl property
    originalUrl *string
    // The playbackSpeeds property
    playbackSpeeds *string
    // The playerKeyColor property
    playerKeyColor *string
    // The previewUrl property
    previewUrl *string
    // The seekPath property
    seekPath *string
    // The showHeatmap property
    showHeatmap *bool
    // The thumbnailUrl property
    thumbnailUrl *string
    // The tokenAuthEnabled property
    tokenAuthEnabled *bool
    // The uiLanguage property
    uiLanguage *string
    // The vastTagUrl property
    vastTagUrl *string
    // The video property
    video i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable
    // The videoPlaylistUrl property
    videoPlaylistUrl *string
}
// NewVideoPlayData instantiates a new VideoPlayData and sets the default values.
func NewVideoPlayData()(*VideoPlayData) {
    m := &VideoPlayData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideoPlayDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoPlayDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideoPlayData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VideoPlayData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowEarlyPlay gets the allowEarlyPlay property value. The allowEarlyPlay property
// returns a *bool when successful
func (m *VideoPlayData) GetAllowEarlyPlay()(*bool) {
    return m.allowEarlyPlay
}
// GetCaptionsBackground gets the captionsBackground property value. The captionsBackground property
// returns a *string when successful
func (m *VideoPlayData) GetCaptionsBackground()(*string) {
    return m.captionsBackground
}
// GetCaptionsFontColor gets the captionsFontColor property value. The captionsFontColor property
// returns a *string when successful
func (m *VideoPlayData) GetCaptionsFontColor()(*string) {
    return m.captionsFontColor
}
// GetCaptionsFontSize gets the captionsFontSize property value. The captionsFontSize property
// returns a *int32 when successful
func (m *VideoPlayData) GetCaptionsFontSize()(*int32) {
    return m.captionsFontSize
}
// GetCaptionsPath gets the captionsPath property value. The captionsPath property
// returns a *string when successful
func (m *VideoPlayData) GetCaptionsPath()(*string) {
    return m.captionsPath
}
// GetControls gets the controls property value. The controls property
// returns a *string when successful
func (m *VideoPlayData) GetControls()(*string) {
    return m.controls
}
// GetDrmVersion gets the drmVersion property value. The drmVersion property
// returns a *int32 when successful
func (m *VideoPlayData) GetDrmVersion()(*int32) {
    return m.drmVersion
}
// GetEnableDRM gets the enableDRM property value. The enableDRM property
// returns a *bool when successful
func (m *VideoPlayData) GetEnableDRM()(*bool) {
    return m.enableDRM
}
// GetEnableMP4Fallback gets the enableMP4Fallback property value. The enableMP4Fallback property
// returns a *bool when successful
func (m *VideoPlayData) GetEnableMP4Fallback()(*bool) {
    return m.enableMP4Fallback
}
// GetFallbackUrl gets the fallbackUrl property value. The fallbackUrl property
// returns a *string when successful
func (m *VideoPlayData) GetFallbackUrl()(*string) {
    return m.fallbackUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideoPlayData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowEarlyPlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEarlyPlay(val)
        }
        return nil
    }
    res["captionsBackground"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsBackground(val)
        }
        return nil
    }
    res["captionsFontColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsFontColor(val)
        }
        return nil
    }
    res["captionsFontSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsFontSize(val)
        }
        return nil
    }
    res["captionsPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsPath(val)
        }
        return nil
    }
    res["controls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControls(val)
        }
        return nil
    }
    res["drmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrmVersion(val)
        }
        return nil
    }
    res["enableDRM"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDRM(val)
        }
        return nil
    }
    res["enableMP4Fallback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableMP4Fallback(val)
        }
        return nil
    }
    res["fallbackUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFallbackUrl(val)
        }
        return nil
    }
    res["fontFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontFamily(val)
        }
        return nil
    }
    res["originalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalUrl(val)
        }
        return nil
    }
    res["playbackSpeeds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaybackSpeeds(val)
        }
        return nil
    }
    res["playerKeyColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayerKeyColor(val)
        }
        return nil
    }
    res["previewUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewUrl(val)
        }
        return nil
    }
    res["seekPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeekPath(val)
        }
        return nil
    }
    res["showHeatmap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowHeatmap(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["tokenAuthEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenAuthEnabled(val)
        }
        return nil
    }
    res["uiLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUiLanguage(val)
        }
        return nil
    }
    res["vastTagUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVastTagUrl(val)
        }
        return nil
    }
    res["video"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.CreateVideoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val.(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable))
        }
        return nil
    }
    res["videoPlaylistUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoPlaylistUrl(val)
        }
        return nil
    }
    return res
}
// GetFontFamily gets the fontFamily property value. The fontFamily property
// returns a *string when successful
func (m *VideoPlayData) GetFontFamily()(*string) {
    return m.fontFamily
}
// GetOriginalUrl gets the originalUrl property value. The originalUrl property
// returns a *string when successful
func (m *VideoPlayData) GetOriginalUrl()(*string) {
    return m.originalUrl
}
// GetPlaybackSpeeds gets the playbackSpeeds property value. The playbackSpeeds property
// returns a *string when successful
func (m *VideoPlayData) GetPlaybackSpeeds()(*string) {
    return m.playbackSpeeds
}
// GetPlayerKeyColor gets the playerKeyColor property value. The playerKeyColor property
// returns a *string when successful
func (m *VideoPlayData) GetPlayerKeyColor()(*string) {
    return m.playerKeyColor
}
// GetPreviewUrl gets the previewUrl property value. The previewUrl property
// returns a *string when successful
func (m *VideoPlayData) GetPreviewUrl()(*string) {
    return m.previewUrl
}
// GetSeekPath gets the seekPath property value. The seekPath property
// returns a *string when successful
func (m *VideoPlayData) GetSeekPath()(*string) {
    return m.seekPath
}
// GetShowHeatmap gets the showHeatmap property value. The showHeatmap property
// returns a *bool when successful
func (m *VideoPlayData) GetShowHeatmap()(*bool) {
    return m.showHeatmap
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
// returns a *string when successful
func (m *VideoPlayData) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetTokenAuthEnabled gets the tokenAuthEnabled property value. The tokenAuthEnabled property
// returns a *bool when successful
func (m *VideoPlayData) GetTokenAuthEnabled()(*bool) {
    return m.tokenAuthEnabled
}
// GetUiLanguage gets the uiLanguage property value. The uiLanguage property
// returns a *string when successful
func (m *VideoPlayData) GetUiLanguage()(*string) {
    return m.uiLanguage
}
// GetVastTagUrl gets the vastTagUrl property value. The vastTagUrl property
// returns a *string when successful
func (m *VideoPlayData) GetVastTagUrl()(*string) {
    return m.vastTagUrl
}
// GetVideo gets the video property value. The video property
// returns a Videoable when successful
func (m *VideoPlayData) GetVideo()(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable) {
    return m.video
}
// GetVideoPlaylistUrl gets the videoPlaylistUrl property value. The videoPlaylistUrl property
// returns a *string when successful
func (m *VideoPlayData) GetVideoPlaylistUrl()(*string) {
    return m.videoPlaylistUrl
}
// Serialize serializes information the current object
func (m *VideoPlayData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowEarlyPlay", m.GetAllowEarlyPlay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("captionsBackground", m.GetCaptionsBackground())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("captionsFontColor", m.GetCaptionsFontColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("captionsFontSize", m.GetCaptionsFontSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("captionsPath", m.GetCaptionsPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("controls", m.GetControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("drmVersion", m.GetDrmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableDRM", m.GetEnableDRM())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableMP4Fallback", m.GetEnableMP4Fallback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fallbackUrl", m.GetFallbackUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fontFamily", m.GetFontFamily())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("originalUrl", m.GetOriginalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("playbackSpeeds", m.GetPlaybackSpeeds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("playerKeyColor", m.GetPlayerKeyColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewUrl", m.GetPreviewUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("seekPath", m.GetSeekPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showHeatmap", m.GetShowHeatmap())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("tokenAuthEnabled", m.GetTokenAuthEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uiLanguage", m.GetUiLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vastTagUrl", m.GetVastTagUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("videoPlaylistUrl", m.GetVideoPlaylistUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VideoPlayData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowEarlyPlay sets the allowEarlyPlay property value. The allowEarlyPlay property
func (m *VideoPlayData) SetAllowEarlyPlay(value *bool)() {
    m.allowEarlyPlay = value
}
// SetCaptionsBackground sets the captionsBackground property value. The captionsBackground property
func (m *VideoPlayData) SetCaptionsBackground(value *string)() {
    m.captionsBackground = value
}
// SetCaptionsFontColor sets the captionsFontColor property value. The captionsFontColor property
func (m *VideoPlayData) SetCaptionsFontColor(value *string)() {
    m.captionsFontColor = value
}
// SetCaptionsFontSize sets the captionsFontSize property value. The captionsFontSize property
func (m *VideoPlayData) SetCaptionsFontSize(value *int32)() {
    m.captionsFontSize = value
}
// SetCaptionsPath sets the captionsPath property value. The captionsPath property
func (m *VideoPlayData) SetCaptionsPath(value *string)() {
    m.captionsPath = value
}
// SetControls sets the controls property value. The controls property
func (m *VideoPlayData) SetControls(value *string)() {
    m.controls = value
}
// SetDrmVersion sets the drmVersion property value. The drmVersion property
func (m *VideoPlayData) SetDrmVersion(value *int32)() {
    m.drmVersion = value
}
// SetEnableDRM sets the enableDRM property value. The enableDRM property
func (m *VideoPlayData) SetEnableDRM(value *bool)() {
    m.enableDRM = value
}
// SetEnableMP4Fallback sets the enableMP4Fallback property value. The enableMP4Fallback property
func (m *VideoPlayData) SetEnableMP4Fallback(value *bool)() {
    m.enableMP4Fallback = value
}
// SetFallbackUrl sets the fallbackUrl property value. The fallbackUrl property
func (m *VideoPlayData) SetFallbackUrl(value *string)() {
    m.fallbackUrl = value
}
// SetFontFamily sets the fontFamily property value. The fontFamily property
func (m *VideoPlayData) SetFontFamily(value *string)() {
    m.fontFamily = value
}
// SetOriginalUrl sets the originalUrl property value. The originalUrl property
func (m *VideoPlayData) SetOriginalUrl(value *string)() {
    m.originalUrl = value
}
// SetPlaybackSpeeds sets the playbackSpeeds property value. The playbackSpeeds property
func (m *VideoPlayData) SetPlaybackSpeeds(value *string)() {
    m.playbackSpeeds = value
}
// SetPlayerKeyColor sets the playerKeyColor property value. The playerKeyColor property
func (m *VideoPlayData) SetPlayerKeyColor(value *string)() {
    m.playerKeyColor = value
}
// SetPreviewUrl sets the previewUrl property value. The previewUrl property
func (m *VideoPlayData) SetPreviewUrl(value *string)() {
    m.previewUrl = value
}
// SetSeekPath sets the seekPath property value. The seekPath property
func (m *VideoPlayData) SetSeekPath(value *string)() {
    m.seekPath = value
}
// SetShowHeatmap sets the showHeatmap property value. The showHeatmap property
func (m *VideoPlayData) SetShowHeatmap(value *bool)() {
    m.showHeatmap = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *VideoPlayData) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetTokenAuthEnabled sets the tokenAuthEnabled property value. The tokenAuthEnabled property
func (m *VideoPlayData) SetTokenAuthEnabled(value *bool)() {
    m.tokenAuthEnabled = value
}
// SetUiLanguage sets the uiLanguage property value. The uiLanguage property
func (m *VideoPlayData) SetUiLanguage(value *string)() {
    m.uiLanguage = value
}
// SetVastTagUrl sets the vastTagUrl property value. The vastTagUrl property
func (m *VideoPlayData) SetVastTagUrl(value *string)() {
    m.vastTagUrl = value
}
// SetVideo sets the video property value. The video property
func (m *VideoPlayData) SetVideo(value i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable)() {
    m.video = value
}
// SetVideoPlaylistUrl sets the videoPlaylistUrl property value. The videoPlaylistUrl property
func (m *VideoPlayData) SetVideoPlaylistUrl(value *string)() {
    m.videoPlaylistUrl = value
}
type VideoPlayDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowEarlyPlay()(*bool)
    GetCaptionsBackground()(*string)
    GetCaptionsFontColor()(*string)
    GetCaptionsFontSize()(*int32)
    GetCaptionsPath()(*string)
    GetControls()(*string)
    GetDrmVersion()(*int32)
    GetEnableDRM()(*bool)
    GetEnableMP4Fallback()(*bool)
    GetFallbackUrl()(*string)
    GetFontFamily()(*string)
    GetOriginalUrl()(*string)
    GetPlaybackSpeeds()(*string)
    GetPlayerKeyColor()(*string)
    GetPreviewUrl()(*string)
    GetSeekPath()(*string)
    GetShowHeatmap()(*bool)
    GetThumbnailUrl()(*string)
    GetTokenAuthEnabled()(*bool)
    GetUiLanguage()(*string)
    GetVastTagUrl()(*string)
    GetVideo()(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable)
    GetVideoPlaylistUrl()(*string)
    SetAllowEarlyPlay(value *bool)()
    SetCaptionsBackground(value *string)()
    SetCaptionsFontColor(value *string)()
    SetCaptionsFontSize(value *int32)()
    SetCaptionsPath(value *string)()
    SetControls(value *string)()
    SetDrmVersion(value *int32)()
    SetEnableDRM(value *bool)()
    SetEnableMP4Fallback(value *bool)()
    SetFallbackUrl(value *string)()
    SetFontFamily(value *string)()
    SetOriginalUrl(value *string)()
    SetPlaybackSpeeds(value *string)()
    SetPlayerKeyColor(value *string)()
    SetPreviewUrl(value *string)()
    SetSeekPath(value *string)()
    SetShowHeatmap(value *bool)()
    SetThumbnailUrl(value *string)()
    SetTokenAuthEnabled(value *bool)()
    SetUiLanguage(value *string)()
    SetVastTagUrl(value *string)()
    SetVideo(value i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable)()
    SetVideoPlaylistUrl(value *string)()
}
