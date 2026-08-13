func findSmallestRegion(source [][]string, region1 string, region2 string) string {
	regions := mapRegions(source)

	return string(regions.LCA(RegionName(region1), RegionName(region2)))
}

type RegionName string

// maps each non-root region to its direct parent
type Regions map[RegionName]RegionName

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Has(val T) bool {
	_, exists := s[val]
	return exists
}

func mapRegions(source [][]string) *Regions {
	regions := make(Regions)

	for _, region := range source {
		for _, child := range region[1:] {
			regions[RegionName(child)] = RegionName(region[0])
		}
	}

	return &regions
}

func (r Regions) LCA(r1, r2 RegionName) RegionName {
	r1Parents := make(Set[RegionName])

	for parent := r1; parent != ""; parent = r[parent] {
		r1Parents.Add(parent)
	}

	result := r2
	for !r1Parents.Has(result) {
		result = r[result]
	}

	return result
}

func findRoot(regions *Regions) RegionName {
	for region, parent := range *regions {
		if parent == "" {
			return region
		}
	}

	panic("cannot find root")
}
