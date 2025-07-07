package integration

import (
	"askon/support-api/storage"
	"database/sql"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	composePath := filepath.Join("..", "..", "docker-compose.test.yml")
	cmd := exec.Command("docker-compose", "-f", composePath, "up", "-d", "--wait")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to start container: %v\nOutput: %s", err, string(output))
	}

	var db *sql.DB
	var dbErr error
	for i := 0; i < 20; i++ {
		db, dbErr = sql.Open("mysql", "test_user:test_password@tcp(localhost:3307)/test_db?parseTime=true")
		if dbErr == nil {
			if err = db.Ping(); err == nil {
				break
			}
			db.Close()
		}
		time.Sleep(3 * time.Second)
	}

	if dbErr != nil {
		t.Fatal("Failed to connect after 20 attempts:", err)
	}

	return db
}

func TestArticleStorage_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer func() {
		db.Close()
		composePath := filepath.Join("..", "..", "docker-compose.test.yml")
		exec.Command("docker-compose", "-f", composePath, "down").Run()
	}()
	storage := storage.NewArticleStorage(db)

	t.Run("Empty table", func(t *testing.T) {
		_, err := db.Exec("TRUNCATE TABLE article")
		assert.NoError(t, err)

		result, err := storage.GetArticleTime(
			time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
		)
		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("With test data", func(t *testing.T) {

		_, err := db.Exec("TRUNCATE TABLE article")
		assert.NoError(t, err)

		err = execSQLFile(db, "testdata.sql")
		assert.NoError(t, err)

		result, err := storage.GetArticleTime(
			time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
		)
		assert.NoError(t, err)
		assert.Equal(t, map[string]int{
			"2025-01-01T00:00:00Z": 2,
		}, result)
	})
}

func execSQLFile(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	return err
}
