package library

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Base64 encoded captions file
    captionsFile *string
    // The text description label for the caption
    label *string
    // The unique srclang shortcode for the caption
    srclang *string
}
// NewItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody instantiates a new ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody and sets the default values.
func NewItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody()(*ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) {
    m := &ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCaptionsFile gets the captionsFile property value. Base64 encoded captions file
// returns a *string when successful
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) GetCaptionsFile()(*string) {
    return m.captionsFile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["captionsFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptionsFile(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["srclang"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSrclang(val)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. The text description label for the caption
// returns a *string when successful
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) GetLabel()(*string) {
    return m.label
}
// GetSrclang gets the srclang property value. The unique srclang shortcode for the caption
// returns a *string when successful
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) GetSrclang()(*string) {
    return m.srclang
}
// Serialize serializes information the current object
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("captionsFile", m.GetCaptionsFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("srclang", m.GetSrclang())
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
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCaptionsFile sets the captionsFile property value. Base64 encoded captions file
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) SetCaptionsFile(value *string)() {
    m.captionsFile = value
}
// SetLabel sets the label property value. The text description label for the caption
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) SetLabel(value *string)() {
    m.label = value
}
// SetSrclang sets the srclang property value. The unique srclang shortcode for the caption
func (m *ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBody) SetSrclang(value *string)() {
    m.srclang = value
}
type ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCaptionsFile()(*string)
    GetLabel()(*string)
    GetSrclang()(*string)
    SetCaptionsFile(value *string)()
    SetLabel(value *string)()
    SetSrclang(value *string)()
}
