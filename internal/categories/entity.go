package categories

import (
	"errors"
)

type Category struct {
	Perishable    string `json:"perishable"`
	NonPerishable string `json:"non_perishable"`
}

func _(c *Category) (error, bool) {
	if c.Perishable == "" || c.NonPerishable == "" {
		return errors.New("it is necessary category type"), false
	}
	return nil, true
}
