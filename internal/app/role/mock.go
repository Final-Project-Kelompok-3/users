package role

import (
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Final-Project-Kelompok-3/users/internal/model"
	"github.com/Final-Project-Kelompok-3/users/pkg/util/mock"
	"gorm.io/gorm"
)
func NewMock(expectQuery string) (*gorm.DB) {
	db, mock := mock.DBConnection()
	
	var roles = []model.Role{
		{
			Model: model.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Admin",
			Description: "Admin sistem",
		},
		{
			Model: model.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Sekolah",
			Description: "Sekolah",
		},
		{
			Model: model.Model{
				ID:        3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Siswa",
			Description: "Siswa",
		},
	}
	
	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at", "deleted_at"})
	
	for _, v := range roles {
		rows = rows.AddRow(v.ID, v.Name, v.Description, v.CreatedAt, v.UpdatedAt, v.DeletedAt)
	}
	
	mock.ExpectQuery(regexp.QuoteMeta(expectQuery)).WillReturnRows(rows)

	return db
}