package enums

func GenerateAuthEnums() map[int]string {
	authEnum := map[int]string{
		1:   "READ",
		2:   "BULK_READ",
		4:   "WRITE",
		8:   "BULK_WRITE",
		16:  "UPDATE",
		32:  "BULK_UPDATE",
		64:  "DELETE",
		128: "BULK_DELETE",
	}

	return authEnum
}

func GetPermissionArrayFromAuthLevel(level int) []string {
	permissionArray := []string{}
	authEnum := GenerateAuthEnums()

	nearestPowerOfTwo, index := GetPowerOfTwoJustBelowNum(level)
	num := level

	for num > 0 {
		permissionArray = append(permissionArray, authEnum[nearestPowerOfTwo])
		num -= nearestPowerOfTwo
		if index > 0 {
			index--
			nearestPowerOfTwo = PowersOfTwo[index]
		}
	}

	return permissionArray
}
