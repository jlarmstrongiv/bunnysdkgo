package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ShieldZoneRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dDoSChallengeWindow property
    dDoSChallengeWindow *int32
    // The dDoSShieldSensitivity property
    dDoSShieldSensitivity *float64
    // The learningMode property
    learningMode *bool
    // The premiumPlan property
    premiumPlan *bool
    // The shieldZoneId property
    shieldZoneId *int64
    // The wafDisabledRuleGroups property
    wafDisabledRuleGroups []string
    // The wafDisabledRules property
    wafDisabledRules []string
    // The wafEnabled property
    wafEnabled *bool
    // The wafEngineConfig property
    wafEngineConfig []PullZoneWafConfigVariableModelable
    // The wafExecutionMode property
    wafExecutionMode *float64
    // The wafLogOnlyRules property
    wafLogOnlyRules []string
    // The wafProfileId property
    wafProfileId *int32
    // The wafRequestHeaderLoggingEnabled property
    wafRequestHeaderLoggingEnabled *bool
    // The wafRequestIgnoredHeaders property
    wafRequestIgnoredHeaders []string
}
// NewShieldZoneRequest instantiates a new ShieldZoneRequest and sets the default values.
func NewShieldZoneRequest()(*ShieldZoneRequest) {
    m := &ShieldZoneRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateShieldZoneRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateShieldZoneRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShieldZoneRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ShieldZoneRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDDoSChallengeWindow gets the dDoSChallengeWindow property value. The dDoSChallengeWindow property
// returns a *int32 when successful
func (m *ShieldZoneRequest) GetDDoSChallengeWindow()(*int32) {
    return m.dDoSChallengeWindow
}
// GetDDoSShieldSensitivity gets the dDoSShieldSensitivity property value. The dDoSShieldSensitivity property
// returns a *float64 when successful
func (m *ShieldZoneRequest) GetDDoSShieldSensitivity()(*float64) {
    return m.dDoSShieldSensitivity
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ShieldZoneRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dDoSChallengeWindow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDDoSChallengeWindow(val)
        }
        return nil
    }
    res["dDoSShieldSensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDDoSShieldSensitivity(val)
        }
        return nil
    }
    res["learningMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLearningMode(val)
        }
        return nil
    }
    res["premiumPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPremiumPlan(val)
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
    res["wafDisabledRuleGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafDisabledRuleGroups(res)
        }
        return nil
    }
    res["wafDisabledRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafDisabledRules(res)
        }
        return nil
    }
    res["wafEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafEnabled(val)
        }
        return nil
    }
    res["wafEngineConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePullZoneWafConfigVariableModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullZoneWafConfigVariableModelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullZoneWafConfigVariableModelable)
                }
            }
            m.SetWafEngineConfig(res)
        }
        return nil
    }
    res["wafExecutionMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafExecutionMode(val)
        }
        return nil
    }
    res["wafLogOnlyRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafLogOnlyRules(res)
        }
        return nil
    }
    res["wafProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafProfileId(val)
        }
        return nil
    }
    res["wafRequestHeaderLoggingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafRequestHeaderLoggingEnabled(val)
        }
        return nil
    }
    res["wafRequestIgnoredHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafRequestIgnoredHeaders(res)
        }
        return nil
    }
    return res
}
// GetLearningMode gets the learningMode property value. The learningMode property
// returns a *bool when successful
func (m *ShieldZoneRequest) GetLearningMode()(*bool) {
    return m.learningMode
}
// GetPremiumPlan gets the premiumPlan property value. The premiumPlan property
// returns a *bool when successful
func (m *ShieldZoneRequest) GetPremiumPlan()(*bool) {
    return m.premiumPlan
}
// GetShieldZoneId gets the shieldZoneId property value. The shieldZoneId property
// returns a *int64 when successful
func (m *ShieldZoneRequest) GetShieldZoneId()(*int64) {
    return m.shieldZoneId
}
// GetWafDisabledRuleGroups gets the wafDisabledRuleGroups property value. The wafDisabledRuleGroups property
// returns a []string when successful
func (m *ShieldZoneRequest) GetWafDisabledRuleGroups()([]string) {
    return m.wafDisabledRuleGroups
}
// GetWafDisabledRules gets the wafDisabledRules property value. The wafDisabledRules property
// returns a []string when successful
func (m *ShieldZoneRequest) GetWafDisabledRules()([]string) {
    return m.wafDisabledRules
}
// GetWafEnabled gets the wafEnabled property value. The wafEnabled property
// returns a *bool when successful
func (m *ShieldZoneRequest) GetWafEnabled()(*bool) {
    return m.wafEnabled
}
// GetWafEngineConfig gets the wafEngineConfig property value. The wafEngineConfig property
// returns a []PullZoneWafConfigVariableModelable when successful
func (m *ShieldZoneRequest) GetWafEngineConfig()([]PullZoneWafConfigVariableModelable) {
    return m.wafEngineConfig
}
// GetWafExecutionMode gets the wafExecutionMode property value. The wafExecutionMode property
// returns a *float64 when successful
func (m *ShieldZoneRequest) GetWafExecutionMode()(*float64) {
    return m.wafExecutionMode
}
// GetWafLogOnlyRules gets the wafLogOnlyRules property value. The wafLogOnlyRules property
// returns a []string when successful
func (m *ShieldZoneRequest) GetWafLogOnlyRules()([]string) {
    return m.wafLogOnlyRules
}
// GetWafProfileId gets the wafProfileId property value. The wafProfileId property
// returns a *int32 when successful
func (m *ShieldZoneRequest) GetWafProfileId()(*int32) {
    return m.wafProfileId
}
// GetWafRequestHeaderLoggingEnabled gets the wafRequestHeaderLoggingEnabled property value. The wafRequestHeaderLoggingEnabled property
// returns a *bool when successful
func (m *ShieldZoneRequest) GetWafRequestHeaderLoggingEnabled()(*bool) {
    return m.wafRequestHeaderLoggingEnabled
}
// GetWafRequestIgnoredHeaders gets the wafRequestIgnoredHeaders property value. The wafRequestIgnoredHeaders property
// returns a []string when successful
func (m *ShieldZoneRequest) GetWafRequestIgnoredHeaders()([]string) {
    return m.wafRequestIgnoredHeaders
}
// Serialize serializes information the current object
func (m *ShieldZoneRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dDoSChallengeWindow", m.GetDDoSChallengeWindow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("dDoSShieldSensitivity", m.GetDDoSShieldSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("learningMode", m.GetLearningMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("premiumPlan", m.GetPremiumPlan())
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
    if m.GetWafDisabledRuleGroups() != nil {
        err := writer.WriteCollectionOfStringValues("wafDisabledRuleGroups", m.GetWafDisabledRuleGroups())
        if err != nil {
            return err
        }
    }
    if m.GetWafDisabledRules() != nil {
        err := writer.WriteCollectionOfStringValues("wafDisabledRules", m.GetWafDisabledRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wafEnabled", m.GetWafEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetWafEngineConfig() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWafEngineConfig()))
        for i, v := range m.GetWafEngineConfig() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("wafEngineConfig", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("wafExecutionMode", m.GetWafExecutionMode())
        if err != nil {
            return err
        }
    }
    if m.GetWafLogOnlyRules() != nil {
        err := writer.WriteCollectionOfStringValues("wafLogOnlyRules", m.GetWafLogOnlyRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wafProfileId", m.GetWafProfileId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wafRequestHeaderLoggingEnabled", m.GetWafRequestHeaderLoggingEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetWafRequestIgnoredHeaders() != nil {
        err := writer.WriteCollectionOfStringValues("wafRequestIgnoredHeaders", m.GetWafRequestIgnoredHeaders())
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
func (m *ShieldZoneRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDDoSChallengeWindow sets the dDoSChallengeWindow property value. The dDoSChallengeWindow property
func (m *ShieldZoneRequest) SetDDoSChallengeWindow(value *int32)() {
    m.dDoSChallengeWindow = value
}
// SetDDoSShieldSensitivity sets the dDoSShieldSensitivity property value. The dDoSShieldSensitivity property
func (m *ShieldZoneRequest) SetDDoSShieldSensitivity(value *float64)() {
    m.dDoSShieldSensitivity = value
}
// SetLearningMode sets the learningMode property value. The learningMode property
func (m *ShieldZoneRequest) SetLearningMode(value *bool)() {
    m.learningMode = value
}
// SetPremiumPlan sets the premiumPlan property value. The premiumPlan property
func (m *ShieldZoneRequest) SetPremiumPlan(value *bool)() {
    m.premiumPlan = value
}
// SetShieldZoneId sets the shieldZoneId property value. The shieldZoneId property
func (m *ShieldZoneRequest) SetShieldZoneId(value *int64)() {
    m.shieldZoneId = value
}
// SetWafDisabledRuleGroups sets the wafDisabledRuleGroups property value. The wafDisabledRuleGroups property
func (m *ShieldZoneRequest) SetWafDisabledRuleGroups(value []string)() {
    m.wafDisabledRuleGroups = value
}
// SetWafDisabledRules sets the wafDisabledRules property value. The wafDisabledRules property
func (m *ShieldZoneRequest) SetWafDisabledRules(value []string)() {
    m.wafDisabledRules = value
}
// SetWafEnabled sets the wafEnabled property value. The wafEnabled property
func (m *ShieldZoneRequest) SetWafEnabled(value *bool)() {
    m.wafEnabled = value
}
// SetWafEngineConfig sets the wafEngineConfig property value. The wafEngineConfig property
func (m *ShieldZoneRequest) SetWafEngineConfig(value []PullZoneWafConfigVariableModelable)() {
    m.wafEngineConfig = value
}
// SetWafExecutionMode sets the wafExecutionMode property value. The wafExecutionMode property
func (m *ShieldZoneRequest) SetWafExecutionMode(value *float64)() {
    m.wafExecutionMode = value
}
// SetWafLogOnlyRules sets the wafLogOnlyRules property value. The wafLogOnlyRules property
func (m *ShieldZoneRequest) SetWafLogOnlyRules(value []string)() {
    m.wafLogOnlyRules = value
}
// SetWafProfileId sets the wafProfileId property value. The wafProfileId property
func (m *ShieldZoneRequest) SetWafProfileId(value *int32)() {
    m.wafProfileId = value
}
// SetWafRequestHeaderLoggingEnabled sets the wafRequestHeaderLoggingEnabled property value. The wafRequestHeaderLoggingEnabled property
func (m *ShieldZoneRequest) SetWafRequestHeaderLoggingEnabled(value *bool)() {
    m.wafRequestHeaderLoggingEnabled = value
}
// SetWafRequestIgnoredHeaders sets the wafRequestIgnoredHeaders property value. The wafRequestIgnoredHeaders property
func (m *ShieldZoneRequest) SetWafRequestIgnoredHeaders(value []string)() {
    m.wafRequestIgnoredHeaders = value
}
type ShieldZoneRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDDoSChallengeWindow()(*int32)
    GetDDoSShieldSensitivity()(*float64)
    GetLearningMode()(*bool)
    GetPremiumPlan()(*bool)
    GetShieldZoneId()(*int64)
    GetWafDisabledRuleGroups()([]string)
    GetWafDisabledRules()([]string)
    GetWafEnabled()(*bool)
    GetWafEngineConfig()([]PullZoneWafConfigVariableModelable)
    GetWafExecutionMode()(*float64)
    GetWafLogOnlyRules()([]string)
    GetWafProfileId()(*int32)
    GetWafRequestHeaderLoggingEnabled()(*bool)
    GetWafRequestIgnoredHeaders()([]string)
    SetDDoSChallengeWindow(value *int32)()
    SetDDoSShieldSensitivity(value *float64)()
    SetLearningMode(value *bool)()
    SetPremiumPlan(value *bool)()
    SetShieldZoneId(value *int64)()
    SetWafDisabledRuleGroups(value []string)()
    SetWafDisabledRules(value []string)()
    SetWafEnabled(value *bool)()
    SetWafEngineConfig(value []PullZoneWafConfigVariableModelable)()
    SetWafExecutionMode(value *float64)()
    SetWafLogOnlyRules(value []string)()
    SetWafProfileId(value *int32)()
    SetWafRequestHeaderLoggingEnabled(value *bool)()
    SetWafRequestIgnoredHeaders(value []string)()
}
