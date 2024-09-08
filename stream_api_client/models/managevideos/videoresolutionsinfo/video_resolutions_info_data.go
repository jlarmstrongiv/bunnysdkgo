package videoresolutionsinfo

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VideoResolutionsInfoData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availableResolutions property
    availableResolutions []string
    // The configuredResolutions property
    configuredResolutions []string
    // The hasOriginal property
    hasOriginal *bool
    // The mp4Resolutions property
    mp4Resolutions Resolutionsable
    // The playlistResolutions property
    playlistResolutions Resolutionsable
    // The storageResolutions property
    storageResolutions Resolutionsable
    // The videoId property
    videoId *string
    // The videoLibraryId property
    videoLibraryId *int64
}
// NewVideoResolutionsInfoData instantiates a new VideoResolutionsInfoData and sets the default values.
func NewVideoResolutionsInfoData()(*VideoResolutionsInfoData) {
    m := &VideoResolutionsInfoData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideoResolutionsInfoDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoResolutionsInfoDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideoResolutionsInfoData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VideoResolutionsInfoData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableResolutions gets the availableResolutions property value. The availableResolutions property
// returns a []string when successful
func (m *VideoResolutionsInfoData) GetAvailableResolutions()([]string) {
    return m.availableResolutions
}
// GetConfiguredResolutions gets the configuredResolutions property value. The configuredResolutions property
// returns a []string when successful
func (m *VideoResolutionsInfoData) GetConfiguredResolutions()([]string) {
    return m.configuredResolutions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideoResolutionsInfoData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAvailableResolutions(res)
        }
        return nil
    }
    res["configuredResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetConfiguredResolutions(res)
        }
        return nil
    }
    res["hasOriginal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOriginal(val)
        }
        return nil
    }
    res["mp4Resolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResolutionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMp4Resolutions(val.(Resolutionsable))
        }
        return nil
    }
    res["playlistResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResolutionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaylistResolutions(val.(Resolutionsable))
        }
        return nil
    }
    res["storageResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResolutionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageResolutions(val.(Resolutionsable))
        }
        return nil
    }
    res["videoId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoId(val)
        }
        return nil
    }
    res["videoLibraryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoLibraryId(val)
        }
        return nil
    }
    return res
}
// GetHasOriginal gets the hasOriginal property value. The hasOriginal property
// returns a *bool when successful
func (m *VideoResolutionsInfoData) GetHasOriginal()(*bool) {
    return m.hasOriginal
}
// GetMp4Resolutions gets the mp4Resolutions property value. The mp4Resolutions property
// returns a Resolutionsable when successful
func (m *VideoResolutionsInfoData) GetMp4Resolutions()(Resolutionsable) {
    return m.mp4Resolutions
}
// GetPlaylistResolutions gets the playlistResolutions property value. The playlistResolutions property
// returns a Resolutionsable when successful
func (m *VideoResolutionsInfoData) GetPlaylistResolutions()(Resolutionsable) {
    return m.playlistResolutions
}
// GetStorageResolutions gets the storageResolutions property value. The storageResolutions property
// returns a Resolutionsable when successful
func (m *VideoResolutionsInfoData) GetStorageResolutions()(Resolutionsable) {
    return m.storageResolutions
}
// GetVideoId gets the videoId property value. The videoId property
// returns a *string when successful
func (m *VideoResolutionsInfoData) GetVideoId()(*string) {
    return m.videoId
}
// GetVideoLibraryId gets the videoLibraryId property value. The videoLibraryId property
// returns a *int64 when successful
func (m *VideoResolutionsInfoData) GetVideoLibraryId()(*int64) {
    return m.videoLibraryId
}
// Serialize serializes information the current object
func (m *VideoResolutionsInfoData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableResolutions() != nil {
        err := writer.WriteCollectionOfStringValues("availableResolutions", m.GetAvailableResolutions())
        if err != nil {
            return err
        }
    }
    if m.GetConfiguredResolutions() != nil {
        err := writer.WriteCollectionOfStringValues("configuredResolutions", m.GetConfiguredResolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasOriginal", m.GetHasOriginal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mp4Resolutions", m.GetMp4Resolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("playlistResolutions", m.GetPlaylistResolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageResolutions", m.GetStorageResolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("videoId", m.GetVideoId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("videoLibraryId", m.GetVideoLibraryId())
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
func (m *VideoResolutionsInfoData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableResolutions sets the availableResolutions property value. The availableResolutions property
func (m *VideoResolutionsInfoData) SetAvailableResolutions(value []string)() {
    m.availableResolutions = value
}
// SetConfiguredResolutions sets the configuredResolutions property value. The configuredResolutions property
func (m *VideoResolutionsInfoData) SetConfiguredResolutions(value []string)() {
    m.configuredResolutions = value
}
// SetHasOriginal sets the hasOriginal property value. The hasOriginal property
func (m *VideoResolutionsInfoData) SetHasOriginal(value *bool)() {
    m.hasOriginal = value
}
// SetMp4Resolutions sets the mp4Resolutions property value. The mp4Resolutions property
func (m *VideoResolutionsInfoData) SetMp4Resolutions(value Resolutionsable)() {
    m.mp4Resolutions = value
}
// SetPlaylistResolutions sets the playlistResolutions property value. The playlistResolutions property
func (m *VideoResolutionsInfoData) SetPlaylistResolutions(value Resolutionsable)() {
    m.playlistResolutions = value
}
// SetStorageResolutions sets the storageResolutions property value. The storageResolutions property
func (m *VideoResolutionsInfoData) SetStorageResolutions(value Resolutionsable)() {
    m.storageResolutions = value
}
// SetVideoId sets the videoId property value. The videoId property
func (m *VideoResolutionsInfoData) SetVideoId(value *string)() {
    m.videoId = value
}
// SetVideoLibraryId sets the videoLibraryId property value. The videoLibraryId property
func (m *VideoResolutionsInfoData) SetVideoLibraryId(value *int64)() {
    m.videoLibraryId = value
}
type VideoResolutionsInfoDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableResolutions()([]string)
    GetConfiguredResolutions()([]string)
    GetHasOriginal()(*bool)
    GetMp4Resolutions()(Resolutionsable)
    GetPlaylistResolutions()(Resolutionsable)
    GetStorageResolutions()(Resolutionsable)
    GetVideoId()(*string)
    GetVideoLibraryId()(*int64)
    SetAvailableResolutions(value []string)()
    SetConfiguredResolutions(value []string)()
    SetHasOriginal(value *bool)()
    SetMp4Resolutions(value Resolutionsable)()
    SetPlaylistResolutions(value Resolutionsable)()
    SetStorageResolutions(value Resolutionsable)()
    SetVideoId(value *string)()
    SetVideoLibraryId(value *int64)()
}
