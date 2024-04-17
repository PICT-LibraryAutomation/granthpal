package utils

func Map[T any, U any](ts []T, f func(T) U) []U {
	var us []U
	for _, t := range ts {
		us = append(us, f(t))
	}
	return us
}
