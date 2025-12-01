package types

import "time"

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	format := `"2006-01-02T15:04:05"`
	if b[len(b)-2] == 'Z' {
		format = `"2006-01-02T15:04:05Z07:00"`
	}

	date, err := time.Parse(format, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
