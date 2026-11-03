// Package services holds only functions: the other half of the same context.
package services

import "github.com/victorhugomazzeo/100_go_mistakes/012-project-misorganization/mistake/models"

func Customers() []models.Customer {
	return []models.Customer{{Name: "Gopher"}}
}
