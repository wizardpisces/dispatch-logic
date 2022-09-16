package structures

// type PortalEnum string

// const (
// 	Seller PortalEnum = "Seller"
// 	Normal PortalEnum = "Normal"
// )
type PortalEnum int

const (
	Seller PortalEnum = iota
	Normal
)

func (p PortalEnum) String() string {
	switch p {
	case Seller:
		return "Seller"
	case Normal:
		return "Normal"
	default:
		return "Unknown"
	}
}
