package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullZoneWafConfigVariableModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The valueEncoded property
    valueEncoded *string
}
// NewPullZoneWafConfigVariableModel instantiates a new PullZoneWafConfigVariableModel and sets the default values.
func NewPullZoneWafConfigVariableModel()(*PullZoneWafConfigVariableModel) {
    m := &PullZoneWafConfigVariableModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullZoneWafConfigVariableModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullZoneWafConfigVariableModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullZoneWafConfigVariableModel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullZoneWafConfigVariableModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullZoneWafConfigVariableModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["valueEncoded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueEncoded(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PullZoneWafConfigVariableModel) GetName()(*string) {
    return m.name
}
// GetValueEncoded gets the valueEncoded property value. The valueEncoded property
// returns a *string when successful
func (m *PullZoneWafConfigVariableModel) GetValueEncoded()(*string) {
    return m.valueEncoded
}
// Serialize serializes information the current object
func (m *PullZoneWafConfigVariableModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueEncoded", m.GetValueEncoded())
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
func (m *PullZoneWafConfigVariableModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *PullZoneWafConfigVariableModel) SetName(value *string)() {
    m.name = value
}
// SetValueEncoded sets the valueEncoded property value. The valueEncoded property
func (m *PullZoneWafConfigVariableModel) SetValueEncoded(value *string)() {
    m.valueEncoded = value
}
type PullZoneWafConfigVariableModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetValueEncoded()(*string)
    SetName(value *string)()
    SetValueEncoded(value *string)()
}
