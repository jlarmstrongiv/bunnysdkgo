package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdgeScriptVariable struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DefaultValue property
    defaultValue *string
    // The Id property
    id *int64
    // The Name property
    name *string
    // The Required property
    required *bool
}
// NewEdgeScriptVariable instantiates a new EdgeScriptVariable and sets the default values.
func NewEdgeScriptVariable()(*EdgeScriptVariable) {
    m := &EdgeScriptVariable{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdgeScriptVariableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdgeScriptVariableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeScriptVariable(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EdgeScriptVariable) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultValue gets the DefaultValue property value. The DefaultValue property
// returns a *string when successful
func (m *EdgeScriptVariable) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdgeScriptVariable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["DefaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
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
    res["Required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *EdgeScriptVariable) GetId()(*int64) {
    return m.id
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *EdgeScriptVariable) GetName()(*string) {
    return m.name
}
// GetRequired gets the Required property value. The Required property
// returns a *bool when successful
func (m *EdgeScriptVariable) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *EdgeScriptVariable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("DefaultValue", m.GetDefaultValue())
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
        err := writer.WriteBoolValue("Required", m.GetRequired())
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
func (m *EdgeScriptVariable) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultValue sets the DefaultValue property value. The DefaultValue property
func (m *EdgeScriptVariable) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetId sets the Id property value. The Id property
func (m *EdgeScriptVariable) SetId(value *int64)() {
    m.id = value
}
// SetName sets the Name property value. The Name property
func (m *EdgeScriptVariable) SetName(value *string)() {
    m.name = value
}
// SetRequired sets the Required property value. The Required property
func (m *EdgeScriptVariable) SetRequired(value *bool)() {
    m.required = value
}
type EdgeScriptVariableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetId()(*int64)
    GetName()(*string)
    GetRequired()(*bool)
    SetDefaultValue(value *string)()
    SetId(value *int64)()
    SetName(value *string)()
    SetRequired(value *bool)()
}
