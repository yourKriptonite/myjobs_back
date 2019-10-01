package interfaces

import (
	"github.com/google/uuid"
)

type SeekerReg struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Password   string `json:"password"`
}

type EmployerReg struct {
	CompanyName      string `json:"company_name"`
	Site             string `json:"site"`
	FirstName        string `json:"first_name"`
	SecondName       string `json:"second_name"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	ExtraPhoneNumber string `json:"extra_phone_number"`
	Password         string `json:"password"`
	City             string `json:"city"`
	EmplNum          string `json:"empl_num"`
}

type Seeker struct {
	Email      string      `json:"email"`
	FirstName  string      `json:"first_name"`
	SecondName string      `json:"second_name"`
	Password   string      `json:"password"`
	PathToImg  string      `json:"path_to_img"`
	Resumes    []uuid.UUID `json:"resumes"` //should be fixed
} //add extra fields

type Employer struct {
	CompanyName      string      `json:"company_name"`
	Site             string      `json:"site"`
	FirstName        string      `json:"first_name"`
	SecondName       string      `json:"second_name"`
	Email            string      `json:"email"`
	PhoneNumber      string      `json:"number"`
	ExtraPhoneNumber string      `json:"extra_number"`
	Password         string      `json:"password"`
	City             string      `json:"city"`
	EmplNum          string      `json:"empl_num"`
	PathToImg        string      `json:"path_to_img"`
	Vacancies        []uuid.UUID `json:"-"`
} //add extra fields

type Resume struct {
	OwnerID uuid.UUID `json:"-"`
	// ID uuid.UUID
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	City        string `json:"city"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birth_date"`
	Sex         string `json:"sex"`
	Citizenship string `json:"citizenship"`
	Experience  string `json:"experience"`
	Profession  string `json:"profession"`
	Position    string `json:"position"`
	Wage        string `json:"wage"`
	Education   string `json:"education"`
	About       string `json:"about"`
}

type Message struct {
	Body string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}

type Id struct {
	Id string `json:"id"`
}

type Class struct {
	Class string `json:"class"`
}

type Vacancy struct {
	OwnerID      uuid.UUID `json:"-"`
	CompanyName  string    `json:"company_name"`
	Experience   string    `json:"experience"`
	Profession   string    `json:"profession"`
	Position     string    `json:"position"`
	Tasks        string    `json:"task"`
	Requirements string    `json:"requirements"`
	Wage         string    `json:"wage"`
	Conditions   string    `json:"conditions"`
	About        string    `json:"about"`
}

type AuthStorageValue struct {
	ID      uuid.UUID
	Expires string
	Class   string
}

type UserAuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// type Handler struct {
// 	Storage AuthStorage
// 	Mu      *sync.Mutex
// }