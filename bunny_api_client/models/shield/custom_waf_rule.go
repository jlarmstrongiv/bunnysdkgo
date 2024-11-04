package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomWafRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int64
    // The ruleConfiguration property
    ruleConfiguration CreateCustomWafRuleModelable
    // The ruleDescription property
    ruleDescription *string
    // The ruleJson property
    ruleJson *string
    // The ruleName property
    ruleName *string
    // The shieldZoneId property
    shieldZoneId *int64
    // The userId property
    userId *string
}
// NewCustomWafRule instantiates a new CustomWafRule and sets the default values.
func NewCustomWafRule()(*CustomWafRule) {
    m := &CustomWafRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomWafRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomWafRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomWafRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomWafRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomWafRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["ruleConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreateCustomWafRuleModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleConfiguration(val.(CreateCustomWafRuleModelable))
        }
        return nil
    }
    res["ruleDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleDescription(val)
        }
        return nil
    }
    res["ruleJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleJson(val)
        }
        return nil
    }
    res["ruleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    res["shieldZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldZoneId(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CustomWafRule) GetId()(*int64) {
    return m.id
}
// GetRuleConfiguration gets the ruleConfiguration property value. The ruleConfiguration property
// returns a CreateCustomWafRuleModelable when successful
func (m *CustomWafRule) GetRuleConfiguration()(CreateCustomWafRuleModelable) {
    return m.ruleConfiguration
}
// GetRuleDescription gets the ruleDescription property value. The ruleDescription property
// returns a *string when successful
func (m *CustomWafRule) GetRuleDescription()(*string) {
    return m.ruleDescription
}
// GetRuleJson gets the ruleJson property value. The ruleJson property
// returns a *string when successful
func (m *CustomWafRule) GetRuleJson()(*string) {
    return m.ruleJson
}
// GetRuleName gets the ruleName property value. The ruleName property
// returns a *string when successful
func (m *CustomWafRule) GetRuleName()(*string) {
    return m.ruleName
}
// GetShieldZoneId gets the shieldZoneId property value. The shieldZoneId property
// returns a *int64 when successful
func (m *CustomWafRule) GetShieldZoneId()(*int64) {
    return m.shieldZoneId
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *CustomWafRule) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *CustomWafRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ruleConfiguration", m.GetRuleConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleDescription", m.GetRuleDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleJson", m.GetRuleJson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("shieldZoneId", m.GetShieldZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *CustomWafRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *CustomWafRule) SetId(value *int64)() {
    m.id = value
}
// SetRuleConfiguration sets the ruleConfiguration property value. The ruleConfiguration property
func (m *CustomWafRule) SetRuleConfiguration(value CreateCustomWafRuleModelable)() {
    m.ruleConfiguration = value
}
// SetRuleDescription sets the ruleDescription property value. The ruleDescription property
func (m *CustomWafRule) SetRuleDescription(value *string)() {
    m.ruleDescription = value
}
// SetRuleJson sets the ruleJson property value. The ruleJson property
func (m *CustomWafRule) SetRuleJson(value *string)() {
    m.ruleJson = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *CustomWafRule) SetRuleName(value *string)() {
    m.ruleName = value
}
// SetShieldZoneId sets the shieldZoneId property value. The shieldZoneId property
func (m *CustomWafRule) SetShieldZoneId(value *int64)() {
    m.shieldZoneId = value
}
// SetUserId sets the userId property value. The userId property
func (m *CustomWafRule) SetUserId(value *string)() {
    m.userId = value
}
type CustomWafRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int64)
    GetRuleConfiguration()(CreateCustomWafRuleModelable)
    GetRuleDescription()(*string)
    GetRuleJson()(*string)
    GetRuleName()(*string)
    GetShieldZoneId()(*int64)
    GetUserId()(*string)
    SetId(value *int64)()
    SetRuleConfiguration(value CreateCustomWafRuleModelable)()
    SetRuleDescription(value *string)()
    SetRuleJson(value *string)()
    SetRuleName(value *string)()
    SetShieldZoneId(value *int64)()
    SetUserId(value *string)()
}
