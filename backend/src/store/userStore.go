package store

import (
	"github.com/akzhigitov/highload-social-network/src/models"
)

func CreateUser(user *models.User) (err error) {
	_, err = db.Exec("INSERT INTO user(user_id, email, password,  last_name, first_name, city, sex, birthday, interests) VALUES (?,?,?,?,?,?,?,?,?)",
		user.Id, user.Email, user.Password, user.LastName, user.FirstName, user.City, user.Sex, user.Birthday, user.Interests)
	return err
}

func UpdateUser(userID int, user *models.User) (err error) {
	_, err = db.Exec("UPDATE user SET last_name=?, first_name=?, city=?, sex=?, interests=? WHERE user_id = ?",
		user.LastName, user.FirstName, user.City, user.Sex, user.Interests, userID)
	return err
}

func GetPasswordByEmail(email string) (password string, err error) {
	err = db.QueryRow("SELECT password FROM user WHERE email=?", email).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

func GetUserInfoByEmail(email string) (*models.User, error) {
	var userId int
	var password []byte

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.
		QueryRow("SELECT user_id, password FROM user WHERE email=?", email).
		Scan(&userId, &password)
	if err != nil {
		return nil, err
	}

	u := models.User{
		Password: password,
	}
	u.Id = userId
	tx.Commit()
	return &u, nil
}

func GetUser(userID int) (user *models.User, err error) {

	var userId int
	var lastName string
	var firstName string
	var city string
	var sex byte
	var birthday string
	var interests string

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.
		QueryRow("SELECT user_id, last_name, first_name, city, sex, birthday, interests FROM user WHERE user_id=?", userID).
		Scan(&userId, &lastName, &firstName, &city, &sex, &birthday, &interests)
	if err != nil {
		return nil, err
	}
	var followers []models.Friend
	var following []models.Friend
	friendsRows, err := tx.Query("SELECT user_id1, u1.first_name, u1.last_name,user_id2, u2.first_name, u2.last_name FROM friends join user u1 on user_id1 = u1.user_id join user u2 on user_id2 = u2.user_id WHERE user_id1 =? OR user_id2 = ?", userId, userId)

	if err != nil {
		return nil, err
	}
	defer friendsRows.Close()
	var userID1 int
	var userID2 int
	var userFirstName1 string
	var userFirstName2 string
	var userLastName1 string
	var userLastName2 string
	for friendsRows.Next() {
		err = friendsRows.Scan(&userID1, &userFirstName1, &userLastName1, &userID2, &userFirstName2, &userLastName2)
		if err != nil {
			return nil, err
		}
		if userID1 != userId {
			followers = append(followers, models.Friend{
				Model: models.Model{
					Id: userID1,
				},
				FirstName: userFirstName1,
				LastName:  userLastName1,
			})
		} else {
			following = append(following, models.Friend{
				Model: models.Model{
					Id: userID2,
				},
				FirstName: userFirstName2,
				LastName:  userLastName2,
			})
		}
	}

	u := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Sex:       sex,
		Interests: interests,
		Followers: followers,
		Following: following,
		City:      city,
	}
	u.Id = userId
	tx.Commit()
	return &u, nil
}

func Follow(user1, user2 int) (err error) {
	_, err = db.Exec("INSERT friends(user_id1, user_id2) VALUES (?,?)", user1, user2)
	return err
}

func UnFollow(user1, user2 int) (err error) {
	_, err = db.Exec("DELETE FROM friends WHERE user_id1=? AND user_id2=?", user1, user2)
	return err
}

func GetUsers() (users []*models.User, err error) {

	rows, err := db.Query("SELECT user_id, last_name, first_name, city, sex, birthday, email FROM user")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userId int
	var lastName string
	var firstName string
	var city string
	var sex byte
	var birthday string
	var email string

	for rows.Next() {
		err = rows.Scan(&userId, &lastName, &firstName, &city, &sex, &birthday, &email)
		if err != nil {
			return nil, err
		}
		u := models.User{
			Model: models.Model{
				Id: userId,
			},
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday,
			Sex:       sex,
			City:      city,
			Email:     email,
		}
		users = append(users, &u)
	}

	return users, nil
}
