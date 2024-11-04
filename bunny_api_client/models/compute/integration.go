package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Integration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DeployConfiguration property
    deployConfiguration Integration_DeployConfigurationable
    // The IntegrationId property
    integrationId *int64
    // The RepositorySettings property
    repositorySettings Integration_RepositorySettingsable
}
// NewIntegration instantiates a new Integration and sets the default values.
func NewIntegration()(*Integration) {
    m := &Integration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Integration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeployConfiguration gets the DeployConfiguration property value. The DeployConfiguration property
// returns a Integration_DeployConfigurationable when successful
func (m *Integration) GetDeployConfiguration()(Integration_DeployConfigurationable) {
    return m.deployConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Integration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["DeployConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegration_DeployConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployConfiguration(val.(Integration_DeployConfigurationable))
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
    res["RepositorySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegration_RepositorySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySettings(val.(Integration_RepositorySettingsable))
        }
        return nil
    }
    return res
}
// GetIntegrationId gets the IntegrationId property value. The IntegrationId property
// returns a *int64 when successful
func (m *Integration) GetIntegrationId()(*int64) {
    return m.integrationId
}
// GetRepositorySettings gets the RepositorySettings property value. The RepositorySettings property
// returns a Integration_RepositorySettingsable when successful
func (m *Integration) GetRepositorySettings()(Integration_RepositorySettingsable) {
    return m.repositorySettings
}
// Serialize serializes information the current object
func (m *Integration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("IntegrationId", m.GetIntegrationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("RepositorySettings", m.GetRepositorySettings())
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
func (m *Integration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeployConfiguration sets the DeployConfiguration property value. The DeployConfiguration property
func (m *Integration) SetDeployConfiguration(value Integration_DeployConfigurationable)() {
    m.deployConfiguration = value
}
// SetIntegrationId sets the IntegrationId property value. The IntegrationId property
func (m *Integration) SetIntegrationId(value *int64)() {
    m.integrationId = value
}
// SetRepositorySettings sets the RepositorySettings property value. The RepositorySettings property
func (m *Integration) SetRepositorySettings(value Integration_RepositorySettingsable)() {
    m.repositorySettings = value
}
type Integrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeployConfiguration()(Integration_DeployConfigurationable)
    GetIntegrationId()(*int64)
    GetRepositorySettings()(Integration_RepositorySettingsable)
    SetDeployConfiguration(value Integration_DeployConfigurationable)()
    SetIntegrationId(value *int64)()
    SetRepositorySettings(value Integration_RepositorySettingsable)()
}
