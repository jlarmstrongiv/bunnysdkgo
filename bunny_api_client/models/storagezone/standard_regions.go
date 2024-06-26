package storagezone
type StandardRegions int

const (
    DE_STANDARDREGIONS StandardRegions = iota
    UK_STANDARDREGIONS
    SE_STANDARDREGIONS
    NY_STANDARDREGIONS
    LA_STANDARDREGIONS
    SG_STANDARDREGIONS
    SYD_STANDARDREGIONS
    BR_STANDARDREGIONS
    JH_STANDARDREGIONS
)

func (i StandardRegions) String() string {
    return []string{"DE", "UK", "SE", "NY", "LA", "SG", "SYD", "BR", "JH"}[i]
}
func ParseStandardRegions(v string) (any, error) {
    result := DE_STANDARDREGIONS
    switch v {
        case "DE":
            result = DE_STANDARDREGIONS
        case "UK":
            result = UK_STANDARDREGIONS
        case "SE":
            result = SE_STANDARDREGIONS
        case "NY":
            result = NY_STANDARDREGIONS
        case "LA":
            result = LA_STANDARDREGIONS
        case "SG":
            result = SG_STANDARDREGIONS
        case "SYD":
            result = SYD_STANDARDREGIONS
        case "BR":
            result = BR_STANDARDREGIONS
        case "JH":
            result = JH_STANDARDREGIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStandardRegions(values []StandardRegions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StandardRegions) isMultiValue() bool {
    return false
}
