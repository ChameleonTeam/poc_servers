package main

import (
    "database/sql"
    "fmt"
    
    _ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./USERS.db")
    if err != nil {
        return nil, err
    }
    if db == nil {
        return nil, fmt.Errorf("DB nil")
    }
    return db, err
}

func PopulateDb(db *sql.DB) error {
    // create table if not exists
    sql_table := `
	CREATE TABLE IF NOT EXISTS USERS(
		DNI     TEXT PRIMARY KEY,
		NAMEP   TEXT,
		SURNAME TEXT,
		GENDER  TEXT,
		ADDR	TEXT,
		PHONE   TEXT,
		WEIGHT  INTEGER
	);
	`
    _, err := db.Exec(sql_table)
    if err != nil {
        return err
    }
    
    statement, err := db.Prepare("INSERT INTO USERS (DNI, NAMEP, SURNAME, GENDER, ADDR, PHONE, WEIGHT) VALUES (?, ?, ?, ?, ?, ?, ?)")
    statement.Exec("85910025F", "Felipe", "Murillo", "Male", "ADDR1", "666666666", 88)
    statement.Exec("52145265L", "Virgilio", "Garcia", "Male", "ADDR1", "666666666", 100)
    statement.Exec("32514521A", "Dani", "Escribano", "Male", "ADDR1", "666666666", 81)
    statement.Exec("69852235X", "Andres", "Ruiz", "Male", "ADDR1", "666666666", 70)
    statement.Exec("65256985A", "Carlos", "Montero", "Male", "ADDR1", "666666666", 80)
    statement.Exec("32601125T", "Alejandro", "Galindo", "Male", "ADDR1", "666666666", 85)
    
    return nil
}

func insertUser(db *sql.DB, pr *Person) (*Person, error) {
    
    statement, err := db.Prepare("INSERT INTO USERS (DNI, NAMEP, SURNAME, GENDER, ADDR, PHONE, WEIGHT) VALUES (?, ?, ?, ?, ?, ?, ?)")
    result, err := statement.Exec(pr.Dni, pr.Name, pr.Surname, pr.Gender, pr.Addr, pr.Phone, pr.Weight)
    
    fmt.Println(result)
    
    if err != nil {
        return nil, err
    }
    
    return pr, nil
}

func updateUser(db *sql.DB, pr *Person, dni string) (*Person, error) {

    statement, err := db.Prepare("UPDATE USERS SET DNI=?, NAMEP=?, SURNAME=?, GENDER=?, ADDR=?, PHONE=?, WEIGHT=? WHERE DNI=?")
    result, err := statement.Exec(pr.Dni, pr.Name, pr.Surname, pr.Gender, pr.Addr, pr.Phone, pr.Weight, dni)

    fmt.Println(result)

    if err != nil {
        return nil, err
    }

    return pr, nil

}

func getUser(db *sql.DB, dni string) (*Person, error) {
    
    var person Person
    
    rows, err := db.Query(`SELECT DNI, NAMEP, SURNAME, GENDER, ADDR, PHONE, WEIGHT FROM USERS where DNI = ?`, dni)
    
    if err != nil {
        return nil, err
    }
    
    for rows.Next() {
        err := rows.Scan(&person.Dni,
            &person.Name,
            &person.Surname,
            &person.Gender,
            &person.Addr,
            &person.Phone,
            &person.Weight,
        )
        if err != nil {
            return nil, err
        }
    }
    defer rows.Close()
    return &person, nil
}

func listUsers(db *sql.DB) ([]Person, error) {
    
    personList := make([]Person, 0)
    
    rows, err := db.Query(`SELECT * FROM USERS`)
    
    if err != nil {
        return nil, err
    }
    
    var person Person
    
    for rows.Next() {
        _ = rows.Scan(&person.Dni,
            &person.Name,
            &person.Surname,
            &person.Gender,
            &person.Addr,
            &person.Phone,
            &person.Weight,
        )
        
        personList = append(personList, person)
    }
    defer rows.Close()
    
    return personList, nil
}

func deleteUser(db *sql.DB, dni string) (*Person, error) {

    var person *Person

    person, err := getUser(db, dni)

    statement, err := db.Prepare(`DELETE FROM USERS where DNI = ?`)
    _, err = statement.Exec(dni)
    if err != nil {
        return nil, err
    }

    return person, nil
}

func DropTable(db *sql.DB) error {
    statement, _ := db.Prepare("DROP TABLE IF EXISTS USERS")
    
    _, err := statement.Exec()
    
    if err != nil {
        return err
    }
    return nil
}
