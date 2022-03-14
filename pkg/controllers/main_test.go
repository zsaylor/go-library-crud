package controllers

import (
	"go-library/pkg/models"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := models.InitDB("root:password@tcp(127.0.0.1:3306)/test-lib?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	db := models.GetDB()
	db.Exec("DELETE FROM books;")
	db.Exec("ALTER TABLE books AUTO_INCREMENT = 1;")

	exitVal := m.Run()
	os.Exit(exitVal)
}
