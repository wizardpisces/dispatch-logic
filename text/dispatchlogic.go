package text

import (
	portal "github.com/wizardpisces/dispatch-logic/text/portal"

	"github.com/wizardpisces/dispatch-logic/structures"
)

type Portal struct {
}

// type PortalRegion struct {
// }

// Hello returns a greeting for the named person.
func (p Portal) Text(portalName structures.PortalEnum, name string) string {
	// Return a greeting that embeds the name in a message.
	if portalName == structures.Seller {
		return portal.HelloSeller(name)
	}
	return portal.HelloNormal(name)
}

// // Hello returns a greeting for the named person.
// func (p PortalRegion) PortalRegionText(portalName structures.PortalEnum, region string, text string) string {

// 	return fmt.Sprintf("%s,%s,%s", portalName, region, text)
// }
