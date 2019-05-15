package gss

import (
	"errors"

	"github.com/apcera/gssapi"
)

func NewSSHGSSAPIClientSide() (*sshGSSAPIClientSide, error) {
	lib, err := gssapi.Load(nil)
	if err != nil {
		return nil, err
	}
	return &sshGSSAPIClientSide{
		lib: lib,
	}, nil
}

type sshGSSAPIClientSide struct {
	lib *gssapi.Lib
	ctx *gssapi.CtxId
}

func (s *sshGSSAPIClientSide) InitSecContext(target string, token []byte, isGSSDelegCreds bool) ([]byte, bool, error) {
	targetBuffer, err := s.lib.MakeBufferString(target)
	defer targetBuffer.Release()
	if err != nil {
		return nil, false, err
	}
	targetName, err := targetBuffer.Name(s.lib.GSS_C_NT_HOSTBASED_SERVICE)
	defer targetName.Release()
	if err != nil {
		return nil, false, err
	}
	var regFlags uint32
	regFlags = gssapi.GSS_C_PROT_READY_FLAG | gssapi.GSS_C_INTEG_FLAG | gssapi.GSS_C_MUTUAL_FLAG
	if isGSSDelegCreds {
		regFlags |= gssapi.GSS_C_DELEG_FLAG
	}
	var tokenBuffer *gssapi.Buffer
	if token != nil {
		var err error
		tokenBuffer, err = s.lib.MakeBufferBytes(token)
		defer tokenBuffer.Release()
		if err != nil {
			return nil, false, err
		}
	} else {
		tokenBuffer = s.lib.GSS_C_NO_BUFFER
	}
	var ctxIn *gssapi.CtxId
	if s.ctx == nil {
		ctxIn = s.lib.GSS_C_NO_CONTEXT
	} else {
		ctxIn = s.ctx
	}

	ctx, _, outToken, _, _, err := s.lib.InitSecContext(s.lib.GSS_C_NO_CREDENTIAL,
		ctxIn, targetName, s.lib.GSS_MECH_KRB5, regFlags, 0, s.lib.GSS_C_NO_CHANNEL_BINDINGS, tokenBuffer)
	defer outToken.Release()
	s.ctx = ctx
	if err != nil {
		if err == gssapi.ErrContinueNeeded {
			return outToken.Bytes(), true, nil
		}
		return outToken.Bytes(), false, err
	}
	return outToken.Bytes(), false, nil
}

func (s *sshGSSAPIClientSide) GetMIC(micFiled []byte) ([]byte, error) {
	if s.ctx == nil {
		return nil, errors.New("ctx is nil, call InitSecContext before GetMIC")
	}
	messageBuffer, err := s.lib.MakeBufferBytes(micFiled)
	defer messageBuffer.Release()
	if err != nil {
		return nil, err
	}
	messageToken, err := s.ctx.GetMIC(0, messageBuffer)
	defer messageToken.Release()
	if err != nil {
		return nil, err
	}
	return messageToken.Bytes(), nil
}

func (s *sshGSSAPIClientSide) DeleteSecContext() error {
	if s.ctx != nil {
		s.ctx.Release()
	}
	return nil
}
