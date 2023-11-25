package shared

import "time"

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (c UpdatedAt) Value() time.Time {
	return time.Time(c)
}

func (c UpdatedAt) Equal(c2 UpdatedAt) bool {
	return c.Value().Equal(c2.Value())
}
