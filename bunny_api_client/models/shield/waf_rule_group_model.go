package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WafRuleGroupModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The description property
    description *string
    // The fileName property
    fileName *string
    // The id property
    id *int64
    // The mainGroup property
    mainGroup *string
    // The name property
    name *string
    // The rules property
    rules []WafRuleModelable
    // The ruleset property
    ruleset *string
}
// NewWafRuleGroupModel instantiates a new WafRuleGroupModel and sets the default values.
func NewWafRuleGroupModel()(*WafRuleGroupModel) {
    m := &WafRuleGroupModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWafRuleGroupModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWafRuleGroupModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWafRuleGroupModel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WafRuleGroupModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *WafRuleGroupModel) GetCode()(*string) {
    return m.code
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *WafRuleGroupModel) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WafRuleGroupModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
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
    res["mainGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMainGroup(val)
        }
        return nil
    }
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
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWafRuleModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WafRuleModelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WafRuleModelable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    res["ruleset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleset(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
// returns a *string when successful
func (m *WafRuleGroupModel) GetFileName()(*string) {
    return m.fileName
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *WafRuleGroupModel) GetId()(*int64) {
    return m.id
}
// GetMainGroup gets the mainGroup property value. The mainGroup property
// returns a *string when successful
func (m *WafRuleGroupModel) GetMainGroup()(*string) {
    return m.mainGroup
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *WafRuleGroupModel) GetName()(*string) {
    return m.name
}
// GetRules gets the rules property value. The rules property
// returns a []WafRuleModelable when successful
func (m *WafRuleGroupModel) GetRules()([]WafRuleModelable) {
    return m.rules
}
// GetRuleset gets the ruleset property value. The ruleset property
// returns a *string when successful
func (m *WafRuleGroupModel) GetRuleset()(*string) {
    return m.ruleset
}
// Serialize serializes information the current object
func (m *WafRuleGroupModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mainGroup", m.GetMainGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleset", m.GetRuleset())
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
func (m *WafRuleGroupModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *WafRuleGroupModel) SetCode(value *string)() {
    m.code = value
}
// SetDescription sets the description property value. The description property
func (m *WafRuleGroupModel) SetDescription(value *string)() {
    m.description = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *WafRuleGroupModel) SetFileName(value *string)() {
    m.fileName = value
}
// SetId sets the id property value. The id property
func (m *WafRuleGroupModel) SetId(value *int64)() {
    m.id = value
}
// SetMainGroup sets the mainGroup property value. The mainGroup property
func (m *WafRuleGroupModel) SetMainGroup(value *string)() {
    m.mainGroup = value
}
// SetName sets the name property value. The name property
func (m *WafRuleGroupModel) SetName(value *string)() {
    m.name = value
}
// SetRules sets the rules property value. The rules property
func (m *WafRuleGroupModel) SetRules(value []WafRuleModelable)() {
    m.rules = value
}
// SetRuleset sets the ruleset property value. The ruleset property
func (m *WafRuleGroupModel) SetRuleset(value *string)() {
    m.ruleset = value
}
type WafRuleGroupModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDescription()(*string)
    GetFileName()(*string)
    GetId()(*int64)
    GetMainGroup()(*string)
    GetName()(*string)
    GetRules()([]WafRuleModelable)
    GetRuleset()(*string)
    SetCode(value *string)()
    SetDescription(value *string)()
    SetFileName(value *string)()
    SetId(value *int64)()
    SetMainGroup(value *string)()
    SetName(value *string)()
    SetRules(value []WafRuleModelable)()
    SetRuleset(value *string)()
}
