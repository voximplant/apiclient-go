package structure

import "time"

type Timestamp time.Time
//type Date time.Time

type Date struct {
	time.Time
}

// UnmarshalJSON will transform the JIRA date into a time.Time
// during the transformation of the JIRA JSON response
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse("\"2006-01-02 15:04:05\"", string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(ti)
	return nil
}

// MarshalJSON will transform the Date object into a short
// date string as JIRA expects during the creation of a
// JIRA request
func (t Timestamp) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(t)
	return []byte(timestamp.Format("\"2006-01-02 15:04:05\"")), nil
}

// UnmarshalJSON will transform the JIRA date into a time.Time
// during the transformation of the JIRA JSON response
func (t *Date) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*t = Date{ti}
	return nil
}

var nilTime = (time.Time{}).UnixNano()
// MarshalJSON will transform the Date object into a short
// date string as JIRA expects during the creation of a
// JIRA request
func (t Date) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	date := time.Time(t.Time)
	return []byte(date.Format("\"2006-01-02\"")), nil
}