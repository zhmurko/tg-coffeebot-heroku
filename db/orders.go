package db

import (
	// "database/sql"
	"log"
)
// keep Orders in db
type Order struct {
	User_id int
	Coffee  string
}

// store an Order in db
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

// get Orders by User
func List(userId int) []Order {
	q := `
        select o.id, o.user_id, c.name 
        from orders as o
          inner join coffee as c
          on o.coffee_id = c.id
		  where o.user_id = $1;
    `
	rows, err := DB.Query(q, userId)
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
		if err := rows.Scan(&id, &o.User_id, &o.Coffee); err != nil {
			log.Fatal(err)
		}
		s = append(s, o)
	}
	return s
}
