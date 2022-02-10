package provider

import "gorm.io/gorm"

type WebserviceDI struct {
	Gorm *gorm.DB
}

func GetWebserviceDI() WebserviceDI {
	return WebserviceDI{
		Gorm: GetGormMain(),
	}
}
