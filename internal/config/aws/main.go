package aws

import (
	"github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/aws_session"
	sns_client "github.com/julioceno/desafio-anotaai-backend-golang/internal/config/aws/sns"
)

func NewHandler() {
	aws_session.NewHandler()

	sns_client.NewHandler()
}
