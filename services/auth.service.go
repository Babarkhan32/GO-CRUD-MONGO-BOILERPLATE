package services
import "exmaple.com/job-x-apis/models"

type AuthService interface {
	Login(*models.Credentials) (*models.User,error) 

}