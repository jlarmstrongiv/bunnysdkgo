package regionlist

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Region struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The AllowLatencyRouting property
    allowLatencyRouting *bool
    // The ContinentCode property
    continentCode *string
    // The CountryCode property
    countryCode *string
    // The Id property
    id *int64
    // The Latitude property
    latitude *float64
    // The Longitude property
    longitude *float64
    // The Name property
    name *string
    // The PricePerGigabyte property
    pricePerGigabyte *float64
    // The RegionCode property
    regionCode *string
}
// NewRegion instantiates a new Region and sets the default values.
func NewRegion()(*Region) {
    m := &Region{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRegionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRegionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Region) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowLatencyRouting gets the AllowLatencyRouting property value. The AllowLatencyRouting property
// returns a *bool when successful
func (m *Region) GetAllowLatencyRouting()(*bool) {
    return m.allowLatencyRouting
}
// GetContinentCode gets the ContinentCode property value. The ContinentCode property
// returns a *string when successful
func (m *Region) GetContinentCode()(*string) {
    return m.continentCode
}
// GetCountryCode gets the CountryCode property value. The CountryCode property
// returns a *string when successful
func (m *Region) GetCountryCode()(*string) {
    return m.countryCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Region) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AllowLatencyRouting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLatencyRouting(val)
        }
        return nil
    }
    res["ContinentCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinentCode(val)
        }
        return nil
    }
    res["CountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
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
    res["Latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["Longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
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
    res["PricePerGigabyte"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricePerGigabyte(val)
        }
        return nil
    }
    res["RegionCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionCode(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *Region) GetId()(*int64) {
    return m.id
}
// GetLatitude gets the Latitude property value. The Latitude property
// returns a *float64 when successful
func (m *Region) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the Longitude property value. The Longitude property
// returns a *float64 when successful
func (m *Region) GetLongitude()(*float64) {
    return m.longitude
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Region) GetName()(*string) {
    return m.name
}
// GetPricePerGigabyte gets the PricePerGigabyte property value. The PricePerGigabyte property
// returns a *float64 when successful
func (m *Region) GetPricePerGigabyte()(*float64) {
    return m.pricePerGigabyte
}
// GetRegionCode gets the RegionCode property value. The RegionCode property
// returns a *string when successful
func (m *Region) GetRegionCode()(*string) {
    return m.regionCode
}
// Serialize serializes information the current object
func (m *Region) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("AllowLatencyRouting", m.GetAllowLatencyRouting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ContinentCode", m.GetContinentCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CountryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Longitude", m.GetLongitude())
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
        err := writer.WriteFloat64Value("PricePerGigabyte", m.GetPricePerGigabyte())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("RegionCode", m.GetRegionCode())
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
func (m *Region) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowLatencyRouting sets the AllowLatencyRouting property value. The AllowLatencyRouting property
func (m *Region) SetAllowLatencyRouting(value *bool)() {
    m.allowLatencyRouting = value
}
// SetContinentCode sets the ContinentCode property value. The ContinentCode property
func (m *Region) SetContinentCode(value *string)() {
    m.continentCode = value
}
// SetCountryCode sets the CountryCode property value. The CountryCode property
func (m *Region) SetCountryCode(value *string)() {
    m.countryCode = value
}
// SetId sets the Id property value. The Id property
func (m *Region) SetId(value *int64)() {
    m.id = value
}
// SetLatitude sets the Latitude property value. The Latitude property
func (m *Region) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the Longitude property value. The Longitude property
func (m *Region) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetName sets the Name property value. The Name property
func (m *Region) SetName(value *string)() {
    m.name = value
}
// SetPricePerGigabyte sets the PricePerGigabyte property value. The PricePerGigabyte property
func (m *Region) SetPricePerGigabyte(value *float64)() {
    m.pricePerGigabyte = value
}
// SetRegionCode sets the RegionCode property value. The RegionCode property
func (m *Region) SetRegionCode(value *string)() {
    m.regionCode = value
}
type Regionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowLatencyRouting()(*bool)
    GetContinentCode()(*string)
    GetCountryCode()(*string)
    GetId()(*int64)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetName()(*string)
    GetPricePerGigabyte()(*float64)
    GetRegionCode()(*string)
    SetAllowLatencyRouting(value *bool)()
    SetContinentCode(value *string)()
    SetCountryCode(value *string)()
    SetId(value *int64)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetName(value *string)()
    SetPricePerGigabyte(value *float64)()
    SetRegionCode(value *string)()
}
