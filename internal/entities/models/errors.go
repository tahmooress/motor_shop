package models

import "github.com/tahmooress/motor-shop/internal/pkg/customeerror"

var (
	ErrEmptyEnvironment = customeerror.New("1000", "environment variables must be not empty",
		"پارامترهای وروری برنامه خالی هستند")
	ErrEmptyIPANDPORT = customeerror.New("1001", "some fields are empty in cli: IP, Port, routers",
		"پارامترهای آی پی و پورت وارد نشده است")
	ErrUnknown                 = customeerror.New("1002", "Unknown error", "ارور ناشناخته")
	ErrAuthorization           = customeerror.New("1003", "unauthorized user", "خطای دسترسی کاربر")
	ErrUserNotFound            = customeerror.New("1004", "user not found", "کاربری با این مشخصات در سیستم موجود تیست")
	ErrPasswordIsWrong         = customeerror.New("1005", "password is wrong", "رمز عبور اشتباه است")
	ErrAdminAccessibilityEmpty = customeerror.New("1006", "admin accessibility cant be empty",
		"دسترسی ادمین نمیتواند خالی باشد")
	ErrUserIsTaken = customeerror.New("1007", "user name is already taken",
		"کاربر با این نام قبلا در سیستم ذخیره شده است")
	ErrParams = customeerror.New("1029", "required parameters for request are empty",
		"پارامترهای لازم در درخواست ارسال نشده است",
	)
	ErrMotorIsNotExist = customeerror.New("1008", "motor with this id and name is not exist in shop",
		"موتور با این مشخصات در سیستم موجود نیست")
	ErrMobile = customeerror.New("1009", "mobile number is not valid",
		"شماره موبایل صحیح نمیباشد.")
	ErrShopNotExist = customeerror.New("1010", "shop with this id is not exist",
		"فروشگاهی با این شماره شناسایی موجود نیست")
)
