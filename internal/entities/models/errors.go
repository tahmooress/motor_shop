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
	ErrShopAlreadyExist = customeerror.New("1011", "shop with this name is already exist",
		"فروشگاهی با این نام و مشخصات قبلا در سیستم ثبت شده است.")
	ErrEquityID = customeerror.New("1012", "equity with this id is not exits",
		"حسابی با این شماره شناسیایی در سیستم موجود تیسن")
	ErrFactorNotExist = customeerror.New("1013", "factor with this number is not exist",
		"فاکتوری با این شماره در سیستم موجود نیست")
	ErrTxTypeAndSubject = customeerror.New("1014", "type or subject of transaction is not valid",
		"فیلد نوع یا عنوان برای تراکنش ورودی معتبر نیست")
	ErrIDIsNotValid = customeerror.New("1015", "entry id number is not valid",
		"شماره شناسایی ورودی یک شماره شناسیایی معتبر نمیباشد")
)
