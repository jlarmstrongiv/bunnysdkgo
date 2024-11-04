package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScriptCreate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Code property
    code *string
    // The CreateLinkedPullZone property
    createLinkedPullZone *bool
    // The Integration property
    integration Integrationable
    // The Name property
    name *string
    // The ScriptType property
    scriptType *float64
}
// NewScriptCreate instantiates a new ScriptCreate and sets the default values.
func NewScriptCreate()(*ScriptCreate) {
    m := &ScriptCreate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScriptCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScriptCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScriptCreate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScriptCreate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the Code property value. The Code property
// returns a *string when successful
func (m *ScriptCreate) GetCode()(*string) {
    return m.code
}
// GetCreateLinkedPullZone gets the CreateLinkedPullZone property value. The CreateLinkedPullZone property
// returns a *bool when successful
func (m *ScriptCreate) GetCreateLinkedPullZone()(*bool) {
    return m.createLinkedPullZone
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScriptCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["CreateLinkedPullZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateLinkedPullZone(val)
        }
        return nil
    }
    res["Integration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegration(val.(Integrationable))
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
    res["ScriptType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptType(val)
        }
        return nil
    }
    return res
}
// GetIntegration gets the Integration property value. The Integration property
// returns a Integrationable when successful
func (m *ScriptCreate) GetIntegration()(Integrationable) {
    return m.integration
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *ScriptCreate) GetName()(*string) {
    return m.name
}
// GetScriptType gets the ScriptType property value. The ScriptType property
// returns a *float64 when successful
func (m *ScriptCreate) GetScriptType()(*float64) {
    return m.scriptType
}
// Serialize serializes information the current object
func (m *ScriptCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("CreateLinkedPullZone", m.GetCreateLinkedPullZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("Integration", m.GetIntegration())
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
        err := writer.WriteFloat64Value("ScriptType", m.GetScriptType())
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
func (m *ScriptCreate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the Code property value. The Code property
func (m *ScriptCreate) SetCode(value *string)() {
    m.code = value
}
// SetCreateLinkedPullZone sets the CreateLinkedPullZone property value. The CreateLinkedPullZone property
func (m *ScriptCreate) SetCreateLinkedPullZone(value *bool)() {
    m.createLinkedPullZone = value
}
// SetIntegration sets the Integration property value. The Integration property
func (m *ScriptCreate) SetIntegration(value Integrationable)() {
    m.integration = value
}
// SetName sets the Name property value. The Name property
func (m *ScriptCreate) SetName(value *string)() {
    m.name = value
}
// SetScriptType sets the ScriptType property value. The ScriptType property
func (m *ScriptCreate) SetScriptType(value *float64)() {
    m.scriptType = value
}
type ScriptCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetCreateLinkedPullZone()(*bool)
    GetIntegration()(Integrationable)
    GetName()(*string)
    GetScriptType()(*float64)
    SetCode(value *string)()
    SetCreateLinkedPullZone(value *bool)()
    SetIntegration(value Integrationable)()
    SetName(value *string)()
    SetScriptType(value *float64)()
}
