package db

import (
	"database/sql"
	"log"
)

type OrdersORM interface {
	Add()
	Get()
	List()
}

type Orders struct {
	db *sql.DB
}

func (x Orders) Add(who int, what string) {
    q := `
        insert into orders (user_id, coffee_id)
            select $1, c.id
            from coffee as c
            where c.name = $2;
    `
    _, err := x.db.Exec(q, who, what)
    if err != nil {
		log.Fatal(err)
	}
}

func (x Orders) List(userId int) {
	q := `
        select o.id, o.user_id, c.name 
        from orders as o
          inner join coffee as c
          on o.coffee_id = c.id;
    `
	rows, err := x.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
    log.Printf("Orders:")
	for rows.Next() {
		var (
			id        int
			coffee string
			user_id  int
		)
		if err := rows.Scan(&id, &user_id, &coffee); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %d of %s\n", id, user_id, coffee)
	}
}
