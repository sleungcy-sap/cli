package constant

type UserRole int

const (
	OrgManager UserRole = iota
	BillingManager
	OrgAuditor
	SpaceManager
	SpaceDeveloper
	SpaceAuditor
)
