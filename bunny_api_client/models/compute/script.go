package compute

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Script struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CurrentReleaseId property
    currentReleaseId *int64
    // The DefaultHostname property
    defaultHostname *string
    // The Deleted property
    deleted *bool
    // The DeploymentKey property
    deploymentKey *string
    // The EdgeScriptVariables property
    edgeScriptVariables []EdgeScriptVariableable
    // The Id property
    id *int64
    // The Integration property
    integration Integrationable
    // The IntegrationId property
    integrationId *int64
    // The LastModified property
    lastModified *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The LinkedPullZones property
    linkedPullZones []LinkedPullZoneable
    // The MonthlyCost property
    monthlyCost *float64
    // The MonthlyCpuTime property
    monthlyCpuTime *int32
    // The MonthlyRequestCount property
    monthlyRequestCount *int32
    // The Name property
    name *string
    // The RepositoryId property
    repositoryId *int64
    // The ScriptType property
    scriptType *float64
    // The SystemHostname property
    systemHostname *string
}
// NewScript instantiates a new Script and sets the default values.
func NewScript()(*Script) {
    m := &Script{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScript(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Script) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentReleaseId gets the CurrentReleaseId property value. The CurrentReleaseId property
// returns a *int64 when successful
func (m *Script) GetCurrentReleaseId()(*int64) {
    return m.currentReleaseId
}
// GetDefaultHostname gets the DefaultHostname property value. The DefaultHostname property
// returns a *string when successful
func (m *Script) GetDefaultHostname()(*string) {
    return m.defaultHostname
}
// GetDeleted gets the Deleted property value. The Deleted property
// returns a *bool when successful
func (m *Script) GetDeleted()(*bool) {
    return m.deleted
}
// GetDeploymentKey gets the DeploymentKey property value. The DeploymentKey property
// returns a *string when successful
func (m *Script) GetDeploymentKey()(*string) {
    return m.deploymentKey
}
// GetEdgeScriptVariables gets the EdgeScriptVariables property value. The EdgeScriptVariables property
// returns a []EdgeScriptVariableable when successful
func (m *Script) GetEdgeScriptVariables()([]EdgeScriptVariableable) {
    return m.edgeScriptVariables
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Script) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CurrentReleaseId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentReleaseId(val)
        }
        return nil
    }
    res["DefaultHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultHostname(val)
        }
        return nil
    }
    res["Deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
        }
        return nil
    }
    res["DeploymentKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentKey(val)
        }
        return nil
    }
    res["EdgeScriptVariables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdgeScriptVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdgeScriptVariableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdgeScriptVariableable)
                }
            }
            m.SetEdgeScriptVariables(res)
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
    res["IntegrationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationId(val)
        }
        return nil
    }
    res["LastModified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModified(val)
        }
        return nil
    }
    res["LinkedPullZones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkedPullZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkedPullZoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LinkedPullZoneable)
                }
            }
            m.SetLinkedPullZones(res)
        }
        return nil
    }
    res["MonthlyCost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyCost(val)
        }
        return nil
    }
    res["MonthlyCpuTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyCpuTime(val)
        }
        return nil
    }
    res["MonthlyRequestCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyRequestCount(val)
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
    res["RepositoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
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
    res["SystemHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemHostname(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *Script) GetId()(*int64) {
    return m.id
}
// GetIntegration gets the Integration property value. The Integration property
// returns a Integrationable when successful
func (m *Script) GetIntegration()(Integrationable) {
    return m.integration
}
// GetIntegrationId gets the IntegrationId property value. The IntegrationId property
// returns a *int64 when successful
func (m *Script) GetIntegrationId()(*int64) {
    return m.integrationId
}
// GetLastModified gets the LastModified property value. The LastModified property
// returns a *Time when successful
func (m *Script) GetLastModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModified
}
// GetLinkedPullZones gets the LinkedPullZones property value. The LinkedPullZones property
// returns a []LinkedPullZoneable when successful
func (m *Script) GetLinkedPullZones()([]LinkedPullZoneable) {
    return m.linkedPullZones
}
// GetMonthlyCost gets the MonthlyCost property value. The MonthlyCost property
// returns a *float64 when successful
func (m *Script) GetMonthlyCost()(*float64) {
    return m.monthlyCost
}
// GetMonthlyCpuTime gets the MonthlyCpuTime property value. The MonthlyCpuTime property
// returns a *int32 when successful
func (m *Script) GetMonthlyCpuTime()(*int32) {
    return m.monthlyCpuTime
}
// GetMonthlyRequestCount gets the MonthlyRequestCount property value. The MonthlyRequestCount property
// returns a *int32 when successful
func (m *Script) GetMonthlyRequestCount()(*int32) {
    return m.monthlyRequestCount
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Script) GetName()(*string) {
    return m.name
}
// GetRepositoryId gets the RepositoryId property value. The RepositoryId property
// returns a *int64 when successful
func (m *Script) GetRepositoryId()(*int64) {
    return m.repositoryId
}
// GetScriptType gets the ScriptType property value. The ScriptType property
// returns a *float64 when successful
func (m *Script) GetScriptType()(*float64) {
    return m.scriptType
}
// GetSystemHostname gets the SystemHostname property value. The SystemHostname property
// returns a *string when successful
func (m *Script) GetSystemHostname()(*string) {
    return m.systemHostname
}
// Serialize serializes information the current object
func (m *Script) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Script) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentReleaseId sets the CurrentReleaseId property value. The CurrentReleaseId property
func (m *Script) SetCurrentReleaseId(value *int64)() {
    m.currentReleaseId = value
}
// SetDefaultHostname sets the DefaultHostname property value. The DefaultHostname property
func (m *Script) SetDefaultHostname(value *string)() {
    m.defaultHostname = value
}
// SetDeleted sets the Deleted property value. The Deleted property
func (m *Script) SetDeleted(value *bool)() {
    m.deleted = value
}
// SetDeploymentKey sets the DeploymentKey property value. The DeploymentKey property
func (m *Script) SetDeploymentKey(value *string)() {
    m.deploymentKey = value
}
// SetEdgeScriptVariables sets the EdgeScriptVariables property value. The EdgeScriptVariables property
func (m *Script) SetEdgeScriptVariables(value []EdgeScriptVariableable)() {
    m.edgeScriptVariables = value
}
// SetId sets the Id property value. The Id property
func (m *Script) SetId(value *int64)() {
    m.id = value
}
// SetIntegration sets the Integration property value. The Integration property
func (m *Script) SetIntegration(value Integrationable)() {
    m.integration = value
}
// SetIntegrationId sets the IntegrationId property value. The IntegrationId property
func (m *Script) SetIntegrationId(value *int64)() {
    m.integrationId = value
}
// SetLastModified sets the LastModified property value. The LastModified property
func (m *Script) SetLastModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModified = value
}
// SetLinkedPullZones sets the LinkedPullZones property value. The LinkedPullZones property
func (m *Script) SetLinkedPullZones(value []LinkedPullZoneable)() {
    m.linkedPullZones = value
}
// SetMonthlyCost sets the MonthlyCost property value. The MonthlyCost property
func (m *Script) SetMonthlyCost(value *float64)() {
    m.monthlyCost = value
}
// SetMonthlyCpuTime sets the MonthlyCpuTime property value. The MonthlyCpuTime property
func (m *Script) SetMonthlyCpuTime(value *int32)() {
    m.monthlyCpuTime = value
}
// SetMonthlyRequestCount sets the MonthlyRequestCount property value. The MonthlyRequestCount property
func (m *Script) SetMonthlyRequestCount(value *int32)() {
    m.monthlyRequestCount = value
}
// SetName sets the Name property value. The Name property
func (m *Script) SetName(value *string)() {
    m.name = value
}
// SetRepositoryId sets the RepositoryId property value. The RepositoryId property
func (m *Script) SetRepositoryId(value *int64)() {
    m.repositoryId = value
}
// SetScriptType sets the ScriptType property value. The ScriptType property
func (m *Script) SetScriptType(value *float64)() {
    m.scriptType = value
}
// SetSystemHostname sets the SystemHostname property value. The SystemHostname property
func (m *Script) SetSystemHostname(value *string)() {
    m.systemHostname = value
}
type Scriptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentReleaseId()(*int64)
    GetDefaultHostname()(*string)
    GetDeleted()(*bool)
    GetDeploymentKey()(*string)
    GetEdgeScriptVariables()([]EdgeScriptVariableable)
    GetId()(*int64)
    GetIntegration()(Integrationable)
    GetIntegrationId()(*int64)
    GetLastModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinkedPullZones()([]LinkedPullZoneable)
    GetMonthlyCost()(*float64)
    GetMonthlyCpuTime()(*int32)
    GetMonthlyRequestCount()(*int32)
    GetName()(*string)
    GetRepositoryId()(*int64)
    GetScriptType()(*float64)
    GetSystemHostname()(*string)
    SetCurrentReleaseId(value *int64)()
    SetDefaultHostname(value *string)()
    SetDeleted(value *bool)()
    SetDeploymentKey(value *string)()
    SetEdgeScriptVariables(value []EdgeScriptVariableable)()
    SetId(value *int64)()
    SetIntegration(value Integrationable)()
    SetIntegrationId(value *int64)()
    SetLastModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinkedPullZones(value []LinkedPullZoneable)()
    SetMonthlyCost(value *float64)()
    SetMonthlyCpuTime(value *int32)()
    SetMonthlyRequestCount(value *int32)()
    SetName(value *string)()
    SetRepositoryId(value *int64)()
    SetScriptType(value *float64)()
    SetSystemHostname(value *string)()
}
