package utils

import 
("fmt" 
"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string, error){
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err !=nil{
		return"", fmt.Errorf("hashing failed %w",err)
	}
	return string(hashPassword), nil
}

func ValidatePassword(hashedPassword string, loginPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(loginPassword))
}