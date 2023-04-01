package enums

func GenerateAuthEnums () map[int]string {
 	authEnum := map[int]string {
		1: "READ",
		2: "WRITE",
		4: "UPDATE",
		8: "DELETE",
		16: "BULK_WRITE",
		32: "BULK_UPDATE",
		64: "BULK_DELETE",
	}

	return authEnum
}

func GetPermissionArrayFromAuthLevel (level int) []string {
	permissionArray := []string{}
	authEnum := GenerateAuthEnums()
	
	nearestPowerOfTwo, index := GetPowerOfTwoJustBelowNum(level);
	num := level;

	for num > 0 {
		permissionArray = append(permissionArray, authEnum[nearestPowerOfTwo])
		num -= nearestPowerOfTwo
		nearestPowerOfTwo = PowersOfTwo[index -1]
		index--
	}

	return permissionArray
}
