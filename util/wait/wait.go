package wait

import "fmt"

// Until iterates over buffered error channel and:
// * upon receiving non-nil value from the channel, makes an early return with this value
// * if no non-nil values were received from iteration over the channel, it just returns nil
// Used to wait for a series of goroutines, launched altogether from the same loop, to finish.
func Until(done chan error) error {
	i := 0

	for err := range done {
		if err != nil {
			return err
		}

		i++

		if i >= cap(done) {
			close(done)
		}
	}

	return nil
}

// WithTolerance is the same as Until, but it does not return immediately on errors
// rather loops through all channel capacity returning "composite" error in the end.
func WithTolerance(done chan error) error {
	var errMessage string

	i := 0

	for err := range done {
		if err != nil {
			errMessage = errMessage + err.Error() + "\n"
		}

		i++

		if i >= cap(done) {
			close(done)
		}
	}

	if len(errMessage) == 0 {
		return nil
	}

	return fmt.Errorf(errMessage)
}
