package models

import "errors"

type Character struct {
	ID      int
	Name    string
	Class   string
	Race    string
	Level   int
	History string
	Owner   string // nombre de usuario sacado de la cookie
}

func CreateCharacter(character Character) error {
	character.ID = NextCharacterID
	NextCharacterID++
	Characters = append(Characters, character)
	return nil
}

func GetAllCharacters() ([]Character, error) {
	charactersCopy := []Character{}
	for i := len(Characters) - 1; i >= 0; i-- {
		charactersCopy = append(charactersCopy, Characters[i])
	}
	return charactersCopy, nil
}

// Solo devuelve los personajes del usuario que ha iniciado sesión.
func GetCharactersByOwner(owner string) ([]Character, error) {
	charactersCopy := []Character{}
	for i := len(Characters) - 1; i >= 0; i-- {
		if Characters[i].Owner == owner {
			charactersCopy = append(charactersCopy, Characters[i])
		}
	}
	return charactersCopy, nil
}

func DeleteCharacter(id int, owner string) error {
	for i, c := range Characters {
		if c.ID == id {
			if c.Owner != owner {
				return errors.New("este personaje no te pertenece")
			}
			Characters = append(Characters[:i], Characters[i+1:]...)
			return nil
		}
	}
	return errors.New("personaje no encontrado")
}
