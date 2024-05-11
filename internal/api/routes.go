package api

import (
	"git.doolta.com/doolta/go-kit/web"
)

func Routes(s *web.Server) {
	s.Log.Info("Installing API routes")
	s.AddRoute("POST /api/v1/members", s.WithServerContext(CreateMemberHandlerFunc))
	s.AddRoute("GET /api/v1/members/{id}", s.WithServerContext(ReadMemberHandlerFunc))
	s.AddRoute("PUT /api/v1/members/{id}", s.WithServerContext(UpdateMemberHandlerFunc))
	s.AddRoute("DELETE /api/v1/members/{id}", s.WithServerContext(DeleteMemberHandlerFunc))
}
