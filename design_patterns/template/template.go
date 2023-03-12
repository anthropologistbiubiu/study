package main

import "fmt"

type IOtp interface {
	genRandomOtp(length int) string
	saveOtpCache(string) error
	getMessage(string) string
	sendNotifaction(string) error
}
type Otp struct {
	IOtp
}

func (this *Otp) genAndSendOtp(length int) bool {
	opt := this.IOtp.genRandomOtp(length)
	err := this.saveOtpCache(opt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	message := this.IOtp.getMessage(opt)
	err = this.sendNotifaction(message)
	if err != nil {
		return false
	}
	return true
}

type emailOtp struct {
	Otp
}

func (this *emailOtp) genRandomOtp(lenghth int) string {
	fmt.Println("get email random otp")
	return fmt.Sprintf("email otp length = %d", lenghth)
}
func (this *emailOtp) saveOtpCache(opt string) error {
	fmt.Println("save emailOpt Cache")
	return nil
}
func (this *emailOtp) getMessage(opt string) string {
	fmt.Println("email opt get ready")
	return "email opt message is " + opt
}
func (this *emailOtp) sendNotifaction(message string) error {
	fmt.Printf("login email otp have send  + message = %s\n", message)
	return nil
}

type smsOtp struct {
	Otp
}

func (this *smsOtp) genRandomOtp(lenght int) string {
	fmt.Println("get sms random otp")
	return fmt.Sprintf("sms opt length = %d", lenght)
}
func (this *smsOtp) saveOtpCache(opt string) error {
	fmt.Println("save smsOpt Cache")
	return nil
}
func (this *smsOtp) getMessage(opt string) string {
	fmt.Println("sms opt message ready")
	return "sms opt message is " + opt
}
func (this *smsOtp) sendNotifaction(message string) error {
	fmt.Printf("login sms otp have send + message = %s\n", message)
	return nil
}
func main() {
	email := &emailOtp{}
	emailOtp := &Otp{
		IOtp: email,
	}

	emailOtp.genAndSendOtp(4)

	sms := &smsOtp{}
	smsOtp := &Otp{
		IOtp: sms,
	}
	smsOtp.genAndSendOtp(5)

}
