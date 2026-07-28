package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	const (
		minutesPerHour = 60
		minutesPerDay  = 24 * minutesPerHour
	)

	totalMinutes := h*minutesPerHour + m
	signedDayMinutes := totalMinutes % minutesPerDay
	clockMinutes := (signedDayMinutes + minutesPerDay) % minutesPerDay

	return Clock{hours: clockMinutes / minutesPerHour, minutes: clockMinutes % minutesPerHour}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
