package leetcode

func findSmallestRegion(source [][]RegionName, region1 RegionName, region2 RegionName) string {
	regions := mapRegions(source)
}

type RegionName string

// maps each non-root region to its direct parent
type Regions map[RegionName]RegionName

func mapRegions(source [][]RegionName) *Regions {
	regions := make(Regions)

	for _, region := range source {
		for _, child := range region[1:] {
			regions[child] = region[0]
		}
	}

	return &regions
}
