package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"korzadivpn/pkg/utils"
)

func main() {
	// 1. Validaciones iniciales - Acepta 64 hex chars (32 bytes)
	keyHex := os.Getenv("ENCRYPTION_KEY")
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil || len(keyBytes) != 32 {
		log.Fatal("ENCRYPTION_KEY debe ser una cadena hexadecimal de 64 caracteres (32 bytes)")
	}

	// 2. IMPORTANTÍSIMO: El paquete utils usa os.Getenv("ENCRYPTION_KEY").
	// Para que funcione, debemos poner la clave como un string de 32 bytes literales en el entorno.
	os.Setenv("ENCRYPTION_KEY", string(keyBytes))

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "korzadivpn.db"
	}
// ... resto sin cambios

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Backup
	if err := backupDB(dbPath); err != nil {
		log.Fatal(err)
	}

	// 3. Ejecutar migraciones en una transacción
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	tables := []struct {
		name   string
		column string
	}{
		{"servers", "server_private_key"},
		{"vpn_profiles", "private_key"},
		{"vpn_core_clients", "private_key"},
	}

	for _, t := range tables {
		if err := migrateTable(tx, t.name, t.column); err != nil {
			tx.Rollback()
			log.Fatalf("Error migrando tabla %s, rollback realizado: %v", t.name, err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	log.Println("Migración completada exitosamente.")
}

func backupDB(path string) error {
	source, err := os.Open(path)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(path + ".bak")
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func migrateTable(tx *sql.Tx, tableName, colName string) error {
	// Verificar si la tabla existe
	var exists int
	err := tx.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&exists)
	if err != nil {
		return err
	}
	if exists == 0 {
		log.Printf("Tabla %s no encontrada, saltando.\n", tableName)
		return nil
	}

	rows, err := tx.Query(fmt.Sprintf("SELECT id, %s FROM %s", colName, tableName))
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var value sql.NullString
		if err := rows.Scan(&id, &value); err != nil {
			return err
		}

		if !value.Valid || value.String == "" {
			continue
		}

		// Intentar descifrar para ver si ya está cifrado
		_, err := utils.Decrypt(value.String)
		if err == nil {
			log.Printf("Tabla %s, ID %d: Ya cifrado, saltando.\n", tableName, id)
			continue
		}

		// Si no se puede descifrar, asumimos texto plano y ciframos
		encrypted, err := utils.Encrypt(value.String)
		if err != nil {
			return fmt.Errorf("error cifrando ID %d: %v", id, err)
		}

		_, err = tx.Exec(fmt.Sprintf("UPDATE %s SET %s = ? WHERE id = ?", tableName, colName), encrypted, id)
		if err != nil {
			return err
		}
		log.Printf("Tabla %s, ID %d: Migrado a cifrado.\n", tableName, id)
	}
	return nil
}
