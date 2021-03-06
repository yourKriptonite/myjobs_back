package repository

import (
	. "2019_2_IBAT/pkg/pkg/models"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestDBUserStorage_GetVacancies_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}

	rows := sqlmock.
		NewRows([]string{"id", "own_id", "company_name", "experience",
			"position", "tasks", "requirements", "wage_from", "wage_to", "conditions", "about",
		})
	expect := []Vacancy{
		{
			ID:           uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
			OwnerID:      uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
			CompanyName:  "MC",
			Experience:   "7 years",
			Position:     "cleaner",
			Tasks:        "cleaning rooms",
			Requirements: "work for 24 hours per week",
			WageFrom:     "100 500 руб",
			WageTo:       "101 500.00 руб",
			Conditions:   "Nice geolocation",
			About:        "Hello employer",
		},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID.String(), item.OwnerID.String(), item.CompanyName, item.Experience,
			item.Position, item.Tasks, item.Requirements,
			item.WageFrom, item.WageTo, item.Conditions, item.About,
		)
	}

	mock.
		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
			"v.region, v.type_of_employment, v.work_schedule " +
			" FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id;").
		WithArgs().
		WillReturnRows(rows)

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	dummyMap := make(map[string]interface{})
	AuthRec := AuthStorageValue{
		ID: uuid.New(),
	}
	vacancies, err := repo.GetVacancies(AuthRec, dummyMap)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(vacancies, expect) {
		t.Errorf("results not match,\n want\n%v,\n have\n %v\n", expect, vacancies)
		return
	}
}

func TestDBUserStorage_GetVacancies_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}

	mock.
		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
			"v.region, v.type_of_employment, v.work_schedule " +
			" FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id;").
		WithArgs().
		WillReturnError(errors.New("GetVacancies: error while querying"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	dummyMap := make(map[string]interface{})
	authRec := AuthStorageValue{
		ID: uuid.New(),
	}
	vacancies, err := repo.GetVacancies(authRec, dummyMap)
	fmt.Println(vacancies)

	if err == nil {
		t.Errorf("Expected err")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

}

func TestDBUserStorage_GetVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	rows := sqlmock.
		NewRows([]string{"id", "own_id", "company_name", "experience",
			"position", "tasks", "requirements", "wage_from", "wage_to", "conditions", "about",
		})
	expect := []Vacancy{
		{
			ID:           uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
			OwnerID:      uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
			CompanyName:  "MC",
			Experience:   "7 years",
			Position:     "mid",
			Tasks:        "cleaning rooms",
			Requirements: "work for 24 hours per week",
			WageFrom:     "100 500.00 руб",
			WageTo:       "101 500.00 руб",
			Conditions:   "Nice geolocation",
			About:        "Hello employer",
		},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID.String(), item.OwnerID.String(), item.CompanyName, item.Experience,
			item.Position, item.Tasks, item.Requirements,
			item.WageFrom, item.WageTo, item.Conditions, item.About,
		)
	}

	mock.
		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
			"v.region, v.type_of_employment, v.work_schedule FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id WHERE").
		WithArgs().
		WillReturnRows(rows)

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a")
	userID := uuid.New()
	item, err := repo.GetVacancy(id, userID)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match,\n want\n%v,\n have\n %v\n", expect[0], item)
		return
	}
}

