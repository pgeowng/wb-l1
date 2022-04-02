package main

import "fmt"

// Паттерн "Адаптер"

// Позволяет использовать объект недоступный для модификации,
// через специально созданный интерфейс.
// Например, у нас есть старый код для работы с сервисом.
// По некоторым причинам(производительность, удобство, расширяемость)
// мы хотим сменить интерфейс взаимодействия с сервисом.
// Допустим одна команда работает над кодом нового сервиса,
// другая занята остальными модулями.
// Чтобы начать переход вне зависимости от готовности сервиса используем адаптер.

type INewService interface {
	Send(name string, owner string, payload []byte) error
}

type OldService struct{}

func NewOldService() *OldService {
	return &OldService{}
}

func (s *OldService) Upload(payload []byte) string {
	fmt.Println("old service upload:", string(payload))
	return "x4d3c889"
}

type OldServiceAdapter struct {
	*OldService
}

// То же поведение, только другой вызов методов объекта.
// type OldServiceAdapter struct {
// 	OldService *OldService
// }

func NewOldServiceAdapter(oldService *OldService) INewService {
	return &OldServiceAdapter{oldService}
}

func (s *OldServiceAdapter) Send(name string, owner string, payload []byte) error {
	fmt.Println("old service adapter upload:", string(payload))
	s.Upload(payload)
	//s.OldService.Upload(payload)
	return nil
}

func main() {
	var osa INewService = NewOldServiceAdapter(NewOldService())

	if err := osa.Send("general", "client-web", []byte("ping")); err != nil {
		fmt.Println("err: ", err)
	}
}
