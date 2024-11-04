package edgerule

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Action struct {
    // The ActionParameter1 property
    actionParameter1 *string
    // The ActionParameter2 property
    actionParameter2 *string
    // The ActionParameter3 property
    actionParameter3 *string
    // The ActionType property
    actionType *float64
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAction instantiates a new Action and sets the default values.
func NewAction()(*Action) {
    m := &Action{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAction(), nil
}
// GetActionParameter1 gets the ActionParameter1 property value. The ActionParameter1 property
// returns a *string when successful
func (m *Action) GetActionParameter1()(*string) {
    return m.actionParameter1
}
// GetActionParameter2 gets the ActionParameter2 property value. The ActionParameter2 property
// returns a *string when successful
func (m *Action) GetActionParameter2()(*string) {
    return m.actionParameter2
}
// GetActionParameter3 gets the ActionParameter3 property value. The ActionParameter3 property
// returns a *string when successful
func (m *Action) GetActionParameter3()(*string) {
    return m.actionParameter3
}
// GetActionType gets the ActionType property value. The ActionType property
// returns a *float64 when successful
func (m *Action) GetActionType()(*float64) {
    return m.actionType
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Action) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Action) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ActionParameter1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionParameter1(val)
        }
        return nil
    }
    res["ActionParameter2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionParameter2(val)
        }
        return nil
    }
    res["ActionParameter3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionParameter3(val)
        }
        return nil
    }
    res["ActionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Action) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ActionParameter1", m.GetActionParameter1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ActionParameter2", m.GetActionParameter2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ActionParameter3", m.GetActionParameter3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("ActionType", m.GetActionType())
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
// SetActionParameter1 sets the ActionParameter1 property value. The ActionParameter1 property
func (m *Action) SetActionParameter1(value *string)() {
    m.actionParameter1 = value
}
// SetActionParameter2 sets the ActionParameter2 property value. The ActionParameter2 property
func (m *Action) SetActionParameter2(value *string)() {
    m.actionParameter2 = value
}
// SetActionParameter3 sets the ActionParameter3 property value. The ActionParameter3 property
func (m *Action) SetActionParameter3(value *string)() {
    m.actionParameter3 = value
}
// SetActionType sets the ActionType property value. The ActionType property
func (m *Action) SetActionType(value *float64)() {
    m.actionType = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Action) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type Actionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionParameter1()(*string)
    GetActionParameter2()(*string)
    GetActionParameter3()(*string)
    GetActionType()(*float64)
    SetActionParameter1(value *string)()
    SetActionParameter2(value *string)()
    SetActionParameter3(value *string)()
    SetActionType(value *float64)()
}
