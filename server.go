package gss

import (
	"errors"

	"github.com/apcera/gssapi"
)

func NewSSHGSSAPIServerSide() (*sshGSSApiServerSide, error) {
	lib, err := gssapi.Load(nil)
	if err != nil {
		return nil, err
	}
	return &sshGSSApiServerSide{
		lib: lib,
	}, nil
}

type sshGSSApiServerSide struct {
	lib *gssapi.Lib
	ctx *gssapi.CtxId
}

func (s *sshGSSApiServerSide) AcceptSecContext(token []byte) ([]byte, string, bool, error) {
	inputToken, err := s.lib.MakeBufferBytes(token)
	defer inputToken.Release()
	if err != nil {
		return nil, "", false, err
	}
	ctx, srcName, _, outToken, _, _, _, err := s.lib.AcceptSecContext(s.lib.GSS_C_NO_CONTEXT, s.lib.GSS_C_NO_CREDENTIAL, inputToken, s.lib.GSS_C_NO_CHANNEL_BINDINGS)
	defer outToken.Release()
	defer srcName.Release()
	s.ctx = ctx
	if err != nil {
		if err == gssapi.ErrContinueNeeded {
			return outToken.Bytes(), "", true, nil
		}
		return outToken.Bytes(), "", false, err
	}
	return outToken.Bytes(), srcName.String(), false, nil
}

func (s *sshGSSApiServerSide) VerifyMIC(micField []byte, micToken []byte) error {
	if s.ctx == nil {
		return errors.New("ctx is nil, acceptSecContext before VerifyMIC")
	}
	messageBuffer, _ := s.lib.MakeBufferBytes(micField)
	defer messageBuffer.Release()
	tokenBuffer, _ := s.lib.MakeBufferBytes(micToken)
	defer tokenBuffer.Release()
	if _, err := s.ctx.VerifyMIC(messageBuffer, tokenBuffer); err != nil {
		return err
	}
	return nil
}

func (s *sshGSSApiServerSide) DeleteSecContext() error {
	if s.ctx != nil {
		s.ctx.Release()
	}
	return nil
}
