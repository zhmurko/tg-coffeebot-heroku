package db

import (
	// "database/sql"
	"log"
)
// Order keep Orders in db
type Order struct {
	UserID int
	Coffee  string
}

// Add an Order in db
func Add(who int, what string) {
	q := `
        insert into orders (user_id, coffee_id)
            select $1, c.id
            from coffee as c
            where c.name = $2;
    `
	_, err := DB.Exec(q, who, what)
	if err != nil {
		log.Fatal(err)
	}
}

// List orders by User
func List(userID int) []Order {
	q := `
        select o.id, o.user_id, c.name 
        from orders as o
          inner join coffee as c
          on o.coffee_id = c.id
		  where o.user_id = $1;
    `
	rows, err := DB.Query(q, userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var s []Order
	for rows.Next() {
		var (
			id int
			o  Order
		)
		if err := rows.Scan(&id, &o.UserID, &o.Coffee); err != nil {
			log.Fatal(err)
		}
		s = append(s, o)
	}
	return s
}
