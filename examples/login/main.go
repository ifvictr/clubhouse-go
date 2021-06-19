package main

import (
	"fmt"

	"github.com/ifvictr/clubhouse-go"
)

func main() {
	client := clubhouse.New()

	var phoneNumber string
	fmt.Print("Phone # (with country code): ")
	fmt.Scanf("%s", &phoneNumber)
	{
		res, _, err := client.StartPhoneNumberAuth(&clubhouse.StartPhoneNumberAuthParams{PhoneNumber: phoneNumber})
		if err != nil {
			fmt.Printf("Failed to start auth: %s", err)
			return
		}
		if !res.Success {
			fmt.Println("Could not start auth")
			return
		}
		if res.IsBlocked {
			fmt.Println("This number is blocked")
			return
		}
	}
	fmt.Println("Verification code sent")

	for {
		var verificationCode string
		fmt.Print("Code: ")
		fmt.Scanf("%s", &verificationCode)
		res, _, err := client.CompletePhoneNumberAuth(&clubhouse.CompletePhoneNumberAuthParams{
			PhoneNumber:      phoneNumber,
			VerificationCode: verificationCode,
		})
		if err != nil {
			fmt.Printf("Failed to complete auth: %s", err)
			return
		}
		if !res.Success {
			fmt.Println("Could not complete auth")
			return
		}
		if res.NumberOfAttemptsRemaining != nil {
			fmt.Printf("You have %d attempts remaining\n", *res.NumberOfAttemptsRemaining)
			if *res.NumberOfAttemptsRemaining == 0 {
				return
			}
			continue
		}

		fmt.Printf("Logged in to @%s\n", *res.UserProfile.Username)
		fmt.Printf("Auth token: %s\n", *res.AuthToken)
		return
	}
}
