package service

import (
	"database/sql"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

type AlbumService struct {
	db *sql.DB
}

func NewAlbumService(db *sql.DB) *AlbumService {
	return &AlbumService{db: db}
}

func (s *AlbumService) CreateAlbum(album *Album) error {
	query := "INSERT INTO ALBUMS (title, artist, price) values(?,?,?)"

	result, err := s.db.Exec(query, album.Title, album.Artist, album.Price)

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	album.ID = int(lastInsertId)

	return nil
}

func (s *AlbumService) GetAlbums() ([]Album, error) {
	query := "SELECT id, title, artist, price FROM albums"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	var albums []Album

	for rows.Next() {
		var album Album

		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(id int) (*Album, error) {
	query := "SELECT id, title, artist, price FROM albums WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var album Album

	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (s *AlbumService) UpdateAlbum(album *Album) error {
	query := "UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?"

	_, err := s.db.Exec(query, album.Title, album.Artist, album.Price, album.ID)

	return err
}

func (s *AlbumService) DeleteAlbum(id int) error {
	query := "DELETE FROM albums WHERE id = ?"

	_, err := s.db.Exec(query, id)

	return err
}

func (s *AlbumService) SearchAlbumByTitle(title string) ([]Album, error) {
	query := "SELECT id, title, artist, price FROM albuns WHERE title like ?"

	rows, err := s.db.Query(query, "%"+title+"%")

	if err != nil {
		return nil, err

	}

	var albums []Album

	for rows.Next() {
		var album Album

		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

		if err != nil {
			return nil, err

		}

		albums = append(albums, album)
	}

	return albums, nil
}

func (s *AlbumService) SearchAlbumByArtist(artist string) ([]Album, error) {
	query := "SELECT id, title, artist, price FROM albums WHERE artist LIKE ?"

	rows, err := s.db.Query(query, "%"+artist+"%")

	if err != nil {
		return nil, err
	}

	var albums []Album

	for rows.Next() {
		var album Album

		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}
	return albums, nil
}
