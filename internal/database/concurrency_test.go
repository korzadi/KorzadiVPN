package database

import (
	"database/sql"
	"fmt"
	"korzadivpn/internal/models"
	"os"
	"sync"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	os.Setenv("ENCRYPTION_KEY", "12345678901234567890123456789012") // 32 bytes
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	CreateVPNClientTable()
	m.Run()
}

func TestConcurrencyAssignIP(t *testing.T) {
	// Setup DB: Using a temporary in-memory DB or resetting the existing one if possible.
	// Since I cannot easily reset the real DB, I will assume a clean state for this test.

	const numRequests = 50
	var wg sync.WaitGroup
	results := make(chan string, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ip, err := GetNextVPNClientIP()
			if err != nil {
				return
			}

			client := models.VPNClient{
				Email:      fmt.Sprintf("user%d@test.com", id),
				ClientName: "Test",
				ClientIP:   ip,
				Status:     "pending",
				// Completar campos requeridos para evitar error de SQL
				ServerID:  1,
				NodeID:    1,
				CreatedAt: "2024-01-01",
				UpdatedAt: "2024-01-01",
			}

			err = CreateVPNClient(client)
			if err == nil {
				results <- ip
			}
		}(i)
	}

	wg.Wait()
	close(results)

	ips := make(map[string]bool)
	count := 0
	for ip := range results {
		if ips[ip] {
			t.Errorf("Duplicate IP detected: %s", ip)
		}
		ips[ip] = true
		count++
	}

	if count == 0 {
		t.Error("No clients were created")
	}
}
