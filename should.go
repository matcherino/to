package to

import "time"

// ShouldBool converts v to bool or returns false.
func ShouldBool(v interface{}) bool {
	b, err := Bool(v)
	if err != nil {
		return false
	}
	return b
}

// ShouldDuration converts v to Duration or returns Duration(0).
func ShouldDuration(v interface{}) time.Duration {
	d, err := Duration(v)
	if err != nil {
		return time.Duration(0)
	}
	return d
}

// ShouldInt converts v to int64 or returns 0.
func ShouldInt(v interface{}) int {
	i, err := Int64(v)
	if err != nil {
		return 0
	}
	return int(i)
}

// ShouldInt64 converts v to int64 or returns 0.
func ShouldInt64(v interface{}) int64 {
	i, err := Int64(v)
	if err != nil {
		return 0
	}
	return i
}

// ShouldFloat converts v to float64 or returns float64(0)
func ShouldFloat(v interface{}) float64 {
	f, err := Float64(v)
	if err != nil {
		return 0
	}
	return f
}

// ShouldString converts v to string or returns ""
func ShouldString(v interface{}) string {
	return String(v)
}

// ShouldTime converts v to Time or returns Time{}
func ShouldTime(v interface{}) time.Time {
	t, err := Time(v)
	if err != nil {
		return time.Time{}
	}
	return t
}
