package handler

import "fmt"

type CreateOpeningRequest struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	WorkModel string `json:"workModel"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" {
		return validateAttribute("role")
	}

	if c.Company == "" {
		return validateAttribute("company")
	}

	if c.Location == "" {
		return validateAttribute("location")
	}

	if c.WorkModel == "" {
		return validateAttribute("workModel")
	}

	if c.Link == "" {
		return validateAttribute("link")
	}

	if c.Salary <= 0 {
		return validateAttribute("salary")
	}

	return nil
}

func validateAttribute(attrName string) error {
	return fmt.Errorf("Param '%s' is required!", attrName)
}

type UpdateOpeningRequest struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	WorkModel string `json:"workModel"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
}

func (u *UpdateOpeningRequest) Validate() error {
	if u.Role != "" || u.Company != "" || u.Location != "" || u.WorkModel != "" || u.Link != "" || u.Salary > 0 {
		return nil
	}

	return fmt.Errorf("At least one valid field must be provided")
}
