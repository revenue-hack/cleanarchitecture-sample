package shared

import "time"

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (c CreatedAt) Value() time.Time {
	return time.Time(c)
}

func (c CreatedAt) Equal(c2 CreatedAt) bool {
	return c.Value().Equal(c2.Value())
}
