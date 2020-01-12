package users

import "errors"

import "golang.org/x/crypto/bcrypt"

import "github.com/jefridev/securityapi/common"

// UserModel this should only be used for database concerns.
type UserModel struct {
	ID           uint64  `gorm:"auto_increment;primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email"`
	Bio          string  `gorm:"column:bio; size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password; not null"`
}

// setPassword hashes password information for a specific user.
func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// checkPassword compares password and tells if password provided is going to be equal.
func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// Update properties for UserModel.
func (u *UserModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(u).Update(data).Error
	return err
}

// FindOneUser input the condition and It will return UserModel.
func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveOne UserModel will be saved in database returning error if It couldnt save or
// nil if everything was done.
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
