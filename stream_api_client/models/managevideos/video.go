package managevideos

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Video struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The available resolutions of the video
    availableResolutions *string
    // The average watch time of the video in seconds
    averageWatchTime *int64
    // The captions property
    captions []Captionable
    // The automatically detected category of the video
    category *Video_category
    // The list of chapters available for the video
    chapters []Chapterable
    // The ID of the collection where the video belongs
    collectionId *string
    // The date when the video was uploaded
    dateUploaded *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current encode progress of the video
    encodeProgress *int32
    // The framerate of the video
    framerate *float64
    // The unique ID of the video
    guid *string
    // Determines if the video has MP4 fallback files generated
    hasMP4Fallback *bool
    // The height of the original video file
    height *int32
    // Determines if the video is publicly accessible
    isPublic *bool
    // The duration of the video in seconds
    length *int32
    // The list of meta tags that have been added to the video
    metaTags []MetaTagable
    // The list of moments available for the video
    moments []Momentable
    // The rotation of the video
    rotation *int32
    // The status of the video.
    status *float64
    // The amount of storage used by this video
    storageSize *int64
    // The number of thumbnails generated for this video
    thumbnailCount *int32
    // The file name of the thumbnail inside of the storage
    thumbnailFileName *string
    // The title of the video
    title *string
    // The total video watch time in seconds
    totalWatchTime *int64
    // The ID of the video library that the video belongs to
    videoLibraryId *int64
    // The number of views the video received
    views *int64
    // The width of the original video file
    width *int32
}
// NewVideo instantiates a new Video and sets the default values.
func NewVideo()(*Video) {
    m := &Video{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Video) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableResolutions gets the availableResolutions property value. The available resolutions of the video
// returns a *string when successful
func (m *Video) GetAvailableResolutions()(*string) {
    return m.availableResolutions
}
// GetAverageWatchTime gets the averageWatchTime property value. The average watch time of the video in seconds
// returns a *int64 when successful
func (m *Video) GetAverageWatchTime()(*int64) {
    return m.averageWatchTime
}
// GetCaptions gets the captions property value. The captions property
// returns a []Captionable when successful
func (m *Video) GetCaptions()([]Captionable) {
    return m.captions
}
// GetCategory gets the category property value. The automatically detected category of the video
// returns a *Video_category when successful
func (m *Video) GetCategory()(*Video_category) {
    return m.category
}
// GetChapters gets the chapters property value. The list of chapters available for the video
// returns a []Chapterable when successful
func (m *Video) GetChapters()([]Chapterable) {
    return m.chapters
}
// GetCollectionId gets the collectionId property value. The ID of the collection where the video belongs
// returns a *string when successful
func (m *Video) GetCollectionId()(*string) {
    return m.collectionId
}
// GetDateUploaded gets the dateUploaded property value. The date when the video was uploaded
// returns a *Time when successful
func (m *Video) GetDateUploaded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateUploaded
}
// GetEncodeProgress gets the encodeProgress property value. The current encode progress of the video
// returns a *int32 when successful
func (m *Video) GetEncodeProgress()(*int32) {
    return m.encodeProgress
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Video) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableResolutions(val)
        }
        return nil
    }
    res["averageWatchTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageWatchTime(val)
        }
        return nil
    }
    res["captions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCaptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Captionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Captionable)
                }
            }
            m.SetCaptions(res)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVideo_category)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*Video_category))
        }
        return nil
    }
    res["chapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChapterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Chapterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Chapterable)
                }
            }
            m.SetChapters(res)
        }
        return nil
    }
    res["collectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionId(val)
        }
        return nil
    }
    res["dateUploaded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateUploaded(val)
        }
        return nil
    }
    res["encodeProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncodeProgress(val)
        }
        return nil
    }
    res["framerate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFramerate(val)
        }
        return nil
    }
    res["guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
        }
        return nil
    }
    res["hasMP4Fallback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMP4Fallback(val)
        }
        return nil
    }
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["isPublic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublic(val)
        }
        return nil
    }
    res["length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["metaTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaTagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaTagable)
                }
            }
            m.SetMetaTags(res)
        }
        return nil
    }
    res["moments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMomentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Momentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Momentable)
                }
            }
            m.SetMoments(res)
        }
        return nil
    }
    res["rotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRotation(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["storageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSize(val)
        }
        return nil
    }
    res["thumbnailCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailCount(val)
        }
        return nil
    }
    res["thumbnailFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailFileName(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["totalWatchTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWatchTime(val)
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
    res["views"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViews(val)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetFramerate gets the framerate property value. The framerate of the video
// returns a *float64 when successful
func (m *Video) GetFramerate()(*float64) {
    return m.framerate
}
// GetGuid gets the guid property value. The unique ID of the video
// returns a *string when successful
func (m *Video) GetGuid()(*string) {
    return m.guid
}
// GetHasMP4Fallback gets the hasMP4Fallback property value. Determines if the video has MP4 fallback files generated
// returns a *bool when successful
func (m *Video) GetHasMP4Fallback()(*bool) {
    return m.hasMP4Fallback
}
// GetHeight gets the height property value. The height of the original video file
// returns a *int32 when successful
func (m *Video) GetHeight()(*int32) {
    return m.height
}
// GetIsPublic gets the isPublic property value. Determines if the video is publicly accessible
// returns a *bool when successful
func (m *Video) GetIsPublic()(*bool) {
    return m.isPublic
}
// GetLength gets the length property value. The duration of the video in seconds
// returns a *int32 when successful
func (m *Video) GetLength()(*int32) {
    return m.length
}
// GetMetaTags gets the metaTags property value. The list of meta tags that have been added to the video
// returns a []MetaTagable when successful
func (m *Video) GetMetaTags()([]MetaTagable) {
    return m.metaTags
}
// GetMoments gets the moments property value. The list of moments available for the video
// returns a []Momentable when successful
func (m *Video) GetMoments()([]Momentable) {
    return m.moments
}
// GetRotation gets the rotation property value. The rotation of the video
// returns a *int32 when successful
func (m *Video) GetRotation()(*int32) {
    return m.rotation
}
// GetStatus gets the status property value. The status of the video.
// returns a *float64 when successful
func (m *Video) GetStatus()(*float64) {
    return m.status
}
// GetStorageSize gets the storageSize property value. The amount of storage used by this video
// returns a *int64 when successful
func (m *Video) GetStorageSize()(*int64) {
    return m.storageSize
}
// GetThumbnailCount gets the thumbnailCount property value. The number of thumbnails generated for this video
// returns a *int32 when successful
func (m *Video) GetThumbnailCount()(*int32) {
    return m.thumbnailCount
}
// GetThumbnailFileName gets the thumbnailFileName property value. The file name of the thumbnail inside of the storage
// returns a *string when successful
func (m *Video) GetThumbnailFileName()(*string) {
    return m.thumbnailFileName
}
// GetTitle gets the title property value. The title of the video
// returns a *string when successful
func (m *Video) GetTitle()(*string) {
    return m.title
}
// GetTotalWatchTime gets the totalWatchTime property value. The total video watch time in seconds
// returns a *int64 when successful
func (m *Video) GetTotalWatchTime()(*int64) {
    return m.totalWatchTime
}
// GetVideoLibraryId gets the videoLibraryId property value. The ID of the video library that the video belongs to
// returns a *int64 when successful
func (m *Video) GetVideoLibraryId()(*int64) {
    return m.videoLibraryId
}
// GetViews gets the views property value. The number of views the video received
// returns a *int64 when successful
func (m *Video) GetViews()(*int64) {
    return m.views
}
// GetWidth gets the width property value. The width of the original video file
// returns a *int32 when successful
func (m *Video) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Video) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCaptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCaptions()))
        for i, v := range m.GetCaptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("captions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChapters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChapters()))
        for i, v := range m.GetChapters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("chapters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("collectionId", m.GetCollectionId())
        if err != nil {
            return err
        }
    }
    if m.GetMetaTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetaTags()))
        for i, v := range m.GetMetaTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("metaTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMoments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMoments()))
        for i, v := range m.GetMoments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("moments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *Video) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableResolutions sets the availableResolutions property value. The available resolutions of the video
func (m *Video) SetAvailableResolutions(value *string)() {
    m.availableResolutions = value
}
// SetAverageWatchTime sets the averageWatchTime property value. The average watch time of the video in seconds
func (m *Video) SetAverageWatchTime(value *int64)() {
    m.averageWatchTime = value
}
// SetCaptions sets the captions property value. The captions property
func (m *Video) SetCaptions(value []Captionable)() {
    m.captions = value
}
// SetCategory sets the category property value. The automatically detected category of the video
func (m *Video) SetCategory(value *Video_category)() {
    m.category = value
}
// SetChapters sets the chapters property value. The list of chapters available for the video
func (m *Video) SetChapters(value []Chapterable)() {
    m.chapters = value
}
// SetCollectionId sets the collectionId property value. The ID of the collection where the video belongs
func (m *Video) SetCollectionId(value *string)() {
    m.collectionId = value
}
// SetDateUploaded sets the dateUploaded property value. The date when the video was uploaded
func (m *Video) SetDateUploaded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateUploaded = value
}
// SetEncodeProgress sets the encodeProgress property value. The current encode progress of the video
func (m *Video) SetEncodeProgress(value *int32)() {
    m.encodeProgress = value
}
// SetFramerate sets the framerate property value. The framerate of the video
func (m *Video) SetFramerate(value *float64)() {
    m.framerate = value
}
// SetGuid sets the guid property value. The unique ID of the video
func (m *Video) SetGuid(value *string)() {
    m.guid = value
}
// SetHasMP4Fallback sets the hasMP4Fallback property value. Determines if the video has MP4 fallback files generated
func (m *Video) SetHasMP4Fallback(value *bool)() {
    m.hasMP4Fallback = value
}
// SetHeight sets the height property value. The height of the original video file
func (m *Video) SetHeight(value *int32)() {
    m.height = value
}
// SetIsPublic sets the isPublic property value. Determines if the video is publicly accessible
func (m *Video) SetIsPublic(value *bool)() {
    m.isPublic = value
}
// SetLength sets the length property value. The duration of the video in seconds
func (m *Video) SetLength(value *int32)() {
    m.length = value
}
// SetMetaTags sets the metaTags property value. The list of meta tags that have been added to the video
func (m *Video) SetMetaTags(value []MetaTagable)() {
    m.metaTags = value
}
// SetMoments sets the moments property value. The list of moments available for the video
func (m *Video) SetMoments(value []Momentable)() {
    m.moments = value
}
// SetRotation sets the rotation property value. The rotation of the video
func (m *Video) SetRotation(value *int32)() {
    m.rotation = value
}
// SetStatus sets the status property value. The status of the video.
func (m *Video) SetStatus(value *float64)() {
    m.status = value
}
// SetStorageSize sets the storageSize property value. The amount of storage used by this video
func (m *Video) SetStorageSize(value *int64)() {
    m.storageSize = value
}
// SetThumbnailCount sets the thumbnailCount property value. The number of thumbnails generated for this video
func (m *Video) SetThumbnailCount(value *int32)() {
    m.thumbnailCount = value
}
// SetThumbnailFileName sets the thumbnailFileName property value. The file name of the thumbnail inside of the storage
func (m *Video) SetThumbnailFileName(value *string)() {
    m.thumbnailFileName = value
}
// SetTitle sets the title property value. The title of the video
func (m *Video) SetTitle(value *string)() {
    m.title = value
}
// SetTotalWatchTime sets the totalWatchTime property value. The total video watch time in seconds
func (m *Video) SetTotalWatchTime(value *int64)() {
    m.totalWatchTime = value
}
// SetVideoLibraryId sets the videoLibraryId property value. The ID of the video library that the video belongs to
func (m *Video) SetVideoLibraryId(value *int64)() {
    m.videoLibraryId = value
}
// SetViews sets the views property value. The number of views the video received
func (m *Video) SetViews(value *int64)() {
    m.views = value
}
// SetWidth sets the width property value. The width of the original video file
func (m *Video) SetWidth(value *int32)() {
    m.width = value
}
type Videoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableResolutions()(*string)
    GetAverageWatchTime()(*int64)
    GetCaptions()([]Captionable)
    GetCategory()(*Video_category)
    GetChapters()([]Chapterable)
    GetCollectionId()(*string)
    GetDateUploaded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEncodeProgress()(*int32)
    GetFramerate()(*float64)
    GetGuid()(*string)
    GetHasMP4Fallback()(*bool)
    GetHeight()(*int32)
    GetIsPublic()(*bool)
    GetLength()(*int32)
    GetMetaTags()([]MetaTagable)
    GetMoments()([]Momentable)
    GetRotation()(*int32)
    GetStatus()(*float64)
    GetStorageSize()(*int64)
    GetThumbnailCount()(*int32)
    GetThumbnailFileName()(*string)
    GetTitle()(*string)
    GetTotalWatchTime()(*int64)
    GetVideoLibraryId()(*int64)
    GetViews()(*int64)
    GetWidth()(*int32)
    SetAvailableResolutions(value *string)()
    SetAverageWatchTime(value *int64)()
    SetCaptions(value []Captionable)()
    SetCategory(value *Video_category)()
    SetChapters(value []Chapterable)()
    SetCollectionId(value *string)()
    SetDateUploaded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEncodeProgress(value *int32)()
    SetFramerate(value *float64)()
    SetGuid(value *string)()
    SetHasMP4Fallback(value *bool)()
    SetHeight(value *int32)()
    SetIsPublic(value *bool)()
    SetLength(value *int32)()
    SetMetaTags(value []MetaTagable)()
    SetMoments(value []Momentable)()
    SetRotation(value *int32)()
    SetStatus(value *float64)()
    SetStorageSize(value *int64)()
    SetThumbnailCount(value *int32)()
    SetThumbnailFileName(value *string)()
    SetTitle(value *string)()
    SetTotalWatchTime(value *int64)()
    SetVideoLibraryId(value *int64)()
    SetViews(value *int64)()
    SetWidth(value *int32)()
}
