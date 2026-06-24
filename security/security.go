package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Params define as configurações de custo do Argon2id
type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func Validate(password string) {
	// 1. set argon params
	params := &Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	// 2. generate hash
	encodedHash, err := GenerateHash(password, params)
	if err != nil {
		log.Fatalf("Erro ao gerar hash: %v", err)
	}
	fmt.Printf("Hash Gerado: %s\n\n", encodedHash)

	// 2. Validar correspondência (Sucesso)
	valido, err := ComparePassword(senha, encodedHash)
	if err != nil {
		log.Fatalf("Erro ao comparar: %v", err)
	}
	fmt.Printf("Senha correta? %t\n", valido)

	// 3. Validar correspondência (Falha)
	valido, _ = ComparePassword("SenhaErrada", encodedHash)
	fmt.Printf("Senha errada detectada? %t\n", !valido)
}

// GenerateHash processa a senha e retorna uma string formatada do Argon2
func GenerateHash(password string, p *Params) (string, error) {
	// Gerar Salt Criptograficamente Seguro
	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Executar o algoritmo Argon2id
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Codificar em formato string padrão ($argon2id$v=19$m=65536,t=3,p=4$...)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encoded, nil
}

// ComparePassword extrai os parâmetros da string gravada e valida a tentativa
func ComparePassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("formato de hash inválido")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, fmt.Errorf("versão incompatível do argon2")
	}

	var params Params
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Computar o hash da senha fornecida usando os mesmos parâmetros extraídos
	otherHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, uint32(len(hash)))

	// Comparação em Tempo Constante contra ataques de temporização (Timing Attacks)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}
