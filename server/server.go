package server

import (
	"context"
	"log"
	"sync"

	linksaver "github.com/Gipohub/linksaver/proto"
)

type linkSaverServer struct {
	linksaver.UnimplementedLinkSaverServer
	mu    sync.Mutex
	pages map[string]map[string]struct{} // username -> set of URLs
}

func New() *linkSaverServer {
	return &linkSaverServer{
		pages: make(map[string]map[string]struct{}),
	}
}

func (s *linkSaverServer) Save(ctx context.Context, req *linksaver.SaveRequest) (*linksaver.SaveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Printf("i am in saveGRPCserver ")
	if s.pages[req.Username] == nil {
		s.pages[req.Username] = make(map[string]struct{})
	}
	s.pages[req.Username][req.Url] = struct{}{}

	return &linksaver.SaveResponse{Status: "ok"}, nil
}

func (s *linkSaverServer) IsExists(ctx context.Context, page *linksaver.Page) (*linksaver.ExistsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.pages[page.Username][page.Url]
	return &linksaver.ExistsResponse{Exists: exists}, nil
}

func (s *linkSaverServer) PickRandom(ctx context.Context, user *linksaver.User) (*linksaver.Page, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for url := range s.pages[user.Username] {
		return &linksaver.Page{Url: url, Username: user.Username}, nil
	}
	return nil, nil
}

func (s *linkSaverServer) PickAll(ctx context.Context, user *linksaver.User) (*linksaver.PageList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := &linksaver.PageList{}
	for url := range s.pages[user.Username] {
		list.Pages = append(list.Pages, &linksaver.Page{
			Url:      url,
			Username: user.Username,
		})
	}
	return list, nil
}
