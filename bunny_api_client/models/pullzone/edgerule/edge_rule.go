package edgerule

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdgeRule struct {
    // The Action parameter 1. The value depends on other parameters of the edge rule.
    actionParameter1 *string
    // The Action parameter 2. The value depends on other parameters of the edge rule.
    actionParameter2 *string
    // The ActionType property
    actionType *float64
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the edge rule
    description *string
    // Determines if the edge rule is currently enabled or not
    enabled *bool
    // The ExtraActions property
    extraActions []Actionable
    // The unique GUID of the edge rule
    guid *string
    // The TriggerMatchingType property
    triggerMatchingType *float64
    // The Triggers property
    triggers []Triggerable
}
// NewEdgeRule instantiates a new EdgeRule and sets the default values.
func NewEdgeRule()(*EdgeRule) {
    m := &EdgeRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdgeRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdgeRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeRule(), nil
}
// GetActionParameter1 gets the ActionParameter1 property value. The Action parameter 1. The value depends on other parameters of the edge rule.
// returns a *string when successful
func (m *EdgeRule) GetActionParameter1()(*string) {
    return m.actionParameter1
}
// GetActionParameter2 gets the ActionParameter2 property value. The Action parameter 2. The value depends on other parameters of the edge rule.
// returns a *string when successful
func (m *EdgeRule) GetActionParameter2()(*string) {
    return m.actionParameter2
}
// GetActionType gets the ActionType property value. The ActionType property
// returns a *float64 when successful
func (m *EdgeRule) GetActionType()(*float64) {
    return m.actionType
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EdgeRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the Description property value. The description of the edge rule
// returns a *string when successful
func (m *EdgeRule) GetDescription()(*string) {
    return m.description
}
// GetEnabled gets the Enabled property value. Determines if the edge rule is currently enabled or not
// returns a *bool when successful
func (m *EdgeRule) GetEnabled()(*bool) {
    return m.enabled
}
// GetExtraActions gets the ExtraActions property value. The ExtraActions property
// returns a []Actionable when successful
func (m *EdgeRule) GetExtraActions()([]Actionable) {
    return m.extraActions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdgeRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["Description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["Enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["ExtraActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Actionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Actionable)
                }
            }
            m.SetExtraActions(res)
        }
        return nil
    }
    res["Guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
        }
        return nil
    }
    res["TriggerMatchingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTriggerMatchingType(val)
        }
        return nil
    }
    res["Triggers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Triggerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Triggerable)
                }
            }
            m.SetTriggers(res)
        }
        return nil
    }
    return res
}
// GetGuid gets the Guid property value. The unique GUID of the edge rule
// returns a *string when successful
func (m *EdgeRule) GetGuid()(*string) {
    return m.guid
}
// GetTriggerMatchingType gets the TriggerMatchingType property value. The TriggerMatchingType property
// returns a *float64 when successful
func (m *EdgeRule) GetTriggerMatchingType()(*float64) {
    return m.triggerMatchingType
}
// GetTriggers gets the Triggers property value. The Triggers property
// returns a []Triggerable when successful
func (m *EdgeRule) GetTriggers()([]Triggerable) {
    return m.triggers
}
// Serialize serializes information the current object
func (m *EdgeRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteFloat64Value("ActionType", m.GetActionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetExtraActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtraActions()))
        for i, v := range m.GetExtraActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ExtraActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Guid", m.GetGuid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("TriggerMatchingType", m.GetTriggerMatchingType())
        if err != nil {
            return err
        }
    }
    if m.GetTriggers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTriggers()))
        for i, v := range m.GetTriggers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Triggers", cast)
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
// SetActionParameter1 sets the ActionParameter1 property value. The Action parameter 1. The value depends on other parameters of the edge rule.
func (m *EdgeRule) SetActionParameter1(value *string)() {
    m.actionParameter1 = value
}
// SetActionParameter2 sets the ActionParameter2 property value. The Action parameter 2. The value depends on other parameters of the edge rule.
func (m *EdgeRule) SetActionParameter2(value *string)() {
    m.actionParameter2 = value
}
// SetActionType sets the ActionType property value. The ActionType property
func (m *EdgeRule) SetActionType(value *float64)() {
    m.actionType = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdgeRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the Description property value. The description of the edge rule
func (m *EdgeRule) SetDescription(value *string)() {
    m.description = value
}
// SetEnabled sets the Enabled property value. Determines if the edge rule is currently enabled or not
func (m *EdgeRule) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetExtraActions sets the ExtraActions property value. The ExtraActions property
func (m *EdgeRule) SetExtraActions(value []Actionable)() {
    m.extraActions = value
}
// SetGuid sets the Guid property value. The unique GUID of the edge rule
func (m *EdgeRule) SetGuid(value *string)() {
    m.guid = value
}
// SetTriggerMatchingType sets the TriggerMatchingType property value. The TriggerMatchingType property
func (m *EdgeRule) SetTriggerMatchingType(value *float64)() {
    m.triggerMatchingType = value
}
// SetTriggers sets the Triggers property value. The Triggers property
func (m *EdgeRule) SetTriggers(value []Triggerable)() {
    m.triggers = value
}
type EdgeRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionParameter1()(*string)
    GetActionParameter2()(*string)
    GetActionType()(*float64)
    GetDescription()(*string)
    GetEnabled()(*bool)
    GetExtraActions()([]Actionable)
    GetGuid()(*string)
    GetTriggerMatchingType()(*float64)
    GetTriggers()([]Triggerable)
    SetActionParameter1(value *string)()
    SetActionParameter2(value *string)()
    SetActionType(value *float64)()
    SetDescription(value *string)()
    SetEnabled(value *bool)()
    SetExtraActions(value []Actionable)()
    SetGuid(value *string)()
    SetTriggerMatchingType(value *float64)()
    SetTriggers(value []Triggerable)()
}
