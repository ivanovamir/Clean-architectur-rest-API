package repository

import (
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

const (
	cost        = 10
	userCtx     = "userId"
	passwordNew = 9999999999999999
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(userDTO *dto.RegisterUser) (int, error) {
	//var user, newUser *models.User
	//var code *models.Code
	//
	//switch userDTO.Inn {
	//case "":
	//	switch r.db.Where("phone = ?", userDTO.Phone).Last(&code).Error {
	//	case nil:
	//		switch r.db.Where("phone = ?", userDTO.Phone).First(&user).Error {
	//		case nil:
	//
	//			if userDTO.Code != code.Code {
	//				return 0, fmt.Errorf("wrong code")
	//			} else {
	//				r.db.Where("phone LIKE ?", userDTO.Phone).Unscoped().Delete(&models.Code{})
	//				return int(user.Id), nil
	//			}
	//		default:
	//			newUser = &models.User{
	//				CreatedAt: time.Now(),
	//				Phone:     userDTO.Phone,
	//			}
	//		}
	//	default:
	//		return 0, fmt.Errorf("unauthorized")
	//	}
	//default:
	//	switch r.db.Where("email = ?", userDTO.Email).Last(&code).Error {
	//	case nil:
	//		//TODO inn checker doesnt work good
	//		switch r.db.Where("email = ?", userDTO.Email).Or("inn = ?", userDTO.Inn).First(&user).Error {
	//		case nil:
	//			return 0, fmt.Errorf("user already registered")
	//		default:
	//			if userDTO.Code != code.Code {
	//				return 0, fmt.Errorf("wrong code")
	//			} else {
	//				passwordHash, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), cost)
	//
	//				if err != nil {
	//					panic(err)
	//				}
	//
	//				newUser = &models.User{
	//					CreatedAt:    time.Now(),
	//					Inn:          userDTO.Inn,
	//					Email:        userDTO.Email,
	//					PasswordHash: string(passwordHash),
	//				}
	//				r.db.Where("email LIKE ?", userDTO.Email).Unscoped().Delete(&models.Code{})
	//			}
	//		}
	//	default:
	//		return 0, fmt.Errorf("unregistered")
	//	}
	//}
	//r.db.Create(&newUser)
	//
	//createdCart := &models.Cart{UserID: newUser.Id}
	//r.db.Create(&createdCart)
	//
	//createdFavouriteList := models.Favourites{UserID: newUser.Id}
	//r.db.Create(&createdFavouriteList)
	//
	//return int(newUser.Id), nil
	return 0, nil
}

//func (r *UserRepository) AuthenticateUser(userDTO *dto.UserAuth) (int, error) {
//	var user *models.User
//	var code *models.Code
//
//	switch r.db.Where("phone = ?", userDTO.Phone).Last(&code).Error {
//	case nil:
//		if userDTO.Code != code.Code {
//			return 0, fmt.Errorf("wrong code")
//		} else {
//			r.db.Where("phone LIKE ?", userDTO.Phone).Unscoped().Delete(&models.Code{})
//			switch r.db.Where("phone = ?", userDTO.Phone).First(&user).Error {
//			case nil:
//				return int(user.Id), nil
//			default:
//				return 0, fmt.Errorf("unauthorized")
//			}
//		}
//	default:
//		return 0, fmt.Errorf("unauthorized")
//	}
//}
