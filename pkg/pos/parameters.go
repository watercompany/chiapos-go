package pos

const (
	paramEXT = 5

	paramM = 1 << paramEXT

	paramB = 60

	paramC = 509

	paramBC = paramB * paramC

	paramC1 = 10000
	paramC2 = 10000

	paramStubBits = 4
)

func BucketID(x uint64) uint64 {
	return x / paramBC
}

func GetIDs(x uint64) (bID, cID uint64) {
	y := x % paramBC
	return y / paramC, y % paramC
}