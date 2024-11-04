package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetCustomWafRulesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data []CustomWafRuleable
    // The page property
    page PaginationResponseable
}
// NewGetCustomWafRulesResponse instantiates a new GetCustomWafRulesResponse and sets the default values.
func NewGetCustomWafRulesResponse()(*GetCustomWafRulesResponse) {
    m := &GetCustomWafRulesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetCustomWafRulesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetCustomWafRulesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetCustomWafRulesResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetCustomWafRulesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a []CustomWafRuleable when successful
func (m *GetCustomWafRulesResponse) GetData()([]CustomWafRuleable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetCustomWafRulesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomWafRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomWafRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomWafRuleable)
                }
            }
            m.SetData(res)
        }
        return nil
    }
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePaginationResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val.(PaginationResponseable))
        }
        return nil
    }
    return res
}
// GetPage gets the page property value. The page property
// returns a PaginationResponseable when successful
func (m *GetCustomWafRulesResponse) GetPage()(PaginationResponseable) {
    return m.page
}
// Serialize serializes information the current object
func (m *GetCustomWafRulesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetData()))
        for i, v := range m.GetData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("data", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("page", m.GetPage())
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
func (m *GetCustomWafRulesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *GetCustomWafRulesResponse) SetData(value []CustomWafRuleable)() {
    m.data = value
}
// SetPage sets the page property value. The page property
func (m *GetCustomWafRulesResponse) SetPage(value PaginationResponseable)() {
    m.page = value
}
type GetCustomWafRulesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()([]CustomWafRuleable)
    GetPage()(PaginationResponseable)
    SetData(value []CustomWafRuleable)()
    SetPage(value PaginationResponseable)()
}
