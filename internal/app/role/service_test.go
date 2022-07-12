package role

import (
	"context"
	"testing"

	"github.com/Final-Project-Kelompok-3/users/internal/dto"
	"github.com/Final-Project-Kelompok-3/users/internal/factory"
)

func TestFindByID(t *testing.T) {
	db := NewMock(`SELECT * FROM "roles" WHERE id = $1 AND "roles"."deleted_at" IS NULL ORDER BY "roles"."id" LIMIT 1`)
	
	f := factory.NewFactory(db)
	service := NewService(f)

	payload := new(dto.ByIDRequest)
	payload.ID = 1

	_, err := service.FindByID(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	// if len(data) != 1 {
	// 	t.Error("Expected 1, got ", len(data))
	// }
}