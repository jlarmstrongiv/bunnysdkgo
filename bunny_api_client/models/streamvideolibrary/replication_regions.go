package streamvideolibrary
type ReplicationRegions int

const (
    UK_REPLICATIONREGIONS ReplicationRegions = iota
    SE_REPLICATIONREGIONS
    NY_REPLICATIONREGIONS
    LA_REPLICATIONREGIONS
    SG_REPLICATIONREGIONS
    SYD_REPLICATIONREGIONS
    BR_REPLICATIONREGIONS
    JH_REPLICATIONREGIONS
)

func (i ReplicationRegions) String() string {
    return []string{"UK", "SE", "NY", "LA", "SG", "SYD", "BR", "JH"}[i]
}
func ParseReplicationRegions(v string) (any, error) {
    result := UK_REPLICATIONREGIONS
    switch v {
        case "UK":
            result = UK_REPLICATIONREGIONS
        case "SE":
            result = SE_REPLICATIONREGIONS
        case "NY":
            result = NY_REPLICATIONREGIONS
        case "LA":
            result = LA_REPLICATIONREGIONS
        case "SG":
            result = SG_REPLICATIONREGIONS
        case "SYD":
            result = SYD_REPLICATIONREGIONS
        case "BR":
            result = BR_REPLICATIONREGIONS
        case "JH":
            result = JH_REPLICATIONREGIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReplicationRegions(values []ReplicationRegions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReplicationRegions) isMultiValue() bool {
    return false
}
