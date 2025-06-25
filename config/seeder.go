package config

import "Siberat/entity"

func SeedRoles() {
	roles := []entity.RoleUser{
		{Role: "Petugas Penetapan"},
		{Role: "Petugas Bidang PP"},
		{Role: "Petugas PSIP"},
		{Role: "Pejabat Pengawas"},
	}

	for _, role := range roles {
		DB.FirstOrCreate(&role, entity.RoleUser{Role: role.Role})
	}
}