package main

import "fmt"

type client struct {
}

func (c *client) insertHeadphonesIntoPhone(ph phone) {
	fmt.Println("Client inserts headphones into mobile phone.")
	ph.insertIntoHeadphonePort()
}

type phone interface {
	insertIntoHeadphonePort()
}

type iphone struct {
}

func (iph *iphone) insertIntoHeadphonePort() {
	fmt.Println("Airpods is plugged into iphone.")
}

type android struct {
}

func (andr *android) insertIntoAndroidHeadphonePort() {
	fmt.Println("Common headphones plugged into android phone")
}

type androidAdapter struct {
	androidPhone *android
}

func (phn *androidAdapter) insertIntoHeadphonePort() {
	fmt.Println("Adapter converts airpods to common headphones.")
	phn.androidPhone.insertIntoAndroidHeadphonePort()
}

func main() {

	client := &client{}
	iphone := &iphone{}

	client.insertHeadphonesIntoPhone(iphone)
	androidPhone := &android{}
	androidPhoneAdapter := &androidAdapter{
		androidPhone: androidPhone,
	}

	client.insertHeadphonesIntoPhone(androidPhoneAdapter)
}
