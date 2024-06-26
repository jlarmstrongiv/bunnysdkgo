package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemRecordsPutRequestBody_EnviromentalVariables struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Name property
    name *string
    // The Value property
    value *string
}
// NewItemRecordsPutRequestBody_EnviromentalVariables instantiates a new ItemRecordsPutRequestBody_EnviromentalVariables and sets the default values.
func NewItemRecordsPutRequestBody_EnviromentalVariables()(*ItemRecordsPutRequestBody_EnviromentalVariables) {
    m := &ItemRecordsPutRequestBody_EnviromentalVariables{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemRecordsPutRequestBody_EnviromentalVariablesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRecordsPutRequestBody_EnviromentalVariablesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRecordsPutRequestBody_EnviromentalVariables(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) GetName()(*string) {
    return m.name
}
// GetValue gets the Value property value. The Value property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Value", m.GetValue())
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
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the Name property value. The Name property
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) SetName(value *string)() {
    m.name = value
}
// SetValue sets the Value property value. The Value property
func (m *ItemRecordsPutRequestBody_EnviromentalVariables) SetValue(value *string)() {
    m.value = value
}
type ItemRecordsPutRequestBody_EnviromentalVariablesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetValue(value *string)()
}
