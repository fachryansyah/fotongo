package services

import (
	"fmt"
	"fotongo/app/utils/constant"
	"fotongo/domains/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitDatabase(dbConfig postgres.Config) *gorm.DB {
	if dbInstance == nil {
		var err error

		NewDBInstance, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			//panic(vm.GetError(discordPackage.WhereAmI(), err))
		}

		if constant.IsMigrate {
			err := NewDBInstance.AutoMigrate(&entities.User{})
			if err != nil {
				panic(err)
			}
		}

		//if constant.IsMigrate {
		//	err := NewDBInstance.AutoMigrate(&entities.HomecareCustomer{}, &entities.HomecarePatient{},
		//		&entities.PatientMedicalHistory{}, &entities.PatientMedicalCondition{},
		//		&entities.PatientMedicalState{},
		//		&entities.CustomerOrderHistory{}, &entities.CustomerOrderEventHistory{},
		//		&entities.CustomerOrderHistoryPatient{}, &entities.HomecareNotification{}, &entities.PatientMedicalEquipment{}, &entities.HomecareOrder{})
		//
		//	if err != nil {
		//		panic(vm.GetError(discordPackage.WhereAmI(), err))
		//	}
		//	fmt.Println("Migrate Success")
		//}
		fmt.Println(dbConfig.DSN)

		dbInstance = NewDBInstance

		return dbInstance
	} else {
		return dbInstance
	}
}