func TestDBUserStorage_GetVacancy_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642bbba")
	mock.
		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
			"v.region, v.type_of_employment, v.work_schedule FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id WHERE").
		WithArgs(id).
		WillReturnError(errors.New("GetVacancy: error while querying"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	userID := uuid.New()
	vacancy, err := repo.GetVacancy(id, userID)
	fmt.Println(vacancy)

	if err == nil {
		t.Errorf("Expected err")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDBUserStorage_CreateVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancy := Vacancy{
		ID:               uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
		OwnerID:          uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
		CompanyName:      "MC",
		Experience:       "7 years",
		Position:         "mid",
		Tasks:            "cleaning rooms",
		Requirements:     "work for 24 hours per week",
		WageFrom:         "100 500.00 руб",
		WageTo:           "120 500.00 руб",
		Conditions:       "Nice geolocation",
		About:            "Hello employer",
		TypeOfEmployment: "someType",
		WorkSchedule:     "WorkSchedule",
		Region:           "Moscow",
	}

	mock.
		ExpectExec(`INSERT INTO vacancies`).
		WithArgs(
			vacancy.ID, vacancy.OwnerID, vacancy.Experience, vacancy.Position, vacancy.Tasks,
			vacancy.Requirements, vacancy.Conditions, vacancy.WageFrom, vacancy.WageTo, vacancy.About,
			vacancy.Region, vacancy.TypeOfEmployment, vacancy.WorkSchedule,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.CreateVacancy(vacancy)

	if !ok {
		t.Error("Failed to create vacancy\n")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDBUserStorage_CreateVacancy_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancy := Vacancy{
		ID:               uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
		OwnerID:          uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
		CompanyName:      "MC",
		Experience:       "7 years",
		Position:         "mid",
		Tasks:            "cleaning rooms",
		Requirements:     "work for 24 hours per week",
		WageFrom:         "100 500.00 руб",
		WageTo:           "120 500.00 руб",
		Conditions:       "Nice geolocation",
		About:            "Hello employer",
		TypeOfEmployment: "someType",
		WorkSchedule:     "WorkSchedule",
		Region:           "Moscow",
	}

	mock.
		ExpectExec(`INSERT INTO vacancies`).
		WithArgs(
			vacancy.ID, vacancy.OwnerID, vacancy.Experience, vacancy.Position, vacancy.Tasks,
			vacancy.Requirements, vacancy.Conditions, vacancy.WageFrom, vacancy.WageTo, vacancy.About,
			vacancy.Region, vacancy.TypeOfEmployment, vacancy.WorkSchedule,
		).
		WillReturnError(fmt.Errorf("bad query"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.CreateVacancy(vacancy)

	if ok {
		t.Errorf("expected false, got true")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDBUserStorage_DeleteVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a")

	mock.
		ExpectExec(`DELETE FROM vacancies`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	err = repo.DeleteVacancy(id)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	mock.
		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
			"v.region, v.type_of_employment, v.work_schedule FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id WHERE").
		WithArgs(id).
		WillReturnError(errors.New("GetVacancy: error while querying"))

	userID := uuid.New()
	_, err = repo.GetVacancy(id, userID)

	if err == nil {
		t.Errorf("Expected err")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDBUserStorage_DeleteVacancy_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a")

	mock.
		ExpectExec(`DELETE FROM vacancies`).
		WithArgs(id).
		WillReturnError(errors.Errorf("error"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	err = repo.DeleteVacancy(id)

	if err == nil {
		t.Errorf("Expected err")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDBUserStorage_PutVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancy := Vacancy{
		ID:               uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
		OwnerID:          uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
		CompanyName:      "MC",
		Experience:       "7 years",
		Position:         "mid",
		Tasks:            "cleaning rooms",
		Requirements:     "work for 24 hours per week",
		WageFrom:         "100 500.00 руб",
		WageTo:           "101 500.00 руб",
		Conditions:       "Nice geolocation",
		About:            "Hello employer",
		TypeOfEmployment: "someType",
		WorkSchedule:     "WorkSchedule",
		Region:           "Moscow",
	}

	mock.
		ExpectExec(`UPDATE vacancies SET`).
		WithArgs(
			vacancy.Experience, vacancy.Position, vacancy.Tasks, vacancy.Requirements,
			vacancy.WageFrom, vacancy.WageTo, vacancy.Conditions, vacancy.About,
			vacancy.Region, vacancy.TypeOfEmployment, vacancy.WorkSchedule,
			vacancy.ID, vacancy.OwnerID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.PutVacancy(vacancy, vacancy.OwnerID, vacancy.ID)

	if !ok {
		t.Error("Failed to put vacancy\n")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDBUserStorage_PutVacancy_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancy := Vacancy{
		ID:               uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"), //invalid id
		OwnerID:          uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
		CompanyName:      "MC",
		Experience:       "7 years",
		Position:         "cleaner",
		Tasks:            "cleaning rooms",
		Requirements:     "work for 24 hours per week",
		WageFrom:         "100 500.00 руб",
		WageTo:           "101 500.00 руб",
		Conditions:       "Nice geolocation",
		About:            "Hello employer",
		TypeOfEmployment: "someType",
		WorkSchedule:     "WorkSchedule",
		Region:           "Moscow",
	}

	mock.
		ExpectExec(`UPDATE vacancies SET`).
		WithArgs(
			vacancy.Experience, vacancy.Position, vacancy.Tasks, vacancy.Requirements,
			vacancy.WageFrom, vacancy.WageTo, vacancy.Conditions, vacancy.About,
			vacancy.Region, vacancy.TypeOfEmployment, vacancy.WorkSchedule,
			vacancy.ID, vacancy.OwnerID,
		).
		WillReturnError(fmt.Errorf("bad query"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.PutVacancy(vacancy, vacancy.OwnerID, vacancy.ID)

	if ok {
		t.Errorf("expected false, got true")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
