package comments

import (
	"errors"
	"strings"
)

type Repository interface {
	FindAll() ([]Comment, error)
	FindById(id int) (Comment, error)
	Save(comment NewComment) (Comment, error)
	Delete(id int) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAllComments() ([]Comment, error) {
	comments, err := s.repo.FindAll()
	if err != nil {
		return []Comment{}, err
	}
	if comments == nil {
		return []Comment{}, nil
	}

	return comments, nil
}

func (s *Service) CreateComment(data NewComment) (Comment, error) {
	data.Username = strings.TrimSpace(data.Username)
	data.Message = strings.TrimSpace(data.Message)

	if data.Username == "" || data.Message == "" {
		return Comment{}, errors.New("el nombre de usuario y el mensaje son obligatorios")
	}

	if len(data.Message) < 5 {
		return Comment{}, errors.New("el mensaje es demasiado corto (mínimo 5 caracteres)")
	}

	return s.repo.Save(data)
}

func (s *Service) GetCommentById(id int) (Comment, error) {
	if id <= 0 {
		return Comment{}, errors.New("el ID debe ser un número positivo")
	}
	return s.repo.FindById(id)
}

func (s *Service) RemoveComment(id int) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}