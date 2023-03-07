package user

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"

	"github.com/ardanlabs/service/business/sys/validate"
	"github.com/google/uuid"
)

// Used to check for sql injection problems.
var sqlInjection = regexp.MustCompile("^[A-Za-z0-9_]+$")

// =============================================================================

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID    *uuid.UUID    `validate:"omitempty,uuid4"`
	Name  *string       `validate:"omitempty,min=3"`
	Email *mail.Address `validate:"omitempty,email"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithUserID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithUserID(userID uuid.UUID) {
	var zero uuid.UUID
	if userID != zero {
		qf.ID = &userID
	}
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) error {
	if name != "" {
		if !sqlInjection.MatchString(name) {
			return errors.New("invalid name format")
		}

		qf.Name = &name
	}

	return nil
}

// WithEmail sets the Email field of the QueryFilter value.
func (qf *QueryFilter) WithEmail(email mail.Address) {
	var zero mail.Address
	if email != zero {
		qf.Email = &email
	}
}
