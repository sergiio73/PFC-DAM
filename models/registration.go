package models

import "errors"

type Registration struct {
	ID          int
	GameID      int
	CharacterID int
	PlayerName  string
	Status      string
}

type RegistrationView struct {
	ID            int
	PlayerName    string
	Status        string
	CharacterName string
	Class         string
	Race          string
	Level         int
}

func CreateRegistration(registration Registration) error {
	// Evito que el mismo personaje se apunte dos veces a la misma partida.
	for _, item := range Registrations {
		if item.GameID == registration.GameID && item.CharacterID == registration.CharacterID {
			return errors.New("este personaje ya está apuntado a esta partida")
		}
	}

	registration.ID = NextRegistrationID
	NextRegistrationID++
	registration.Status = "pendiente"

	Registrations = append(Registrations, registration)

	return nil
}

func GetRegistrationsByGame(gameID int) ([]RegistrationView, error) {
	result := []RegistrationView{}

	for _, registration := range Registrations {
		if registration.GameID == gameID {
			character, err := findCharacterByID(registration.CharacterID)
			if err != nil {
				// Si por lo que sea no encuentra el personaje, saltamos ese registro.
				continue
			}

			view := RegistrationView{
				ID:            registration.ID,
				PlayerName:    registration.PlayerName,
				Status:        registration.Status,
				CharacterName: character.Name,
				Class:         character.Class,
				Race:          character.Race,
				Level:         character.Level,
			}

			result = append(result, view)
		}
	}

	return result, nil
}

func UpdateRegistrationStatus(registrationID int, status string) error {
	for i := 0; i < len(Registrations); i++ {
		if Registrations[i].ID == registrationID {
			Registrations[i].Status = status
			return nil
		}
	}

	return errors.New("inscripción no encontrada")
}

func CountAcceptedPlayers(gameID int) (int, error) {
	count := 0

	for _, registration := range Registrations {
		if registration.GameID == gameID && registration.Status == "aceptado" {
			count++
		}
	}

	return count, nil
}

func findCharacterByID(id int) (Character, error) {
	for _, character := range Characters {
		if character.ID == id {
			return character, nil
		}
	}

	return Character{}, errors.New("personaje no encontrado")
}
