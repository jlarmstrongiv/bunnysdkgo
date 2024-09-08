package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScriptItemPublishItemWithUuPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Note property
    note *string
}
// NewScriptItemPublishItemWithUuPostRequestBody instantiates a new ScriptItemPublishItemWithUuPostRequestBody and sets the default values.
func NewScriptItemPublishItemWithUuPostRequestBody()(*ScriptItemPublishItemWithUuPostRequestBody) {
    m := &ScriptItemPublishItemWithUuPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScriptItemPublishItemWithUuPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScriptItemPublishItemWithUuPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScriptItemPublishItemWithUuPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScriptItemPublishItemWithUuPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScriptItemPublishItemWithUuPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    return res
}
// GetNote gets the Note property value. The Note property
// returns a *string when successful
func (m *ScriptItemPublishItemWithUuPostRequestBody) GetNote()(*string) {
    return m.note
}
// Serialize serializes information the current object
func (m *ScriptItemPublishItemWithUuPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Note", m.GetNote())
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
func (m *ScriptItemPublishItemWithUuPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNote sets the Note property value. The Note property
func (m *ScriptItemPublishItemWithUuPostRequestBody) SetNote(value *string)() {
    m.note = value
}
type ScriptItemPublishItemWithUuPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNote()(*string)
    SetNote(value *string)()
}
