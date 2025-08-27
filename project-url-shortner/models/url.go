package models

import "project-url-shortner/db"

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}
// GetLongURL retrieves the long URL from the database for a given short URL.
func GetLongURL(shortURL string)(string, error){
	var longURL string
	err := db.DB.QueryRow("SELECT long_url from urls WHERE short_url = $1", shortURL).Scan(&longURL);
	if err != nil {
		return "", err
	}
	return longURL, nil
}
// InsertURL inserts a new URL mapping into the database.
func InsertURL(shortURL, longURL string)error{
	query := `INSERT INTO urls (short_url, long_url) VALUES ($1, $2)`;
	stmt,err := db.DB.Prepare(query);
	if err != nil{
		return err;
	}
	defer stmt.Close();
	_, err = stmt.Exec(shortURL, longURL);
	if err != nil{	
		return err;
	}
	return nil;
}
