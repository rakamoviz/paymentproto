package proto

import (
	"io/ioutil"
	"log"

	payment "github.com/rakamoviz/paymentproto/payment"
)

func init() {
	log.SetOutput(ioutil.Discard)
	log.Println(payment.PaymentReceived{})
}
