package repository

import (
	. "2019_2_IBAT/pkg/pkg/models"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// func TestDBUserStorage_GetFavoriteVacancies_Correct(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	defer db.Close()
// 	sqlxDB := sqlx.NewDb(db, "sqlmock")

// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}

// 	rows := sqlmock.
// 		NewRows([]string{"id", "own_id", "company_name", "experience",
// 			"position", "tasks", "requirements", "wage_from", "wage_to", "conditions", "about",
// 		})
// 	expect := []Vacancy{
// 		{
// 			ID:           uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
// 			OwnerID:      uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
// 			CompanyName:  "MC",
// 			Experience:   "7 years",
// 			Position:     "cleaner",
// 			Tasks:        "cleaning rooms",
// 			Requirements: "work for 24 hours per week",
// 			WageFrom:     "100 500 руб",
// 			WageTo:       "101 500.00 руб",
// 			Conditions:   "Nice geolocation",
// 			About:        "Hello employer",
// 		},
// 	}

// 	for _, item := range expect {
// 		rows = rows.AddRow(item.ID.String(), item.OwnerID.String(), item.CompanyName, item.Experience,
// 			item.Position, item.Tasks, item.Requirements, item.WageFrom, item.WageTo,
// 			item.Conditions, item.About,
// 		)
// 	}
// 	AuthRec := AuthStorageValue{
// 		ID: uuid.New(),
// 	}

// 	mock.
// 		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience, " +
// 			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
// 			"v.region, v.type_of_employment, v.work_schedule " +

// 			"FROM favorite_vacancies AS fv " +
// 			"JOIN vacancies AS v ON (fv.vacancy_id = v.id) " +
// 			"JOIN companies AS c ON v.own_id = c.own_id WHERE"). //fux query
// 		WithArgs(AuthRec.ID).
// 		WillReturnRows(rows)

// 	repo := DBUserStorage{
// 		DbConn: sqlxDB,
// 	}

// 	vacancies, err := repo.GetFavoriteVacancies(AuthRec)

// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}

// 	require.Equal(t, vacancies, expect, "The two values should be the same.")

// }

// func TestDBUserStorage_GetFavoriteVacancies_Fail(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	defer db.Close()
// 	sqlxDB := sqlx.NewDb(db, "sqlmock")

// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}

// 	AuthRec := AuthStorageValue{
// 		ID: uuid.New(),
// 	}
// 	mock.
// 		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience, " +
// 			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
// 			"v.region, v.type_of_employment, v.work_schedule " +
// 			"FROM favorite_vacancies AS fv " +
// 			"JOIN vacancies AS v ON (fv.vacancy_id = v.id) " +
// 			"JOIN companies AS c ON v.own_id = c.own_id WHERE").
// 		WithArgs(AuthRec.ID).
// 		WillReturnError(errors.New("GetVacancies: error while querying"))

// 	repo := DBUserStorage{
// 		DbConn: sqlxDB,
// 	}

// 	vacancies, err := repo.GetFavoriteVacancies(AuthRec)
// 	fmt.Println(vacancies)

// 	if err == nil {
// 		t.Errorf("Expected err")
// 		return
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}

// }

// func TestDBUserStorage_GetVacancy_Correct(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	defer db.Close()
// 	sqlxDB := sqlx.NewDb(db, "sqlmock")

// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer sqlxDB.Close()

// 	rows := sqlmock.
// 		NewRows([]string{"id", "own_id", "company_name", "experience",
// 			"position", "tasks", "requirements", "wage_from", "wage_to", "conditions", "about",
// 		})
// 	expect := []Vacancy{
// 		{
// 			ID:           uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a"),
// 			OwnerID:      uuid.MustParse("92b77a73-bac7-4597-ab71-7b5fbe53052d"),
// 			CompanyName:  "MC",
// 			Experience:   "7 years",
// 			Position:     "mid",
// 			Tasks:        "cleaning rooms",
// 			Requirements: "work for 24 hours per week",
// 			WageFrom:     "100 500.00 руб",
// 			WageTo:       "101 500.00 руб",
// 			Conditions:   "Nice geolocation",
// 			About:        "Hello employer",
// 		},
// 	}

