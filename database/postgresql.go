package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// DB instance
var DB *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "123456"
	dbname   = "sakura"
)

// Image struct
type Image struct {
	UUID           uuid.UUID `json:"uuid"`
	WidthAndHeight string    `json:"widthAndHeight"`
	URL            string    `json:"url"`
}

// Images struct
type Images struct {
	ImagesList []Image `json:"images"`
}

// Connect function
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}

// func main() {
// 	// Connect with database
// 	if err := Connect(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create a Fiber app
// 	app := fiber.New()

// 	// Get all records from postgreSQL
// 	app.Get("/employee", func(c *fiber.Ctx) error {
// 		// Select all Employee(s) from database
// 		rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
// 		if err != nil {
// 			return c.Status(500).SendString(err.Error())
// 		}
// 		defer rows.Close()
// 		result := Employees{}

// 		for rows.Next() {
// 			employee := Employee{}
// 			if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
// 				return err // Exit if we get an error
// 			}

// 			// Append Employee to Employees
// 			result.Employees = append(result.Employees, employee)
// 		}
// 		// Return Employees in JSON format
// 		return c.JSON(result)
// 	})

// 	// Add record into postgreSQL
// 	app.Post("/employee", func(c *fiber.Ctx) error {
// 		// New Employee struct
// 		u := new(Employee)

// 		// Parse body into struct
// 		if err := c.BodyParser(u); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		// Insert Employee into database
// 		res, err := db.Query("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
// 		if err != nil {
// 			return err
// 		}

// 		// Print result
// 		log.Println(res)

// 		// Return Employee in JSON format
// 		return c.JSON(u)
// 	})

// 	// Update record into postgreSQL
// 	app.Put("/employee", func(c *fiber.Ctx) error {
// 		// New Employee struct
// 		u := new(Employee)

// 		// Parse body into struct
// 		if err := c.BodyParser(u); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		// Update Employee into database
// 		res, err := db.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
// 		if err != nil {
// 			return err
// 		}

// 		// Print result
// 		log.Println(res)

// 		// Return Employee in JSON format
// 		return c.Status(201).JSON(u)
// 	})

// 	// Delete record from postgreSQL
// 	app.Delete("/employee", func(c *fiber.Ctx) error {
// 		// New Employee struct
// 		u := new(Employee)

// 		// Parse body into struct
// 		if err := c.BodyParser(u); err != nil {
// 			return c.Status(400).SendString(err.Error())
// 		}

// 		// Delete Employee from database
// 		res, err := db.Query("DELETE FROM employees WHERE id = $1", u.ID)
// 		if err != nil {
// 			return err
// 		}

// 		// Print result
// 		log.Println(res)

// 		// Return Employee in JSON format
// 		return c.JSON("Deleted")
// 	})

// }
