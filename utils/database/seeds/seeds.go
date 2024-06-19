package seeds

import (
	"capstone-mikti/utils/database/seed"

	"gorm.io/gorm"
)

func All() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "Create Admin",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "admin", "admin@gmail.com", "012345679")
			},
		},
	}
	return seeds
}
