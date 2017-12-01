package pay

import (
	stripe "github.com/stripe/stripe-go"
)

const (
	Key = "sk_test_t1FcBDEKGVFkCoYBnGJJv6HH"
)

func GetInit() {
	params := &stripe.Charge{
		"id": "",
	}
}
