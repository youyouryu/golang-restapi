package main

import "log"

// User represents a user.
type User struct {
	UserID   string `json:"user_id,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

func (u User) create() error {
	stmtIns, err := db.Prepare("insert into users (user_id, password, nickname, comment) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()
	if _, err := stmtIns.Exec(u.UserID, passwordHash(u.Password), u.Nickname, u.Comment); err != nil {
		return err
	}
	return nil
}

func (u User) update() error {
	stmtUpd, err := db.Prepare("update users set nickname = ?, comment = ? where user_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmtUpd.Exec(u.Nickname, u.Comment, u.UserID); err != nil {
		return err
	}
	return nil
}

func (u User) delete() error {
	stmtDel, err := db.Prepare("delete from users where user_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmtDel.Exec(u.UserID); err != nil {
		return err
	}
	return nil
}

func userByID(userID string) (user User, err error) {
	err = db.QueryRow("select user_id, password, nickname, comment from users where user_id = ?", userID).Scan(&user.UserID, &user.Password, &user.Nickname, &user.Comment)
	return
}
