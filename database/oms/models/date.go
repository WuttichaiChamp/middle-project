package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const DateLayout = "2006-01-02"

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // remove quotes
	
	if s == "" || s == "null" {
		return nil
	}
	
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return err
	}
	
	d.Time = t
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format(DateLayout))), nil
}

func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	if t, ok := value.(time.Time); ok {
		d.Time = t
		return nil
	}
	return fmt.Errorf("cannot scan %T into Date", value)
}
