package kms

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// EncryptionContext define the format required for kms encryption context
type EncryptionContext map[string]*string

// Handler Structure encapsulating stuff common to encrypt and decrypt.
type Handler struct {
	// Service provide interfaces for AWS SDK features
	Service AWSIface

	// Context defines encryption context which is a set of non-secret key-value pairs that you can pass to AWS KMS
	// For more information see https://docs.aws.amazon.com/kms/latest/developerguide/encryption-context.html
	Context []string

	// CipherKey is the KMS encrypted hash key
	CipherKey string
}

// Client prepare AWS config
func Client(region string) *kms.KMS {
	return kms.New(session.New(), aws.NewConfig().WithRegion(region))
}

// AWSIface abstract AWS SDK required method
type AWSIface interface {
	Decrypt(*kms.DecryptInput) (*kms.DecryptOutput, error)
}

// ParseEncryptionContext encryption context is required to decrypt the data
func (h *Handler) ParseEncryptionContext() (EncryptionContext, error) {
	context := make(EncryptionContext, len(h.Context))
	for _, s := range h.Context {
		parts := strings.SplitN(s, "=", 2)
		if len(parts) < 2 {
			return nil, fmt.Errorf("context must be provided in NAME=VALUE format")
		}
		context[parts[0]] = &parts[1]
	}
	return context, nil
}

// Decrypt ciphertext.
func (h *Handler) Decrypt() (string, error) {

	ciphertextBlob, err := base64.StdEncoding.DecodeString(h.CipherKey)
	if err != nil {
		return "", err
	}
	ec, err := h.ParseEncryptionContext()
	if err != nil {
		return "", err
	}
	output, err := h.Service.Decrypt(&kms.DecryptInput{
		EncryptionContext: ec,
		CiphertextBlob:    ciphertextBlob,
	})
	if err != nil {
		return "", err
	}
	return string(output.Plaintext), nil
}