// 	for _, item := range expect {
// 		rows = rows.AddRow(item.ID.String(), item.OwnerID.String(), item.CompanyName, item.Experience,
// 			item.Position, item.Tasks, item.Requirements,
// 			item.WageFrom, item.WageTo, item.Conditions, item.About,
// 		)
// 	}

// 	mock.
// 		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
// 			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
// 			"v.region, v.type_of_employment, v.work_schedule FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id WHERE").
// 		WithArgs().
// 		WillReturnRows(rows)

// 	repo := DBUserStorage{
// 		DbConn: sqlxDB,
// 	}

// 	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642ad8a")
// 	userID := uuid.New()
// 	item, err := repo.GetVacancy(id, userID)

// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// 	if !reflect.DeepEqual(item, expect[0]) {
// 		t.Errorf("results not match,\n want\n%v,\n have\n %v\n", expect[0], item)
// 		return
// 	}
// }

// func TestDBUserStorage_GetVacancy_Fail(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	defer db.Close()
// 	sqlxDB := sqlx.NewDb(db, "sqlmock")

// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer sqlxDB.Close()

// 	id := uuid.MustParse("f14c6104-3430-413b-ab4e-e31c8642bbba")
// 	mock.
// 		ExpectQuery("SELECT v.id, v.own_id, c.company_name, v.experience," +
// 			"v.position, v.tasks, v.requirements, v.wage_from, v.wage_to, v.conditions, v.about, " +
// 			"v.region, v.type_of_employment, v.work_schedule FROM vacancies AS v JOIN companies AS c ON v.own_id = c.own_id WHERE").
// 		WithArgs(id).
// 		WillReturnError(errors.New("GetVacancy: error while querying"))

// 	repo := DBUserStorage{
// 		DbConn: sqlxDB,
// 	}

// 	userID := uuid.New()
// 	vacancy, err := repo.GetVacancy(id, userID)
// 	fmt.Println(vacancy)

// 	if err == nil {
// 		t.Errorf("Expected err")
// 		return
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// }

func TestDBUserStorage_CreateFavoriteVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	fv := FavoriteVacancy{
		PersonID:  uuid.New(),
		VacancyID: uuid.New(),
	}
	mock.
		ExpectExec(`INSERT INTO favorite_vacancies`).
		WithArgs(
			fv.PersonID, fv.VacancyID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.CreateFavorite(fv)

	if !ok {
		t.Error("Failed to create favorite vacancy\n")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDBUserStorage_CreateFavoriteVacancy_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	fv := FavoriteVacancy{
		PersonID:  uuid.New(),
		VacancyID: uuid.New(),
	}
	mock.
		ExpectExec(`INSERT INTO favorite_vacancies`).
		WithArgs(
			fv.PersonID, fv.VacancyID,
		).
		WillReturnError(fmt.Errorf("bad query"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	ok := repo.CreateFavorite(fv)

	if ok {
		t.Errorf("expected false, got true")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDBUserStorage_DeleteFavoriteVacancy_Correct(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancyID := uuid.New()
	authInfo := AuthStorageValue{
		ID:   uuid.New(),
		Role: SeekerStr,
	}
	mock.
		ExpectExec(`DELETE FROM favorite_vacancies`).
		WithArgs(vacancyID, authInfo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	err = repo.DeleteFavoriteVacancy(vacancyID, authInfo)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDBUserStorage_DeleteFavoriteVacancy_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer sqlxDB.Close()

	vacancyID := uuid.New()
	authInfo := AuthStorageValue{
		ID:   uuid.New(),
		Role: SeekerStr,
	}
	mock.
		ExpectExec(`DELETE FROM favorite_vacancies`).
		WithArgs(vacancyID, authInfo.ID).
		WillReturnError(errors.Errorf("error"))

	repo := DBUserStorage{
		DbConn: sqlxDB,
	}

	err = repo.DeleteFavoriteVacancy(vacancyID, authInfo)

	if err == nil {
		t.Errorf("Expected err")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
