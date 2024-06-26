package getcountrylist

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Country struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The FlagUrl property
    flagUrl *string
    // The IsEU property
    isEU *bool
    // The IsoCode property
    isoCode *string
    // The Name property
    name *string
    // The PopList property
    popList []string
    // The TaxPrefix property
    taxPrefix *string
    // The TaxRate property
    taxRate *float64
}
// NewCountry instantiates a new Country and sets the default values.
func NewCountry()(*Country) {
    m := &Country{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCountryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCountryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCountry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Country) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Country) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["FlagUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlagUrl(val)
        }
        return nil
    }
    res["IsEU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEU(val)
        }
        return nil
    }
    res["IsoCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsoCode(val)
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
    res["PopList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPopList(res)
        }
        return nil
    }
    res["TaxPrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxPrefix(val)
        }
        return nil
    }
    res["TaxRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxRate(val)
        }
        return nil
    }
    return res
}
// GetFlagUrl gets the FlagUrl property value. The FlagUrl property
// returns a *string when successful
func (m *Country) GetFlagUrl()(*string) {
    return m.flagUrl
}
// GetIsEU gets the IsEU property value. The IsEU property
// returns a *bool when successful
func (m *Country) GetIsEU()(*bool) {
    return m.isEU
}
// GetIsoCode gets the IsoCode property value. The IsoCode property
// returns a *string when successful
func (m *Country) GetIsoCode()(*string) {
    return m.isoCode
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Country) GetName()(*string) {
    return m.name
}
// GetPopList gets the PopList property value. The PopList property
// returns a []string when successful
func (m *Country) GetPopList()([]string) {
    return m.popList
}
// GetTaxPrefix gets the TaxPrefix property value. The TaxPrefix property
// returns a *string when successful
func (m *Country) GetTaxPrefix()(*string) {
    return m.taxPrefix
}
// GetTaxRate gets the TaxRate property value. The TaxRate property
// returns a *float64 when successful
func (m *Country) GetTaxRate()(*float64) {
    return m.taxRate
}
// Serialize serializes information the current object
func (m *Country) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("FlagUrl", m.GetFlagUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("IsEU", m.GetIsEU())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("IsoCode", m.GetIsoCode())
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
    if m.GetPopList() != nil {
        err := writer.WriteCollectionOfStringValues("PopList", m.GetPopList())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("TaxPrefix", m.GetTaxPrefix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("TaxRate", m.GetTaxRate())
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
func (m *Country) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFlagUrl sets the FlagUrl property value. The FlagUrl property
func (m *Country) SetFlagUrl(value *string)() {
    m.flagUrl = value
}
// SetIsEU sets the IsEU property value. The IsEU property
func (m *Country) SetIsEU(value *bool)() {
    m.isEU = value
}
// SetIsoCode sets the IsoCode property value. The IsoCode property
func (m *Country) SetIsoCode(value *string)() {
    m.isoCode = value
}
// SetName sets the Name property value. The Name property
func (m *Country) SetName(value *string)() {
    m.name = value
}
// SetPopList sets the PopList property value. The PopList property
func (m *Country) SetPopList(value []string)() {
    m.popList = value
}
// SetTaxPrefix sets the TaxPrefix property value. The TaxPrefix property
func (m *Country) SetTaxPrefix(value *string)() {
    m.taxPrefix = value
}
// SetTaxRate sets the TaxRate property value. The TaxRate property
func (m *Country) SetTaxRate(value *float64)() {
    m.taxRate = value
}
type Countryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFlagUrl()(*string)
    GetIsEU()(*bool)
    GetIsoCode()(*string)
    GetName()(*string)
    GetPopList()([]string)
    GetTaxPrefix()(*string)
    GetTaxRate()(*float64)
    SetFlagUrl(value *string)()
    SetIsEU(value *bool)()
    SetIsoCode(value *string)()
    SetName(value *string)()
    SetPopList(value []string)()
    SetTaxPrefix(value *string)()
    SetTaxRate(value *float64)()
}
