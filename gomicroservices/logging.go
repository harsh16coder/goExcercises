package main

import (
	"context"
	"fmt"
	"time"
)

type LogginService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LogginService{
		next: next,
	}
}

func (s *LogginService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%s took=%v", fact.Fact, err, time.Since(start))
	}(time.Now())
	return s.next.GetCatFact(ctx)
}
