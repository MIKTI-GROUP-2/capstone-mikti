package customtime

import "time"

type CustomTime struct {
	time.Time
}

const customTimeFormat = "17-07-1945"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	t, err := time.Parse(customTimeFormat, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(customTimeFormat) + `"`), nil
}

func CustomTimeFromString(s string) CustomTime {
	t, _ := time.Parse(customTimeFormat, s)
	return CustomTime{Time: t}
}
