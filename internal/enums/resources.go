package enums

type Resource struct {
	Name          string
	Description   string
	TotalLevelSum int
}

type AppResourcesListType []Resource

var AppResourcesList = AppResourcesListType{
	Resource{
		Name:          "User",
		Description:   "Auth users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Profile",
		Description:   "Profile of users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Address",
		Description:   "Addresses of users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Permission",
		Description:   "Permissions of the users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Role",
		Description:   "Roles of the users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Availability",
		Description:   "Availability of the users in the domain",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Leave",
		Description:   "Leave applications by the users",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Appointment",
		Description:   "Appointment of two users of different roles",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Consumables",
		Description:   "All the items consumed by a patient or the hospital",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Non Consumables",
		Description:   "All the items used by the hospital which do not get depleted on use",
		TotalLevelSum: 255,
	},
	Resource{
		Name:          "Prescription",
		Description:   "All the prescriptions created by the hospital authority",
		TotalLevelSum: 255,
	},
}
