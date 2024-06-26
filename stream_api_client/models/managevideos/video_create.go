package managevideos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VideoCreate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the collection where the video belongs
    collectionId *string
    // Video time in ms to extract the main video thumbnail.
    thumbnailTime *int32
    // The title of the video
    title *string
}
// NewVideoCreate instantiates a new VideoCreate and sets the default values.
func NewVideoCreate()(*VideoCreate) {
    m := &VideoCreate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideoCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideoCreate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VideoCreate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCollectionId gets the collectionId property value. The ID of the collection where the video belongs
// returns a *string when successful
func (m *VideoCreate) GetCollectionId()(*string) {
    return m.collectionId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideoCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["thumbnailTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailTime(val)
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
    return res
}
// GetThumbnailTime gets the thumbnailTime property value. Video time in ms to extract the main video thumbnail.
// returns a *int32 when successful
func (m *VideoCreate) GetThumbnailTime()(*int32) {
    return m.thumbnailTime
}
// GetTitle gets the title property value. The title of the video
// returns a *string when successful
func (m *VideoCreate) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *VideoCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("collectionId", m.GetCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("thumbnailTime", m.GetThumbnailTime())
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
func (m *VideoCreate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCollectionId sets the collectionId property value. The ID of the collection where the video belongs
func (m *VideoCreate) SetCollectionId(value *string)() {
    m.collectionId = value
}
// SetThumbnailTime sets the thumbnailTime property value. Video time in ms to extract the main video thumbnail.
func (m *VideoCreate) SetThumbnailTime(value *int32)() {
    m.thumbnailTime = value
}
// SetTitle sets the title property value. The title of the video
func (m *VideoCreate) SetTitle(value *string)() {
    m.title = value
}
type VideoCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollectionId()(*string)
    GetThumbnailTime()(*int32)
    GetTitle()(*string)
    SetCollectionId(value *string)()
    SetThumbnailTime(value *int32)()
    SetTitle(value *string)()
}
