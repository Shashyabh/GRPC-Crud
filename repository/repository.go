package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shashyabh/mygo/models"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	Close()
	CreateUser(user *models.User) error
	GetUserById(userId string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

func NewPostgresRepository () (UserRepositoryInterface, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/grpcuser?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
 
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &UserRepository{db}, nil
}

func (repo *UserRepository) Close() {
	repo.db.Close()
}
func (r *UserRepository) CreateUser (user *models.User) error {
	query :="Insert into users (id,name, salary,department,address_id) values (?,?,?,?,?)"

	_, err :=r.db.Exec(query,user.ID,user.Name,user.Salary,user.Department,user.AddressId);

	if err != nil {
		log.Println("Error initializing account client:", err)
	}
	return nil
}

func (r *UserRepository) GetUserById ( userId string) (*models.User,error){
	query :="select id,name, salary,department,address_id FROM users where id = ?"

	rows := r.db.QueryRow(query,userId)

	var user models.User;

	if err := rows.Scan(&user.ID, &user.Name,&user.Salary,&user.Department,&user.AddressId); err != nil{
		if err == sql.ErrNoRows {
            return nil, nil
        }
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers ()([] models.User,error){
	query := "SELECT id, name, salary, department, address_id FROM users"

	rows, err :=r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name,&user.Salary,&user.Department,&user.AddressId); err !=nil{
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users,nil
}
