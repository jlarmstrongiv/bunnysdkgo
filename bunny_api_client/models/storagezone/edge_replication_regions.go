package storagezone
type EdgeReplicationRegions int

const (
    DE_EDGEREPLICATIONREGIONS EdgeReplicationRegions = iota
    UK_EDGEREPLICATIONREGIONS
    SE_EDGEREPLICATIONREGIONS
    CZ_EDGEREPLICATIONREGIONS
    ES_EDGEREPLICATIONREGIONS
    NY_EDGEREPLICATIONREGIONS
    LA_EDGEREPLICATIONREGIONS
    WA_EDGEREPLICATIONREGIONS
    MI_EDGEREPLICATIONREGIONS
    SG_EDGEREPLICATIONREGIONS
    HK_EDGEREPLICATIONREGIONS
    JP_EDGEREPLICATIONREGIONS
    SYD_EDGEREPLICATIONREGIONS
    BR_EDGEREPLICATIONREGIONS
    JH_EDGEREPLICATIONREGIONS
)

func (i EdgeReplicationRegions) String() string {
    return []string{"DE", "UK", "SE", "CZ", "ES", "NY", "LA", "WA", "MI", "SG", "HK", "JP", "SYD", "BR", "JH"}[i]
}
func ParseEdgeReplicationRegions(v string) (any, error) {
    result := DE_EDGEREPLICATIONREGIONS
    switch v {
        case "DE":
            result = DE_EDGEREPLICATIONREGIONS
        case "UK":
            result = UK_EDGEREPLICATIONREGIONS
        case "SE":
            result = SE_EDGEREPLICATIONREGIONS
        case "CZ":
            result = CZ_EDGEREPLICATIONREGIONS
        case "ES":
            result = ES_EDGEREPLICATIONREGIONS
        case "NY":
            result = NY_EDGEREPLICATIONREGIONS
        case "LA":
            result = LA_EDGEREPLICATIONREGIONS
        case "WA":
            result = WA_EDGEREPLICATIONREGIONS
        case "MI":
            result = MI_EDGEREPLICATIONREGIONS
        case "SG":
            result = SG_EDGEREPLICATIONREGIONS
        case "HK":
            result = HK_EDGEREPLICATIONREGIONS
        case "JP":
            result = JP_EDGEREPLICATIONREGIONS
        case "SYD":
            result = SYD_EDGEREPLICATIONREGIONS
        case "BR":
            result = BR_EDGEREPLICATIONREGIONS
        case "JH":
            result = JH_EDGEREPLICATIONREGIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEdgeReplicationRegions(values []EdgeReplicationRegions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EdgeReplicationRegions) isMultiValue() bool {
    return false
}
