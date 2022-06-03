package controller

// original page:
// https://dev.to/techschoolguru/how-to-create-and-verify-jwt-paseto-token-in-golang-1l5j

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var key string = "12345678876543211234567887654321"
var maker, initErr = NewPasetoMaker(key)

// will it ensure single-instance?
func GetGlobalMaker () Maker {
	if initErr != nil {
		fmt.Printf("failed to init global token maker, err:%e\n", initErr)
		return nil
	}
	return maker
}

type Maker interface{
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

// token info 
type Payload struct {
	ID uuid.UUID
	Username string
	IssuedAt time.Time
	ExpiredAt time.Time
}

func NewPayload (username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("failed to get new uuid for new payload, err: %e\n", err)
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (py *Payload) Valid () (error) {
	if time.Now().After(py.ExpiredAt){
		return fmt.Errorf("token expired")
	}
	return nil
}

func (py *Payload) get_username() string {
	return py.Username
}
// token gen
type PasetoMaker struct {
	paseto *paseto.V2
	symmetricKey []byte
}
func NewPasetoMaker (symmetricKey string) (Maker, error){
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := PasetoMaker{
		paseto: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		fmt.Printf("failed to new payload for new token, err: %e\n", err)
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		fmt.Printf("failed to verify token, err: %e\n", err)
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		fmt.Printf("payload err: %e\n", err)
		return nil, err
	}

	return payload, nil
}
