package auth

import "golang.org/x/crypto/bcrypt"


//字符串加密
func Encrypt(source string) (string,error) {
	hashedBytes,err :=bcrypt.GenerateFromPassword([]byte(source),bcrypt.DefaultCost)
	return string(hashedBytes),err
}

//比较将加密的文本与纯文本进行比较
func Compare(hashedPassword,password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}
