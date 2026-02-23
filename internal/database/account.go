package database

import "gorm.io/gorm"

type Account struct {
        gorm.Model
        HouseholdID        *uint  `gorm:"index"`
        Email              string `gorm:"uniqueIndex;not null"`
        Name               string `gorm:"not null"`
        Locale             string `gorm:"not null;default:'en'"`
        Active             bool   `gorm:"not null;default:true"`
        Verified           bool   `gorm:"not null;default:false"`
        OnboardingComplete bool   `gorm:"not null;default:false"`
}
