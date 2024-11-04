package streamvideolibrary

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VideoLibrary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines direct play URLs are enabled for the library
    allowDirectPlay *bool
    // Determines if the Early-Play feature is enabled
    allowEarlyPlay *bool
    // The list of allowed referrer domains allowed to access the library
    allowedReferrers []string
    // The API access key for the library. Only added when the includeAccessKey parameter is set.
    apiAccessKey *string
    // The API key used for authenticating with the video library
    apiKey *string
    // The AppleFairPlayDrm property
    appleFairPlayDrm VideoLibrary_AppleFairPlayDrmable
    // The bitrate used for encoding 1080p videos
    bitrate1080p *int32
    // The bitrate used for encoding 1440p videos
    bitrate1440p *int32
    // The bitrate used for encoding 2160p videos
    bitrate2160p *int32
    // The bitrate used for encoding 240p videos
    bitrate240p *int32
    // The bitrate used for encoding 360p videos
    bitrate360p *int32
    // The bitrate used for encoding 480p videos
    bitrate480p *int32
    // The bitrate used for encoding 720p videos
    bitrate720p *int32
    // The list of blocked referrer domains blocked from accessing the library
    blockedReferrers []string
    // Determines if the requests without a referrer are blocked
    blockNoneReferrer *bool
    // The captions display background color
    captionsBackground *string
    // The captions display font color
    captionsFontColor *string
    // The captions display font size
    captionsFontSize *int32
    // The list of controls on the video player.
    controls *string
    // The custom HTMl that is added into the head of the HTML player.
    customHTML *string
    // The date when the video library was created
    dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DrmBasePriceOverride property
    drmBasePriceOverride *float64
    // The DrmCostPerLicenseOverride property
    drmCostPerLicenseOverride *float64
    // The DrmVersion property
    drmVersion *int64
    // Determines if content tagging should be enabled for this library.
    enableContentTagging *bool
    // The comma separated list of enabled resolutions
    enabledResolutions *string
    // Determines if the MediaCage basic DRM is enabled
    enableDRM *bool
    // Determines if the MP4 fallback feature is enabled
    enableMP4Fallback *bool
    // The EnableMultiAudioTrackSupport property
    enableMultiAudioTrackSupport *bool
    // Determines if the automatic audio transcribing is currently enabled for this zone.
    enableTranscribing *bool
    // Determines if automatic transcribing description generation is currently enabled.
    enableTranscribingDescriptionGeneration *bool
    // Determines if automatic transcribing title generation is currently enabled.
    enableTranscribingTitleGeneration *bool
    // The EncodingTier property
    encodingTier *int32
    // The captions font family.
    fontFamily *string
    // The GoogleWidevineDrm property
    googleWidevineDrm VideoLibrary_GoogleWidevineDrmable
    // Determines if the video library has a watermark configured
    hasWatermark *bool
    // The Id property
    id *int64
    // The JitEncodingEnabled property
    jitEncodingEnabled *bool
    // Determines if the original video files should be stored after encoding
    keepOriginalFiles *bool
    // The MonthlyChargesEnterpriseDrm property
    monthlyChargesEnterpriseDrm *float64
    // The MonthlyChargesPremiumEncoding property
    monthlyChargesPremiumEncoding *float64
    // The MonthlyChargesTranscribing property
    monthlyChargesTranscribing *float64
    // The name of the Video Library.
    name *string
    // The OutputCodecs property
    outputCodecs *string
    // The key color of the player.
    playerKeyColor *string
    // Determines if the player token authentication is enabled
    playerTokenAuthenticationEnabled *bool
    // The PremiumEncodingPriceOverride property
    premiumEncodingPriceOverride *float64
    // The ID of the connected underlying pull zone
    pullZoneId *int64
    // The PullZoneType property
    pullZoneType *float64
    // The read-only API key used for authenticating with the video library
    readOnlyApiKey *string
    // The list of languages that the captions will be automatically transcribed to.
    rememberPlayerPosition *bool
    // The geo-replication regions of the underlying storage zone
    replicationRegions []ReplicationRegions
    // Determines if the video watch heatmap should be displayed in the player.
    showHeatmap *bool
    // The total amount of storage used by the library
    storageUsage *int64
    // The ID of the connected underlying storage zone
    storageZoneId *int64
    // The amount of traffic usage this month
    trafficUsage *int64
    // The TranscribingCaptionLanguages property
    transcribingCaptionLanguages []string
    // The TranscribingPriceOverride property
    transcribingPriceOverride *float64
    // The UI language of the player
    uILanguage *string
    // The UseSeparateAudioStream property
    useSeparateAudioStream *bool
    // The URL of the VAST tag endpoint for advertising configuration
    vastTagUrl *string
    // The vi.ai publisher id for advertising configuration
    viAiPublisherId *string
    // The number of videos in the video library
    videoCount *int64
    // The height of the watermark (in %)
    watermarkHeight *int32
    // The left offset of the watermark position (in %)
    watermarkPositionLeft *int32
    // The top offset of the watermark position (in %)
    watermarkPositionTop *int32
    // The WatermarkVersion property
    watermarkVersion *int64
    // The width of the watermark (in %)
    watermarkWidth *int32
    // The webhook URL of the video library
    webhookUrl *string
}
// NewVideoLibrary instantiates a new VideoLibrary and sets the default values.
func NewVideoLibrary()(*VideoLibrary) {
    m := &VideoLibrary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideoLibraryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoLibraryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideoLibrary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VideoLibrary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDirectPlay gets the AllowDirectPlay property value. Determines direct play URLs are enabled for the library
// returns a *bool when successful
func (m *VideoLibrary) GetAllowDirectPlay()(*bool) {
    return m.allowDirectPlay
}
// GetAllowEarlyPlay gets the AllowEarlyPlay property value. Determines if the Early-Play feature is enabled
// returns a *bool when successful
func (m *VideoLibrary) GetAllowEarlyPlay()(*bool) {
    return m.allowEarlyPlay
}
// GetAllowedReferrers gets the AllowedReferrers property value. The list of allowed referrer domains allowed to access the library
// returns a []string when successful
func (m *VideoLibrary) GetAllowedReferrers()([]string) {
    return m.allowedReferrers
}
// GetApiAccessKey gets the ApiAccessKey property value. The API access key for the library. Only added when the includeAccessKey parameter is set.
// returns a *string when successful
func (m *VideoLibrary) GetApiAccessKey()(*string) {
    return m.apiAccessKey
}
// GetApiKey gets the ApiKey property value. The API key used for authenticating with the video library
// returns a *string when successful
func (m *VideoLibrary) GetApiKey()(*string) {
    return m.apiKey
}
// GetAppleFairPlayDrm gets the AppleFairPlayDrm property value. The AppleFairPlayDrm property
// returns a VideoLibrary_AppleFairPlayDrmable when successful
func (m *VideoLibrary) GetAppleFairPlayDrm()(VideoLibrary_AppleFairPlayDrmable) {
    return m.appleFairPlayDrm
}
// GetBitrate1080p gets the Bitrate1080p property value. The bitrate used for encoding 1080p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate1080p()(*int32) {
    return m.bitrate1080p
}
// GetBitrate1440p gets the Bitrate1440p property value. The bitrate used for encoding 1440p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate1440p()(*int32) {
    return m.bitrate1440p
}
// GetBitrate2160p gets the Bitrate2160p property value. The bitrate used for encoding 2160p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate2160p()(*int32) {
    return m.bitrate2160p
}
// GetBitrate240p gets the Bitrate240p property value. The bitrate used for encoding 240p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate240p()(*int32) {
    return m.bitrate240p
}
// GetBitrate360p gets the Bitrate360p property value. The bitrate used for encoding 360p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate360p()(*int32) {
    return m.bitrate360p
}
// GetBitrate480p gets the Bitrate480p property value. The bitrate used for encoding 480p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate480p()(*int32) {
    return m.bitrate480p
}
// GetBitrate720p gets the Bitrate720p property value. The bitrate used for encoding 720p videos
// returns a *int32 when successful
func (m *VideoLibrary) GetBitrate720p()(*int32) {
    return m.bitrate720p
}
// GetBlockedReferrers gets the BlockedReferrers property value. The list of blocked referrer domains blocked from accessing the library
// returns a []string when successful
func (m *VideoLibrary) GetBlockedReferrers()([]string) {
    return m.blockedReferrers
}
// GetBlockNoneReferrer gets the BlockNoneReferrer property value. Determines if the requests without a referrer are blocked
// returns a *bool when successful
func (m *VideoLibrary) GetBlockNoneReferrer()(*bool) {
    return m.blockNoneReferrer
}
// GetCaptionsBackground gets the CaptionsBackground property value. The captions display background color
// returns a *string when successful
func (m *VideoLibrary) GetCaptionsBackground()(*string) {
    return m.captionsBackground
}
// GetCaptionsFontColor gets the CaptionsFontColor property value. The captions display font color
// returns a *string when successful
func (m *VideoLibrary) GetCaptionsFontColor()(*string) {
    return m.captionsFontColor
}
// GetCaptionsFontSize gets the CaptionsFontSize property value. The captions display font size
// returns a *int32 when successful
func (m *VideoLibrary) GetCaptionsFontSize()(*int32) {
    return m.captionsFontSize
}
// GetControls gets the Controls property value. The list of controls on the video player.
// returns a *string when successful
func (m *VideoLibrary) GetControls()(*string) {
    return m.controls
}
// GetCustomHTML gets the CustomHTML property value. The custom HTMl that is added into the head of the HTML player.
// returns a *string when successful
func (m *VideoLibrary) GetCustomHTML()(*string) {
    return m.customHTML
}
// GetDateCreated gets the DateCreated property value. The date when the video library was created
// returns a *Time when successful
func (m *VideoLibrary) GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateCreated
}
// GetDrmBasePriceOverride gets the DrmBasePriceOverride property value. The DrmBasePriceOverride property
// returns a *float64 when successful
func (m *VideoLibrary) GetDrmBasePriceOverride()(*float64) {
    return m.drmBasePriceOverride
}
// GetDrmCostPerLicenseOverride gets the DrmCostPerLicenseOverride property value. The DrmCostPerLicenseOverride property
// returns a *float64 when successful
func (m *VideoLibrary) GetDrmCostPerLicenseOverride()(*float64) {
    return m.drmCostPerLicenseOverride
}
// GetDrmVersion gets the DrmVersion property value. The DrmVersion property
// returns a *int64 when successful
func (m *VideoLibrary) GetDrmVersion()(*int64) {
    return m.drmVersion
}
// GetEnableContentTagging gets the EnableContentTagging property value. Determines if content tagging should be enabled for this library.
// returns a *bool when successful
func (m *VideoLibrary) GetEnableContentTagging()(*bool) {
    return m.enableContentTagging
}
// GetEnabledResolutions gets the EnabledResolutions property value. The comma separated list of enabled resolutions
// returns a *string when successful
func (m *VideoLibrary) GetEnabledResolutions()(*string) {
    return m.enabledResolutions
}
// GetEnableDRM gets the EnableDRM property value. Determines if the MediaCage basic DRM is enabled
// returns a *bool when successful
func (m *VideoLibrary) GetEnableDRM()(*bool) {
    return m.enableDRM
}
// GetEnableMP4Fallback gets the EnableMP4Fallback property value. Determines if the MP4 fallback feature is enabled
// returns a *bool when successful
func (m *VideoLibrary) GetEnableMP4Fallback()(*bool) {
    return m.enableMP4Fallback
}
// GetEnableMultiAudioTrackSupport gets the EnableMultiAudioTrackSupport property value. The EnableMultiAudioTrackSupport property
// returns a *bool when successful
func (m *VideoLibrary) GetEnableMultiAudioTrackSupport()(*bool) {
    return m.enableMultiAudioTrackSupport
}
// GetEnableTranscribing gets the EnableTranscribing property value. Determines if the automatic audio transcribing is currently enabled for this zone.
// returns a *bool when successful
func (m *VideoLibrary) GetEnableTranscribing()(*bool) {
    return m.enableTranscribing
}
// GetEnableTranscribingDescriptionGeneration gets the EnableTranscribingDescriptionGeneration property value. Determines if automatic transcribing description generation is currently enabled.
// returns a *bool when successful
func (m *VideoLibrary) GetEnableTranscribingDescriptionGeneration()(*bool) {
    return m.enableTranscribingDescriptionGeneration
}
// GetEnableTranscribingTitleGeneration gets the EnableTranscribingTitleGeneration property value. Determines if automatic transcribing title generation is currently enabled.
// returns a *bool when successful
func (m *VideoLibrary) GetEnableTranscribingTitleGeneration()(*bool) {
    return m.enableTranscribingTitleGeneration
}
// GetEncodingTier gets the EncodingTier property value. The EncodingTier property
// returns a *int32 when successful
func (m *VideoLibrary) GetEncodingTier()(*int32) {
    return m.encodingTier
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideoLibrary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AllowDirectPlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDirectPlay(val)
        }
        return nil
    }
    res["AllowEarlyPlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEarlyPlay(val)
        }
        return nil
    }
    res["AllowedReferrers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAllowedReferrers(res)
        }
        return nil
    }
    res["ApiAccessKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiAccessKey(val)
        }
        return nil
    }
    res["ApiKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiKey(val)
        }
        return nil
    }
    res["AppleFairPlayDrm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVideoLibrary_AppleFairPlayDrmFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleFairPlayDrm(val.(VideoLibrary_AppleFairPlayDrmable))
        }
        return nil
    }
    res["Bitrate1080p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate1080p(val)
        }
        return nil
    }
    res["Bitrate1440p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate1440p(val)
        }
        return nil
    }
    res["Bitrate2160p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate2160p(val)
        }
        return nil
    }
    res["Bitrate240p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate240p(val)
        }
        return nil
    }
    res["Bitrate360p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate360p(val)
        }
        return nil
    }
    res["Bitrate480p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate480p(val)
        }
        return nil
    }
    res["Bitrate720p"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate720p(val)
        }
        return nil
    }
    res["BlockedReferrers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetBlockedReferrers(res)
        }
        return nil
    }
    res["BlockNoneReferrer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockNoneReferrer(val)
        }
        return nil
    }
    res["CaptionsBackground"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsBackground(val)
        }
        return nil
    }
    res["CaptionsFontColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsFontColor(val)
        }
        return nil
    }
    res["CaptionsFontSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsFontSize(val)
        }
        return nil
    }
    res["Controls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControls(val)
        }
        return nil
    }
    res["CustomHTML"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomHTML(val)
        }
        return nil
    }
    res["DateCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateCreated(val)
        }
        return nil
    }
    res["DrmBasePriceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrmBasePriceOverride(val)
        }
        return nil
    }
    res["DrmCostPerLicenseOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrmCostPerLicenseOverride(val)
        }
        return nil
    }
    res["DrmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrmVersion(val)
        }
        return nil
    }
    res["EnableContentTagging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableContentTagging(val)
        }
        return nil
    }
    res["EnabledResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledResolutions(val)
        }
        return nil
    }
    res["EnableDRM"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDRM(val)
        }
        return nil
    }
    res["EnableMP4Fallback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableMP4Fallback(val)
        }
        return nil
    }
    res["EnableMultiAudioTrackSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableMultiAudioTrackSupport(val)
        }
        return nil
    }
    res["EnableTranscribing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTranscribing(val)
        }
        return nil
    }
    res["EnableTranscribingDescriptionGeneration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTranscribingDescriptionGeneration(val)
        }
        return nil
    }
    res["EnableTranscribingTitleGeneration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTranscribingTitleGeneration(val)
        }
        return nil
    }
    res["EncodingTier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncodingTier(val)
        }
        return nil
    }
    res["FontFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontFamily(val)
        }
        return nil
    }
    res["GoogleWidevineDrm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVideoLibrary_GoogleWidevineDrmFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoogleWidevineDrm(val.(VideoLibrary_GoogleWidevineDrmable))
        }
        return nil
    }
    res["HasWatermark"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasWatermark(val)
        }
        return nil
    }
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["JitEncodingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJitEncodingEnabled(val)
        }
        return nil
    }
    res["KeepOriginalFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepOriginalFiles(val)
        }
        return nil
    }
    res["MonthlyChargesEnterpriseDrm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyChargesEnterpriseDrm(val)
        }
        return nil
    }
    res["MonthlyChargesPremiumEncoding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyChargesPremiumEncoding(val)
        }
        return nil
    }
    res["MonthlyChargesTranscribing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyChargesTranscribing(val)
        }
        return nil
    }
    res["Name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["OutputCodecs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputCodecs(val)
        }
        return nil
    }
    res["PlayerKeyColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayerKeyColor(val)
        }
        return nil
    }
    res["PlayerTokenAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayerTokenAuthenticationEnabled(val)
        }
        return nil
    }
    res["PremiumEncodingPriceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPremiumEncodingPriceOverride(val)
        }
        return nil
    }
    res["PullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["PullZoneType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneType(val)
        }
        return nil
    }
    res["ReadOnlyApiKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnlyApiKey(val)
        }
        return nil
    }
    res["RememberPlayerPosition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRememberPlayerPosition(val)
        }
        return nil
    }
    res["ReplicationRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseReplicationRegions)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReplicationRegions, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ReplicationRegions))
                }
            }
            m.SetReplicationRegions(res)
        }
        return nil
    }
    res["ShowHeatmap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowHeatmap(val)
        }
        return nil
    }
    res["StorageUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsage(val)
        }
        return nil
    }
    res["StorageZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageZoneId(val)
        }
        return nil
    }
    res["TrafficUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficUsage(val)
        }
        return nil
    }
    res["TranscribingCaptionLanguages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTranscribingCaptionLanguages(res)
        }
        return nil
    }
    res["TranscribingPriceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranscribingPriceOverride(val)
        }
        return nil
    }
    res["UILanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUILanguage(val)
        }
        return nil
    }
    res["UseSeparateAudioStream"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSeparateAudioStream(val)
        }
        return nil
    }
    res["VastTagUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVastTagUrl(val)
        }
        return nil
    }
    res["ViAiPublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViAiPublisherId(val)
        }
        return nil
    }
    res["VideoCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoCount(val)
        }
        return nil
    }
    res["WatermarkHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkHeight(val)
        }
        return nil
    }
    res["WatermarkPositionLeft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkPositionLeft(val)
        }
        return nil
    }
    res["WatermarkPositionTop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkPositionTop(val)
        }
        return nil
    }
    res["WatermarkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkVersion(val)
        }
        return nil
    }
    res["WatermarkWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkWidth(val)
        }
        return nil
    }
    res["WebhookUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhookUrl(val)
        }
        return nil
    }
    return res
}
// GetFontFamily gets the FontFamily property value. The captions font family.
// returns a *string when successful
func (m *VideoLibrary) GetFontFamily()(*string) {
    return m.fontFamily
}
// GetGoogleWidevineDrm gets the GoogleWidevineDrm property value. The GoogleWidevineDrm property
// returns a VideoLibrary_GoogleWidevineDrmable when successful
func (m *VideoLibrary) GetGoogleWidevineDrm()(VideoLibrary_GoogleWidevineDrmable) {
    return m.googleWidevineDrm
}
// GetHasWatermark gets the HasWatermark property value. Determines if the video library has a watermark configured
// returns a *bool when successful
func (m *VideoLibrary) GetHasWatermark()(*bool) {
    return m.hasWatermark
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *VideoLibrary) GetId()(*int64) {
    return m.id
}
// GetJitEncodingEnabled gets the JitEncodingEnabled property value. The JitEncodingEnabled property
// returns a *bool when successful
func (m *VideoLibrary) GetJitEncodingEnabled()(*bool) {
    return m.jitEncodingEnabled
}
// GetKeepOriginalFiles gets the KeepOriginalFiles property value. Determines if the original video files should be stored after encoding
// returns a *bool when successful
func (m *VideoLibrary) GetKeepOriginalFiles()(*bool) {
    return m.keepOriginalFiles
}
// GetMonthlyChargesEnterpriseDrm gets the MonthlyChargesEnterpriseDrm property value. The MonthlyChargesEnterpriseDrm property
// returns a *float64 when successful
func (m *VideoLibrary) GetMonthlyChargesEnterpriseDrm()(*float64) {
    return m.monthlyChargesEnterpriseDrm
}
// GetMonthlyChargesPremiumEncoding gets the MonthlyChargesPremiumEncoding property value. The MonthlyChargesPremiumEncoding property
// returns a *float64 when successful
func (m *VideoLibrary) GetMonthlyChargesPremiumEncoding()(*float64) {
    return m.monthlyChargesPremiumEncoding
}
// GetMonthlyChargesTranscribing gets the MonthlyChargesTranscribing property value. The MonthlyChargesTranscribing property
// returns a *float64 when successful
func (m *VideoLibrary) GetMonthlyChargesTranscribing()(*float64) {
    return m.monthlyChargesTranscribing
}
// GetName gets the Name property value. The name of the Video Library.
// returns a *string when successful
func (m *VideoLibrary) GetName()(*string) {
    return m.name
}
// GetOutputCodecs gets the OutputCodecs property value. The OutputCodecs property
// returns a *string when successful
func (m *VideoLibrary) GetOutputCodecs()(*string) {
    return m.outputCodecs
}
// GetPlayerKeyColor gets the PlayerKeyColor property value. The key color of the player.
// returns a *string when successful
func (m *VideoLibrary) GetPlayerKeyColor()(*string) {
    return m.playerKeyColor
}
// GetPlayerTokenAuthenticationEnabled gets the PlayerTokenAuthenticationEnabled property value. Determines if the player token authentication is enabled
// returns a *bool when successful
func (m *VideoLibrary) GetPlayerTokenAuthenticationEnabled()(*bool) {
    return m.playerTokenAuthenticationEnabled
}
// GetPremiumEncodingPriceOverride gets the PremiumEncodingPriceOverride property value. The PremiumEncodingPriceOverride property
// returns a *float64 when successful
func (m *VideoLibrary) GetPremiumEncodingPriceOverride()(*float64) {
    return m.premiumEncodingPriceOverride
}
// GetPullZoneId gets the PullZoneId property value. The ID of the connected underlying pull zone
// returns a *int64 when successful
func (m *VideoLibrary) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetPullZoneType gets the PullZoneType property value. The PullZoneType property
// returns a *float64 when successful
func (m *VideoLibrary) GetPullZoneType()(*float64) {
    return m.pullZoneType
}
// GetReadOnlyApiKey gets the ReadOnlyApiKey property value. The read-only API key used for authenticating with the video library
// returns a *string when successful
func (m *VideoLibrary) GetReadOnlyApiKey()(*string) {
    return m.readOnlyApiKey
}
// GetRememberPlayerPosition gets the RememberPlayerPosition property value. The list of languages that the captions will be automatically transcribed to.
// returns a *bool when successful
func (m *VideoLibrary) GetRememberPlayerPosition()(*bool) {
    return m.rememberPlayerPosition
}
// GetReplicationRegions gets the ReplicationRegions property value. The geo-replication regions of the underlying storage zone
// returns a []ReplicationRegions when successful
func (m *VideoLibrary) GetReplicationRegions()([]ReplicationRegions) {
    return m.replicationRegions
}
// GetShowHeatmap gets the ShowHeatmap property value. Determines if the video watch heatmap should be displayed in the player.
// returns a *bool when successful
func (m *VideoLibrary) GetShowHeatmap()(*bool) {
    return m.showHeatmap
}
// GetStorageUsage gets the StorageUsage property value. The total amount of storage used by the library
// returns a *int64 when successful
func (m *VideoLibrary) GetStorageUsage()(*int64) {
    return m.storageUsage
}
// GetStorageZoneId gets the StorageZoneId property value. The ID of the connected underlying storage zone
// returns a *int64 when successful
func (m *VideoLibrary) GetStorageZoneId()(*int64) {
    return m.storageZoneId
}
// GetTrafficUsage gets the TrafficUsage property value. The amount of traffic usage this month
// returns a *int64 when successful
func (m *VideoLibrary) GetTrafficUsage()(*int64) {
    return m.trafficUsage
}
// GetTranscribingCaptionLanguages gets the TranscribingCaptionLanguages property value. The TranscribingCaptionLanguages property
// returns a []string when successful
func (m *VideoLibrary) GetTranscribingCaptionLanguages()([]string) {
    return m.transcribingCaptionLanguages
}
// GetTranscribingPriceOverride gets the TranscribingPriceOverride property value. The TranscribingPriceOverride property
// returns a *float64 when successful
func (m *VideoLibrary) GetTranscribingPriceOverride()(*float64) {
    return m.transcribingPriceOverride
}
// GetUILanguage gets the UILanguage property value. The UI language of the player
// returns a *string when successful
func (m *VideoLibrary) GetUILanguage()(*string) {
    return m.uILanguage
}
// GetUseSeparateAudioStream gets the UseSeparateAudioStream property value. The UseSeparateAudioStream property
// returns a *bool when successful
func (m *VideoLibrary) GetUseSeparateAudioStream()(*bool) {
    return m.useSeparateAudioStream
}
// GetVastTagUrl gets the VastTagUrl property value. The URL of the VAST tag endpoint for advertising configuration
// returns a *string when successful
func (m *VideoLibrary) GetVastTagUrl()(*string) {
    return m.vastTagUrl
}
// GetViAiPublisherId gets the ViAiPublisherId property value. The vi.ai publisher id for advertising configuration
// returns a *string when successful
func (m *VideoLibrary) GetViAiPublisherId()(*string) {
    return m.viAiPublisherId
}
// GetVideoCount gets the VideoCount property value. The number of videos in the video library
// returns a *int64 when successful
func (m *VideoLibrary) GetVideoCount()(*int64) {
    return m.videoCount
}
// GetWatermarkHeight gets the WatermarkHeight property value. The height of the watermark (in %)
// returns a *int32 when successful
func (m *VideoLibrary) GetWatermarkHeight()(*int32) {
    return m.watermarkHeight
}
// GetWatermarkPositionLeft gets the WatermarkPositionLeft property value. The left offset of the watermark position (in %)
// returns a *int32 when successful
func (m *VideoLibrary) GetWatermarkPositionLeft()(*int32) {
    return m.watermarkPositionLeft
}
// GetWatermarkPositionTop gets the WatermarkPositionTop property value. The top offset of the watermark position (in %)
// returns a *int32 when successful
func (m *VideoLibrary) GetWatermarkPositionTop()(*int32) {
    return m.watermarkPositionTop
}
// GetWatermarkVersion gets the WatermarkVersion property value. The WatermarkVersion property
// returns a *int64 when successful
func (m *VideoLibrary) GetWatermarkVersion()(*int64) {
    return m.watermarkVersion
}
// GetWatermarkWidth gets the WatermarkWidth property value. The width of the watermark (in %)
// returns a *int32 when successful
func (m *VideoLibrary) GetWatermarkWidth()(*int32) {
    return m.watermarkWidth
}
// GetWebhookUrl gets the WebhookUrl property value. The webhook URL of the video library
// returns a *string when successful
func (m *VideoLibrary) GetWebhookUrl()(*string) {
    return m.webhookUrl
}
// Serialize serializes information the current object
func (m *VideoLibrary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("AllowDirectPlay", m.GetAllowDirectPlay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("AllowEarlyPlay", m.GetAllowEarlyPlay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate1080p", m.GetBitrate1080p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate1440p", m.GetBitrate1440p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate2160p", m.GetBitrate2160p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate240p", m.GetBitrate240p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate360p", m.GetBitrate360p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate480p", m.GetBitrate480p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Bitrate720p", m.GetBitrate720p())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("BlockNoneReferrer", m.GetBlockNoneReferrer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CaptionsBackground", m.GetCaptionsBackground())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CaptionsFontColor", m.GetCaptionsFontColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("CaptionsFontSize", m.GetCaptionsFontSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Controls", m.GetControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CustomHTML", m.GetCustomHTML())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("DrmBasePriceOverride", m.GetDrmBasePriceOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("DrmCostPerLicenseOverride", m.GetDrmCostPerLicenseOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableContentTagging", m.GetEnableContentTagging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("EnabledResolutions", m.GetEnabledResolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableDRM", m.GetEnableDRM())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableMP4Fallback", m.GetEnableMP4Fallback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableMultiAudioTrackSupport", m.GetEnableMultiAudioTrackSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableTranscribing", m.GetEnableTranscribing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableTranscribingDescriptionGeneration", m.GetEnableTranscribingDescriptionGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableTranscribingTitleGeneration", m.GetEnableTranscribingTitleGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("EncodingTier", m.GetEncodingTier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("FontFamily", m.GetFontFamily())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("JitEncodingEnabled", m.GetJitEncodingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("KeepOriginalFiles", m.GetKeepOriginalFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyChargesEnterpriseDrm", m.GetMonthlyChargesEnterpriseDrm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyChargesPremiumEncoding", m.GetMonthlyChargesPremiumEncoding())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyChargesTranscribing", m.GetMonthlyChargesTranscribing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OutputCodecs", m.GetOutputCodecs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PlayerKeyColor", m.GetPlayerKeyColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("PlayerTokenAuthenticationEnabled", m.GetPlayerTokenAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("PremiumEncodingPriceOverride", m.GetPremiumEncodingPriceOverride())
        if err != nil {
            return err
        }
    }
    if m.GetReplicationRegions() != nil {
        err := writer.WriteCollectionOfStringValues("ReplicationRegions", SerializeReplicationRegions(m.GetReplicationRegions()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ShowHeatmap", m.GetShowHeatmap())
        if err != nil {
            return err
        }
    }
    if m.GetTranscribingCaptionLanguages() != nil {
        err := writer.WriteCollectionOfStringValues("TranscribingCaptionLanguages", m.GetTranscribingCaptionLanguages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("TranscribingPriceOverride", m.GetTranscribingPriceOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("UILanguage", m.GetUILanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("UseSeparateAudioStream", m.GetUseSeparateAudioStream())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("VastTagUrl", m.GetVastTagUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ViAiPublisherId", m.GetViAiPublisherId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("WatermarkHeight", m.GetWatermarkHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("WatermarkPositionLeft", m.GetWatermarkPositionLeft())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("WatermarkPositionTop", m.GetWatermarkPositionTop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("WatermarkWidth", m.GetWatermarkWidth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("WebhookUrl", m.GetWebhookUrl())
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
func (m *VideoLibrary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDirectPlay sets the AllowDirectPlay property value. Determines direct play URLs are enabled for the library
func (m *VideoLibrary) SetAllowDirectPlay(value *bool)() {
    m.allowDirectPlay = value
}
// SetAllowEarlyPlay sets the AllowEarlyPlay property value. Determines if the Early-Play feature is enabled
func (m *VideoLibrary) SetAllowEarlyPlay(value *bool)() {
    m.allowEarlyPlay = value
}
// SetAllowedReferrers sets the AllowedReferrers property value. The list of allowed referrer domains allowed to access the library
func (m *VideoLibrary) SetAllowedReferrers(value []string)() {
    m.allowedReferrers = value
}
// SetApiAccessKey sets the ApiAccessKey property value. The API access key for the library. Only added when the includeAccessKey parameter is set.
func (m *VideoLibrary) SetApiAccessKey(value *string)() {
    m.apiAccessKey = value
}
// SetApiKey sets the ApiKey property value. The API key used for authenticating with the video library
func (m *VideoLibrary) SetApiKey(value *string)() {
    m.apiKey = value
}
// SetAppleFairPlayDrm sets the AppleFairPlayDrm property value. The AppleFairPlayDrm property
func (m *VideoLibrary) SetAppleFairPlayDrm(value VideoLibrary_AppleFairPlayDrmable)() {
    m.appleFairPlayDrm = value
}
// SetBitrate1080p sets the Bitrate1080p property value. The bitrate used for encoding 1080p videos
func (m *VideoLibrary) SetBitrate1080p(value *int32)() {
    m.bitrate1080p = value
}
// SetBitrate1440p sets the Bitrate1440p property value. The bitrate used for encoding 1440p videos
func (m *VideoLibrary) SetBitrate1440p(value *int32)() {
    m.bitrate1440p = value
}
// SetBitrate2160p sets the Bitrate2160p property value. The bitrate used for encoding 2160p videos
func (m *VideoLibrary) SetBitrate2160p(value *int32)() {
    m.bitrate2160p = value
}
// SetBitrate240p sets the Bitrate240p property value. The bitrate used for encoding 240p videos
func (m *VideoLibrary) SetBitrate240p(value *int32)() {
    m.bitrate240p = value
}
// SetBitrate360p sets the Bitrate360p property value. The bitrate used for encoding 360p videos
func (m *VideoLibrary) SetBitrate360p(value *int32)() {
    m.bitrate360p = value
}
// SetBitrate480p sets the Bitrate480p property value. The bitrate used for encoding 480p videos
func (m *VideoLibrary) SetBitrate480p(value *int32)() {
    m.bitrate480p = value
}
// SetBitrate720p sets the Bitrate720p property value. The bitrate used for encoding 720p videos
func (m *VideoLibrary) SetBitrate720p(value *int32)() {
    m.bitrate720p = value
}
// SetBlockedReferrers sets the BlockedReferrers property value. The list of blocked referrer domains blocked from accessing the library
func (m *VideoLibrary) SetBlockedReferrers(value []string)() {
    m.blockedReferrers = value
}
// SetBlockNoneReferrer sets the BlockNoneReferrer property value. Determines if the requests without a referrer are blocked
func (m *VideoLibrary) SetBlockNoneReferrer(value *bool)() {
    m.blockNoneReferrer = value
}
// SetCaptionsBackground sets the CaptionsBackground property value. The captions display background color
func (m *VideoLibrary) SetCaptionsBackground(value *string)() {
    m.captionsBackground = value
}
// SetCaptionsFontColor sets the CaptionsFontColor property value. The captions display font color
func (m *VideoLibrary) SetCaptionsFontColor(value *string)() {
    m.captionsFontColor = value
}
// SetCaptionsFontSize sets the CaptionsFontSize property value. The captions display font size
func (m *VideoLibrary) SetCaptionsFontSize(value *int32)() {
    m.captionsFontSize = value
}
// SetControls sets the Controls property value. The list of controls on the video player.
func (m *VideoLibrary) SetControls(value *string)() {
    m.controls = value
}
// SetCustomHTML sets the CustomHTML property value. The custom HTMl that is added into the head of the HTML player.
func (m *VideoLibrary) SetCustomHTML(value *string)() {
    m.customHTML = value
}
// SetDateCreated sets the DateCreated property value. The date when the video library was created
func (m *VideoLibrary) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateCreated = value
}
// SetDrmBasePriceOverride sets the DrmBasePriceOverride property value. The DrmBasePriceOverride property
func (m *VideoLibrary) SetDrmBasePriceOverride(value *float64)() {
    m.drmBasePriceOverride = value
}
// SetDrmCostPerLicenseOverride sets the DrmCostPerLicenseOverride property value. The DrmCostPerLicenseOverride property
func (m *VideoLibrary) SetDrmCostPerLicenseOverride(value *float64)() {
    m.drmCostPerLicenseOverride = value
}
// SetDrmVersion sets the DrmVersion property value. The DrmVersion property
func (m *VideoLibrary) SetDrmVersion(value *int64)() {
    m.drmVersion = value
}
// SetEnableContentTagging sets the EnableContentTagging property value. Determines if content tagging should be enabled for this library.
func (m *VideoLibrary) SetEnableContentTagging(value *bool)() {
    m.enableContentTagging = value
}
// SetEnabledResolutions sets the EnabledResolutions property value. The comma separated list of enabled resolutions
func (m *VideoLibrary) SetEnabledResolutions(value *string)() {
    m.enabledResolutions = value
}
// SetEnableDRM sets the EnableDRM property value. Determines if the MediaCage basic DRM is enabled
func (m *VideoLibrary) SetEnableDRM(value *bool)() {
    m.enableDRM = value
}
// SetEnableMP4Fallback sets the EnableMP4Fallback property value. Determines if the MP4 fallback feature is enabled
func (m *VideoLibrary) SetEnableMP4Fallback(value *bool)() {
    m.enableMP4Fallback = value
}
// SetEnableMultiAudioTrackSupport sets the EnableMultiAudioTrackSupport property value. The EnableMultiAudioTrackSupport property
func (m *VideoLibrary) SetEnableMultiAudioTrackSupport(value *bool)() {
    m.enableMultiAudioTrackSupport = value
}
// SetEnableTranscribing sets the EnableTranscribing property value. Determines if the automatic audio transcribing is currently enabled for this zone.
func (m *VideoLibrary) SetEnableTranscribing(value *bool)() {
    m.enableTranscribing = value
}
// SetEnableTranscribingDescriptionGeneration sets the EnableTranscribingDescriptionGeneration property value. Determines if automatic transcribing description generation is currently enabled.
func (m *VideoLibrary) SetEnableTranscribingDescriptionGeneration(value *bool)() {
    m.enableTranscribingDescriptionGeneration = value
}
// SetEnableTranscribingTitleGeneration sets the EnableTranscribingTitleGeneration property value. Determines if automatic transcribing title generation is currently enabled.
func (m *VideoLibrary) SetEnableTranscribingTitleGeneration(value *bool)() {
    m.enableTranscribingTitleGeneration = value
}
// SetEncodingTier sets the EncodingTier property value. The EncodingTier property
func (m *VideoLibrary) SetEncodingTier(value *int32)() {
    m.encodingTier = value
}
// SetFontFamily sets the FontFamily property value. The captions font family.
func (m *VideoLibrary) SetFontFamily(value *string)() {
    m.fontFamily = value
}
// SetGoogleWidevineDrm sets the GoogleWidevineDrm property value. The GoogleWidevineDrm property
func (m *VideoLibrary) SetGoogleWidevineDrm(value VideoLibrary_GoogleWidevineDrmable)() {
    m.googleWidevineDrm = value
}
// SetHasWatermark sets the HasWatermark property value. Determines if the video library has a watermark configured
func (m *VideoLibrary) SetHasWatermark(value *bool)() {
    m.hasWatermark = value
}
// SetId sets the Id property value. The Id property
func (m *VideoLibrary) SetId(value *int64)() {
    m.id = value
}
// SetJitEncodingEnabled sets the JitEncodingEnabled property value. The JitEncodingEnabled property
func (m *VideoLibrary) SetJitEncodingEnabled(value *bool)() {
    m.jitEncodingEnabled = value
}
// SetKeepOriginalFiles sets the KeepOriginalFiles property value. Determines if the original video files should be stored after encoding
func (m *VideoLibrary) SetKeepOriginalFiles(value *bool)() {
    m.keepOriginalFiles = value
}
// SetMonthlyChargesEnterpriseDrm sets the MonthlyChargesEnterpriseDrm property value. The MonthlyChargesEnterpriseDrm property
func (m *VideoLibrary) SetMonthlyChargesEnterpriseDrm(value *float64)() {
    m.monthlyChargesEnterpriseDrm = value
}
// SetMonthlyChargesPremiumEncoding sets the MonthlyChargesPremiumEncoding property value. The MonthlyChargesPremiumEncoding property
func (m *VideoLibrary) SetMonthlyChargesPremiumEncoding(value *float64)() {
    m.monthlyChargesPremiumEncoding = value
}
// SetMonthlyChargesTranscribing sets the MonthlyChargesTranscribing property value. The MonthlyChargesTranscribing property
func (m *VideoLibrary) SetMonthlyChargesTranscribing(value *float64)() {
    m.monthlyChargesTranscribing = value
}
// SetName sets the Name property value. The name of the Video Library.
func (m *VideoLibrary) SetName(value *string)() {
    m.name = value
}
// SetOutputCodecs sets the OutputCodecs property value. The OutputCodecs property
func (m *VideoLibrary) SetOutputCodecs(value *string)() {
    m.outputCodecs = value
}
// SetPlayerKeyColor sets the PlayerKeyColor property value. The key color of the player.
func (m *VideoLibrary) SetPlayerKeyColor(value *string)() {
    m.playerKeyColor = value
}
// SetPlayerTokenAuthenticationEnabled sets the PlayerTokenAuthenticationEnabled property value. Determines if the player token authentication is enabled
func (m *VideoLibrary) SetPlayerTokenAuthenticationEnabled(value *bool)() {
    m.playerTokenAuthenticationEnabled = value
}
// SetPremiumEncodingPriceOverride sets the PremiumEncodingPriceOverride property value. The PremiumEncodingPriceOverride property
func (m *VideoLibrary) SetPremiumEncodingPriceOverride(value *float64)() {
    m.premiumEncodingPriceOverride = value
}
// SetPullZoneId sets the PullZoneId property value. The ID of the connected underlying pull zone
func (m *VideoLibrary) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetPullZoneType sets the PullZoneType property value. The PullZoneType property
func (m *VideoLibrary) SetPullZoneType(value *float64)() {
    m.pullZoneType = value
}
// SetReadOnlyApiKey sets the ReadOnlyApiKey property value. The read-only API key used for authenticating with the video library
func (m *VideoLibrary) SetReadOnlyApiKey(value *string)() {
    m.readOnlyApiKey = value
}
// SetRememberPlayerPosition sets the RememberPlayerPosition property value. The list of languages that the captions will be automatically transcribed to.
func (m *VideoLibrary) SetRememberPlayerPosition(value *bool)() {
    m.rememberPlayerPosition = value
}
// SetReplicationRegions sets the ReplicationRegions property value. The geo-replication regions of the underlying storage zone
func (m *VideoLibrary) SetReplicationRegions(value []ReplicationRegions)() {
    m.replicationRegions = value
}
// SetShowHeatmap sets the ShowHeatmap property value. Determines if the video watch heatmap should be displayed in the player.
func (m *VideoLibrary) SetShowHeatmap(value *bool)() {
    m.showHeatmap = value
}
// SetStorageUsage sets the StorageUsage property value. The total amount of storage used by the library
func (m *VideoLibrary) SetStorageUsage(value *int64)() {
    m.storageUsage = value
}
// SetStorageZoneId sets the StorageZoneId property value. The ID of the connected underlying storage zone
func (m *VideoLibrary) SetStorageZoneId(value *int64)() {
    m.storageZoneId = value
}
// SetTrafficUsage sets the TrafficUsage property value. The amount of traffic usage this month
func (m *VideoLibrary) SetTrafficUsage(value *int64)() {
    m.trafficUsage = value
}
// SetTranscribingCaptionLanguages sets the TranscribingCaptionLanguages property value. The TranscribingCaptionLanguages property
func (m *VideoLibrary) SetTranscribingCaptionLanguages(value []string)() {
    m.transcribingCaptionLanguages = value
}
// SetTranscribingPriceOverride sets the TranscribingPriceOverride property value. The TranscribingPriceOverride property
func (m *VideoLibrary) SetTranscribingPriceOverride(value *float64)() {
    m.transcribingPriceOverride = value
}
// SetUILanguage sets the UILanguage property value. The UI language of the player
func (m *VideoLibrary) SetUILanguage(value *string)() {
    m.uILanguage = value
}
// SetUseSeparateAudioStream sets the UseSeparateAudioStream property value. The UseSeparateAudioStream property
func (m *VideoLibrary) SetUseSeparateAudioStream(value *bool)() {
    m.useSeparateAudioStream = value
}
// SetVastTagUrl sets the VastTagUrl property value. The URL of the VAST tag endpoint for advertising configuration
func (m *VideoLibrary) SetVastTagUrl(value *string)() {
    m.vastTagUrl = value
}
// SetViAiPublisherId sets the ViAiPublisherId property value. The vi.ai publisher id for advertising configuration
func (m *VideoLibrary) SetViAiPublisherId(value *string)() {
    m.viAiPublisherId = value
}
// SetVideoCount sets the VideoCount property value. The number of videos in the video library
func (m *VideoLibrary) SetVideoCount(value *int64)() {
    m.videoCount = value
}
// SetWatermarkHeight sets the WatermarkHeight property value. The height of the watermark (in %)
func (m *VideoLibrary) SetWatermarkHeight(value *int32)() {
    m.watermarkHeight = value
}
// SetWatermarkPositionLeft sets the WatermarkPositionLeft property value. The left offset of the watermark position (in %)
func (m *VideoLibrary) SetWatermarkPositionLeft(value *int32)() {
    m.watermarkPositionLeft = value
}
// SetWatermarkPositionTop sets the WatermarkPositionTop property value. The top offset of the watermark position (in %)
func (m *VideoLibrary) SetWatermarkPositionTop(value *int32)() {
    m.watermarkPositionTop = value
}
// SetWatermarkVersion sets the WatermarkVersion property value. The WatermarkVersion property
func (m *VideoLibrary) SetWatermarkVersion(value *int64)() {
    m.watermarkVersion = value
}
// SetWatermarkWidth sets the WatermarkWidth property value. The width of the watermark (in %)
func (m *VideoLibrary) SetWatermarkWidth(value *int32)() {
    m.watermarkWidth = value
}
// SetWebhookUrl sets the WebhookUrl property value. The webhook URL of the video library
func (m *VideoLibrary) SetWebhookUrl(value *string)() {
    m.webhookUrl = value
}
type VideoLibraryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDirectPlay()(*bool)
    GetAllowEarlyPlay()(*bool)
    GetAllowedReferrers()([]string)
    GetApiAccessKey()(*string)
    GetApiKey()(*string)
    GetAppleFairPlayDrm()(VideoLibrary_AppleFairPlayDrmable)
    GetBitrate1080p()(*int32)
    GetBitrate1440p()(*int32)
    GetBitrate2160p()(*int32)
    GetBitrate240p()(*int32)
    GetBitrate360p()(*int32)
    GetBitrate480p()(*int32)
    GetBitrate720p()(*int32)
    GetBlockedReferrers()([]string)
    GetBlockNoneReferrer()(*bool)
    GetCaptionsBackground()(*string)
    GetCaptionsFontColor()(*string)
    GetCaptionsFontSize()(*int32)
    GetControls()(*string)
    GetCustomHTML()(*string)
    GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDrmBasePriceOverride()(*float64)
    GetDrmCostPerLicenseOverride()(*float64)
    GetDrmVersion()(*int64)
    GetEnableContentTagging()(*bool)
    GetEnabledResolutions()(*string)
    GetEnableDRM()(*bool)
    GetEnableMP4Fallback()(*bool)
    GetEnableMultiAudioTrackSupport()(*bool)
    GetEnableTranscribing()(*bool)
    GetEnableTranscribingDescriptionGeneration()(*bool)
    GetEnableTranscribingTitleGeneration()(*bool)
    GetEncodingTier()(*int32)
    GetFontFamily()(*string)
    GetGoogleWidevineDrm()(VideoLibrary_GoogleWidevineDrmable)
    GetHasWatermark()(*bool)
    GetId()(*int64)
    GetJitEncodingEnabled()(*bool)
    GetKeepOriginalFiles()(*bool)
    GetMonthlyChargesEnterpriseDrm()(*float64)
    GetMonthlyChargesPremiumEncoding()(*float64)
    GetMonthlyChargesTranscribing()(*float64)
    GetName()(*string)
    GetOutputCodecs()(*string)
    GetPlayerKeyColor()(*string)
    GetPlayerTokenAuthenticationEnabled()(*bool)
    GetPremiumEncodingPriceOverride()(*float64)
    GetPullZoneId()(*int64)
    GetPullZoneType()(*float64)
    GetReadOnlyApiKey()(*string)
    GetRememberPlayerPosition()(*bool)
    GetReplicationRegions()([]ReplicationRegions)
    GetShowHeatmap()(*bool)
    GetStorageUsage()(*int64)
    GetStorageZoneId()(*int64)
    GetTrafficUsage()(*int64)
    GetTranscribingCaptionLanguages()([]string)
    GetTranscribingPriceOverride()(*float64)
    GetUILanguage()(*string)
    GetUseSeparateAudioStream()(*bool)
    GetVastTagUrl()(*string)
    GetViAiPublisherId()(*string)
    GetVideoCount()(*int64)
    GetWatermarkHeight()(*int32)
    GetWatermarkPositionLeft()(*int32)
    GetWatermarkPositionTop()(*int32)
    GetWatermarkVersion()(*int64)
    GetWatermarkWidth()(*int32)
    GetWebhookUrl()(*string)
    SetAllowDirectPlay(value *bool)()
    SetAllowEarlyPlay(value *bool)()
    SetAllowedReferrers(value []string)()
    SetApiAccessKey(value *string)()
    SetApiKey(value *string)()
    SetAppleFairPlayDrm(value VideoLibrary_AppleFairPlayDrmable)()
    SetBitrate1080p(value *int32)()
    SetBitrate1440p(value *int32)()
    SetBitrate2160p(value *int32)()
    SetBitrate240p(value *int32)()
    SetBitrate360p(value *int32)()
    SetBitrate480p(value *int32)()
    SetBitrate720p(value *int32)()
    SetBlockedReferrers(value []string)()
    SetBlockNoneReferrer(value *bool)()
    SetCaptionsBackground(value *string)()
    SetCaptionsFontColor(value *string)()
    SetCaptionsFontSize(value *int32)()
    SetControls(value *string)()
    SetCustomHTML(value *string)()
    SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDrmBasePriceOverride(value *float64)()
    SetDrmCostPerLicenseOverride(value *float64)()
    SetDrmVersion(value *int64)()
    SetEnableContentTagging(value *bool)()
    SetEnabledResolutions(value *string)()
    SetEnableDRM(value *bool)()
    SetEnableMP4Fallback(value *bool)()
    SetEnableMultiAudioTrackSupport(value *bool)()
    SetEnableTranscribing(value *bool)()
    SetEnableTranscribingDescriptionGeneration(value *bool)()
    SetEnableTranscribingTitleGeneration(value *bool)()
    SetEncodingTier(value *int32)()
    SetFontFamily(value *string)()
    SetGoogleWidevineDrm(value VideoLibrary_GoogleWidevineDrmable)()
    SetHasWatermark(value *bool)()
    SetId(value *int64)()
    SetJitEncodingEnabled(value *bool)()
    SetKeepOriginalFiles(value *bool)()
    SetMonthlyChargesEnterpriseDrm(value *float64)()
    SetMonthlyChargesPremiumEncoding(value *float64)()
    SetMonthlyChargesTranscribing(value *float64)()
    SetName(value *string)()
    SetOutputCodecs(value *string)()
    SetPlayerKeyColor(value *string)()
    SetPlayerTokenAuthenticationEnabled(value *bool)()
    SetPremiumEncodingPriceOverride(value *float64)()
    SetPullZoneId(value *int64)()
    SetPullZoneType(value *float64)()
    SetReadOnlyApiKey(value *string)()
    SetRememberPlayerPosition(value *bool)()
    SetReplicationRegions(value []ReplicationRegions)()
    SetShowHeatmap(value *bool)()
    SetStorageUsage(value *int64)()
    SetStorageZoneId(value *int64)()
    SetTrafficUsage(value *int64)()
    SetTranscribingCaptionLanguages(value []string)()
    SetTranscribingPriceOverride(value *float64)()
    SetUILanguage(value *string)()
    SetUseSeparateAudioStream(value *bool)()
    SetVastTagUrl(value *string)()
    SetViAiPublisherId(value *string)()
    SetVideoCount(value *int64)()
    SetWatermarkHeight(value *int32)()
    SetWatermarkPositionLeft(value *int32)()
    SetWatermarkPositionTop(value *int32)()
    SetWatermarkVersion(value *int64)()
    SetWatermarkWidth(value *int32)()
    SetWebhookUrl(value *string)()
}
