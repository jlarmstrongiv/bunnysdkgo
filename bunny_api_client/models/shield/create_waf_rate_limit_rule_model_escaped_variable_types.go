package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateWafRateLimitRuleModel_variableTypes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ARGS property
    aRGS *string
    // The ARGS_COMBINED_SIZE property
    aRGS_COMBINED_SIZE *string
    // The ARGS_GET property
    aRGS_GET *string
    // The ARGS_GET_NAMES property
    aRGS_GET_NAMES *string
    // The ARGS_POST property
    aRGS_POST *string
    // The ARGS_POST_NAMES property
    aRGS_POST_NAMES *string
    // The FILES_NAMES property
    fILES_NAMES *string
    // The GEO property
    gEO *string
    // The QUERY_STRING property
    qUERY_STRING *string
    // The REMOTE_ADDR property
    rEMOTE_ADDR *string
    // The REQUEST_BASENAME property
    rEQUEST_BASENAME *string
    // The REQUEST_BODY property
    rEQUEST_BODY *string
    // The REQUEST_COOKIES property
    rEQUEST_COOKIES *string
    // The REQUEST_COOKIES_NAMES property
    rEQUEST_COOKIES_NAMES *string
    // The REQUEST_FILENAME property
    rEQUEST_FILENAME *string
    // The REQUEST_HEADERS property
    rEQUEST_HEADERS *string
    // The REQUEST_HEADERS_NAMES property
    rEQUEST_HEADERS_NAMES *string
    // The REQUEST_LINE property
    rEQUEST_LINE *string
    // The REQUEST_METHOD property
    rEQUEST_METHOD *string
    // The REQUEST_PROTOCOL property
    rEQUEST_PROTOCOL *string
    // The REQUEST_URI property
    rEQUEST_URI *string
    // The REQUEST_URI_RAW property
    rEQUEST_URI_RAW *string
    // The RESPONSE_BODY property
    rESPONSE_BODY *string
    // The RESPONSE_HEADERS property
    rESPONSE_HEADERS *string
    // The RESPONSE_STATUS property
    rESPONSE_STATUS *string
}
// NewCreateWafRateLimitRuleModel_variableTypes instantiates a new CreateWafRateLimitRuleModel_variableTypes and sets the default values.
func NewCreateWafRateLimitRuleModel_variableTypes()(*CreateWafRateLimitRuleModel_variableTypes) {
    m := &CreateWafRateLimitRuleModel_variableTypes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateWafRateLimitRuleModel_variableTypesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateWafRateLimitRuleModel_variableTypesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateWafRateLimitRuleModel_variableTypes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetARGS gets the ARGS property value. The ARGS property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGS()(*string) {
    return m.aRGS
}
// GetARGSCOMBINEDSIZE gets the ARGS_COMBINED_SIZE property value. The ARGS_COMBINED_SIZE property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGSCOMBINEDSIZE()(*string) {
    return m.aRGS_COMBINED_SIZE
}
// GetARGSGET gets the ARGS_GET property value. The ARGS_GET property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGSGET()(*string) {
    return m.aRGS_GET
}
// GetARGSGETNAMES gets the ARGS_GET_NAMES property value. The ARGS_GET_NAMES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGSGETNAMES()(*string) {
    return m.aRGS_GET_NAMES
}
// GetARGSPOST gets the ARGS_POST property value. The ARGS_POST property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGSPOST()(*string) {
    return m.aRGS_POST
}
// GetARGSPOSTNAMES gets the ARGS_POST_NAMES property value. The ARGS_POST_NAMES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetARGSPOSTNAMES()(*string) {
    return m.aRGS_POST_NAMES
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ARGS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGS(val)
        }
        return nil
    }
    res["ARGS_COMBINED_SIZE"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGSCOMBINEDSIZE(val)
        }
        return nil
    }
    res["ARGS_GET"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGSGET(val)
        }
        return nil
    }
    res["ARGS_GET_NAMES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGSGETNAMES(val)
        }
        return nil
    }
    res["ARGS_POST"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGSPOST(val)
        }
        return nil
    }
    res["ARGS_POST_NAMES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetARGSPOSTNAMES(val)
        }
        return nil
    }
    res["FILES_NAMES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFILESNAMES(val)
        }
        return nil
    }
    res["GEO"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGEO(val)
        }
        return nil
    }
    res["QUERY_STRING"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQUERYSTRING(val)
        }
        return nil
    }
    res["REMOTE_ADDR"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREMOTEADDR(val)
        }
        return nil
    }
    res["REQUEST_BASENAME"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTBASENAME(val)
        }
        return nil
    }
    res["REQUEST_BODY"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTBODY(val)
        }
        return nil
    }
    res["REQUEST_COOKIES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTCOOKIES(val)
        }
        return nil
    }
    res["REQUEST_COOKIES_NAMES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTCOOKIESNAMES(val)
        }
        return nil
    }
    res["REQUEST_FILENAME"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTFILENAME(val)
        }
        return nil
    }
    res["REQUEST_HEADERS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTHEADERS(val)
        }
        return nil
    }
    res["REQUEST_HEADERS_NAMES"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTHEADERSNAMES(val)
        }
        return nil
    }
    res["REQUEST_LINE"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTLINE(val)
        }
        return nil
    }
    res["REQUEST_METHOD"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTMETHOD(val)
        }
        return nil
    }
    res["REQUEST_PROTOCOL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTPROTOCOL(val)
        }
        return nil
    }
    res["REQUEST_URI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTURI(val)
        }
        return nil
    }
    res["REQUEST_URI_RAW"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetREQUESTURIRAW(val)
        }
        return nil
    }
    res["RESPONSE_BODY"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRESPONSEBODY(val)
        }
        return nil
    }
    res["RESPONSE_HEADERS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRESPONSEHEADERS(val)
        }
        return nil
    }
    res["RESPONSE_STATUS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRESPONSESTATUS(val)
        }
        return nil
    }
    return res
}
// GetFILESNAMES gets the FILES_NAMES property value. The FILES_NAMES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetFILESNAMES()(*string) {
    return m.fILES_NAMES
}
// GetGEO gets the GEO property value. The GEO property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetGEO()(*string) {
    return m.gEO
}
// GetQUERYSTRING gets the QUERY_STRING property value. The QUERY_STRING property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetQUERYSTRING()(*string) {
    return m.qUERY_STRING
}
// GetREMOTEADDR gets the REMOTE_ADDR property value. The REMOTE_ADDR property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREMOTEADDR()(*string) {
    return m.rEMOTE_ADDR
}
// GetREQUESTBASENAME gets the REQUEST_BASENAME property value. The REQUEST_BASENAME property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTBASENAME()(*string) {
    return m.rEQUEST_BASENAME
}
// GetREQUESTBODY gets the REQUEST_BODY property value. The REQUEST_BODY property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTBODY()(*string) {
    return m.rEQUEST_BODY
}
// GetREQUESTCOOKIES gets the REQUEST_COOKIES property value. The REQUEST_COOKIES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTCOOKIES()(*string) {
    return m.rEQUEST_COOKIES
}
// GetREQUESTCOOKIESNAMES gets the REQUEST_COOKIES_NAMES property value. The REQUEST_COOKIES_NAMES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTCOOKIESNAMES()(*string) {
    return m.rEQUEST_COOKIES_NAMES
}
// GetREQUESTFILENAME gets the REQUEST_FILENAME property value. The REQUEST_FILENAME property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTFILENAME()(*string) {
    return m.rEQUEST_FILENAME
}
// GetREQUESTHEADERS gets the REQUEST_HEADERS property value. The REQUEST_HEADERS property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTHEADERS()(*string) {
    return m.rEQUEST_HEADERS
}
// GetREQUESTHEADERSNAMES gets the REQUEST_HEADERS_NAMES property value. The REQUEST_HEADERS_NAMES property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTHEADERSNAMES()(*string) {
    return m.rEQUEST_HEADERS_NAMES
}
// GetREQUESTLINE gets the REQUEST_LINE property value. The REQUEST_LINE property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTLINE()(*string) {
    return m.rEQUEST_LINE
}
// GetREQUESTMETHOD gets the REQUEST_METHOD property value. The REQUEST_METHOD property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTMETHOD()(*string) {
    return m.rEQUEST_METHOD
}
// GetREQUESTPROTOCOL gets the REQUEST_PROTOCOL property value. The REQUEST_PROTOCOL property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTPROTOCOL()(*string) {
    return m.rEQUEST_PROTOCOL
}
// GetREQUESTURI gets the REQUEST_URI property value. The REQUEST_URI property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTURI()(*string) {
    return m.rEQUEST_URI
}
// GetREQUESTURIRAW gets the REQUEST_URI_RAW property value. The REQUEST_URI_RAW property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetREQUESTURIRAW()(*string) {
    return m.rEQUEST_URI_RAW
}
// GetRESPONSEBODY gets the RESPONSE_BODY property value. The RESPONSE_BODY property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetRESPONSEBODY()(*string) {
    return m.rESPONSE_BODY
}
// GetRESPONSEHEADERS gets the RESPONSE_HEADERS property value. The RESPONSE_HEADERS property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetRESPONSEHEADERS()(*string) {
    return m.rESPONSE_HEADERS
}
// GetRESPONSESTATUS gets the RESPONSE_STATUS property value. The RESPONSE_STATUS property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel_variableTypes) GetRESPONSESTATUS()(*string) {
    return m.rESPONSE_STATUS
}
// Serialize serializes information the current object
func (m *CreateWafRateLimitRuleModel_variableTypes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ARGS", m.GetARGS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ARGS_COMBINED_SIZE", m.GetARGSCOMBINEDSIZE())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ARGS_GET", m.GetARGSGET())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ARGS_GET_NAMES", m.GetARGSGETNAMES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ARGS_POST", m.GetARGSPOST())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ARGS_POST_NAMES", m.GetARGSPOSTNAMES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("FILES_NAMES", m.GetFILESNAMES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("GEO", m.GetGEO())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("QUERY_STRING", m.GetQUERYSTRING())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REMOTE_ADDR", m.GetREMOTEADDR())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_BASENAME", m.GetREQUESTBASENAME())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_BODY", m.GetREQUESTBODY())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_COOKIES", m.GetREQUESTCOOKIES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_COOKIES_NAMES", m.GetREQUESTCOOKIESNAMES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_FILENAME", m.GetREQUESTFILENAME())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_HEADERS", m.GetREQUESTHEADERS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_HEADERS_NAMES", m.GetREQUESTHEADERSNAMES())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_LINE", m.GetREQUESTLINE())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_METHOD", m.GetREQUESTMETHOD())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_PROTOCOL", m.GetREQUESTPROTOCOL())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_URI", m.GetREQUESTURI())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("REQUEST_URI_RAW", m.GetREQUESTURIRAW())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("RESPONSE_BODY", m.GetRESPONSEBODY())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("RESPONSE_HEADERS", m.GetRESPONSEHEADERS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("RESPONSE_STATUS", m.GetRESPONSESTATUS())
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
func (m *CreateWafRateLimitRuleModel_variableTypes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetARGS sets the ARGS property value. The ARGS property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGS(value *string)() {
    m.aRGS = value
}
// SetARGSCOMBINEDSIZE sets the ARGS_COMBINED_SIZE property value. The ARGS_COMBINED_SIZE property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGSCOMBINEDSIZE(value *string)() {
    m.aRGS_COMBINED_SIZE = value
}
// SetARGSGET sets the ARGS_GET property value. The ARGS_GET property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGSGET(value *string)() {
    m.aRGS_GET = value
}
// SetARGSGETNAMES sets the ARGS_GET_NAMES property value. The ARGS_GET_NAMES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGSGETNAMES(value *string)() {
    m.aRGS_GET_NAMES = value
}
// SetARGSPOST sets the ARGS_POST property value. The ARGS_POST property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGSPOST(value *string)() {
    m.aRGS_POST = value
}
// SetARGSPOSTNAMES sets the ARGS_POST_NAMES property value. The ARGS_POST_NAMES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetARGSPOSTNAMES(value *string)() {
    m.aRGS_POST_NAMES = value
}
// SetFILESNAMES sets the FILES_NAMES property value. The FILES_NAMES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetFILESNAMES(value *string)() {
    m.fILES_NAMES = value
}
// SetGEO sets the GEO property value. The GEO property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetGEO(value *string)() {
    m.gEO = value
}
// SetQUERYSTRING sets the QUERY_STRING property value. The QUERY_STRING property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetQUERYSTRING(value *string)() {
    m.qUERY_STRING = value
}
// SetREMOTEADDR sets the REMOTE_ADDR property value. The REMOTE_ADDR property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREMOTEADDR(value *string)() {
    m.rEMOTE_ADDR = value
}
// SetREQUESTBASENAME sets the REQUEST_BASENAME property value. The REQUEST_BASENAME property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTBASENAME(value *string)() {
    m.rEQUEST_BASENAME = value
}
// SetREQUESTBODY sets the REQUEST_BODY property value. The REQUEST_BODY property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTBODY(value *string)() {
    m.rEQUEST_BODY = value
}
// SetREQUESTCOOKIES sets the REQUEST_COOKIES property value. The REQUEST_COOKIES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTCOOKIES(value *string)() {
    m.rEQUEST_COOKIES = value
}
// SetREQUESTCOOKIESNAMES sets the REQUEST_COOKIES_NAMES property value. The REQUEST_COOKIES_NAMES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTCOOKIESNAMES(value *string)() {
    m.rEQUEST_COOKIES_NAMES = value
}
// SetREQUESTFILENAME sets the REQUEST_FILENAME property value. The REQUEST_FILENAME property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTFILENAME(value *string)() {
    m.rEQUEST_FILENAME = value
}
// SetREQUESTHEADERS sets the REQUEST_HEADERS property value. The REQUEST_HEADERS property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTHEADERS(value *string)() {
    m.rEQUEST_HEADERS = value
}
// SetREQUESTHEADERSNAMES sets the REQUEST_HEADERS_NAMES property value. The REQUEST_HEADERS_NAMES property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTHEADERSNAMES(value *string)() {
    m.rEQUEST_HEADERS_NAMES = value
}
// SetREQUESTLINE sets the REQUEST_LINE property value. The REQUEST_LINE property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTLINE(value *string)() {
    m.rEQUEST_LINE = value
}
// SetREQUESTMETHOD sets the REQUEST_METHOD property value. The REQUEST_METHOD property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTMETHOD(value *string)() {
    m.rEQUEST_METHOD = value
}
// SetREQUESTPROTOCOL sets the REQUEST_PROTOCOL property value. The REQUEST_PROTOCOL property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTPROTOCOL(value *string)() {
    m.rEQUEST_PROTOCOL = value
}
// SetREQUESTURI sets the REQUEST_URI property value. The REQUEST_URI property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTURI(value *string)() {
    m.rEQUEST_URI = value
}
// SetREQUESTURIRAW sets the REQUEST_URI_RAW property value. The REQUEST_URI_RAW property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetREQUESTURIRAW(value *string)() {
    m.rEQUEST_URI_RAW = value
}
// SetRESPONSEBODY sets the RESPONSE_BODY property value. The RESPONSE_BODY property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetRESPONSEBODY(value *string)() {
    m.rESPONSE_BODY = value
}
// SetRESPONSEHEADERS sets the RESPONSE_HEADERS property value. The RESPONSE_HEADERS property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetRESPONSEHEADERS(value *string)() {
    m.rESPONSE_HEADERS = value
}
// SetRESPONSESTATUS sets the RESPONSE_STATUS property value. The RESPONSE_STATUS property
func (m *CreateWafRateLimitRuleModel_variableTypes) SetRESPONSESTATUS(value *string)() {
    m.rESPONSE_STATUS = value
}
type CreateWafRateLimitRuleModel_variableTypesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetARGS()(*string)
    GetARGSCOMBINEDSIZE()(*string)
    GetARGSGET()(*string)
    GetARGSGETNAMES()(*string)
    GetARGSPOST()(*string)
    GetARGSPOSTNAMES()(*string)
    GetFILESNAMES()(*string)
    GetGEO()(*string)
    GetQUERYSTRING()(*string)
    GetREMOTEADDR()(*string)
    GetREQUESTBASENAME()(*string)
    GetREQUESTBODY()(*string)
    GetREQUESTCOOKIES()(*string)
    GetREQUESTCOOKIESNAMES()(*string)
    GetREQUESTFILENAME()(*string)
    GetREQUESTHEADERS()(*string)
    GetREQUESTHEADERSNAMES()(*string)
    GetREQUESTLINE()(*string)
    GetREQUESTMETHOD()(*string)
    GetREQUESTPROTOCOL()(*string)
    GetREQUESTURI()(*string)
    GetREQUESTURIRAW()(*string)
    GetRESPONSEBODY()(*string)
    GetRESPONSEHEADERS()(*string)
    GetRESPONSESTATUS()(*string)
    SetARGS(value *string)()
    SetARGSCOMBINEDSIZE(value *string)()
    SetARGSGET(value *string)()
    SetARGSGETNAMES(value *string)()
    SetARGSPOST(value *string)()
    SetARGSPOSTNAMES(value *string)()
    SetFILESNAMES(value *string)()
    SetGEO(value *string)()
    SetQUERYSTRING(value *string)()
    SetREMOTEADDR(value *string)()
    SetREQUESTBASENAME(value *string)()
    SetREQUESTBODY(value *string)()
    SetREQUESTCOOKIES(value *string)()
    SetREQUESTCOOKIESNAMES(value *string)()
    SetREQUESTFILENAME(value *string)()
    SetREQUESTHEADERS(value *string)()
    SetREQUESTHEADERSNAMES(value *string)()
    SetREQUESTLINE(value *string)()
    SetREQUESTMETHOD(value *string)()
    SetREQUESTPROTOCOL(value *string)()
    SetREQUESTURI(value *string)()
    SetREQUESTURIRAW(value *string)()
    SetRESPONSEBODY(value *string)()
    SetRESPONSEHEADERS(value *string)()
    SetRESPONSESTATUS(value *string)()
}
