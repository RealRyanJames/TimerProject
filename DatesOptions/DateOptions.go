package datesoptions

import "time"

type DateOptions struct {
	Date time.Time
}

func (dayTimeHour DateOptions) Get() DateOptions {

	return DateOptions{
		Date: time.Now(),
	}
}
