package userstore

import (
	"database/sql"
	_ "github.com/lib/pq"
   "github.com/saks07/go-api/utils"
)

type User struct {
   ID int `json:"id"`
   Username string `json:"username"`
   Email string `json:"email"`
}

type UserStore interface {
   SaveUser(username string, email string) error
   GetAllUsers() ([]User, error)
}

type SQLUserStore struct {
   DB *sql.DB
}

// Global variables
var dbTable string = "users"

func (s *SQLUserStore) SaveUser(username string, email string) error {
   var query string = utils.QueryStringTable("INSERT INTO {table} (username, email) VALUES ($1, $2)", dbTable)
   stmt, stmtErr := s.DB.Prepare(query)
   
   if stmtErr != nil {
		return stmtErr
	}

   defer stmt.Close()

   _, err := stmt.Exec(username, email)

   return err
}

func (s *SQLUserStore) GetAllUsers() ([]User, error) {
   var query string = utils.QueryStringTable("SELECT id, username, email FROM {table}", dbTable)
   rows, err := s.DB.Query(query)
   if err != nil {
     return nil, err
   }

   defer rows.Close()

   var users []User

   for rows.Next() {
     var user User

     if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
       return nil, err
     }

     users = append(users, user)
   }

   return users, nil
}