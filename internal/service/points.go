package service

import (
    "time"
    "responsible_employee/internal/model"
    "responsible_employee/internal/repository"
)

type PointsService struct {
    repoUser      repository.User
    repoPointEvent repository.PointEvent
}

func NewPointsService(repoUser repository.User, repoPointEvent repository.PointEvent) *PointsService {
    return &PointsService{repoUser: repoUser, repoPointEvent: repoPointEvent}
}

func (s *PointsService) Summary(month int, year int) ([]model.UserPointsBreakdown, error) {
    if year == 0 {
        year = time.Now().Year()
    }
    if month < 1 || month > 12 {
        month = int(time.Now().Month())
    }

    startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

    monthlyEvents, err := s.repoPointEvent.EventsSince(startOfMonth)
    if err != nil {
        return nil, err
    }
    yearlyEvents, err := s.repoPointEvent.EventsSince(startOfYear)
    if err != nil {
        return nil, err
    }

    users, err := s.repoUser.GetUsersSortedByPoints()
    if err != nil {
        return nil, err
    }

    // Build maps for aggregation
    monthlyByUser := map[string]map[string]int{}
    monthlyTotals := map[string]int{}
    for _, e := range monthlyEvents {
        if _, ok := monthlyByUser[e.UserID]; !ok {
            monthlyByUser[e.UserID] = map[string]int{}
        }
        monthlyByUser[e.UserID][e.Source] += e.Points
        monthlyTotals[e.UserID] += e.Points
    }

    yearlyByUser := map[string]map[string]int{}
    yearlyTotals := map[string]int{}
    for _, e := range yearlyEvents {
        if _, ok := yearlyByUser[e.UserID]; !ok {
            yearlyByUser[e.UserID] = map[string]int{}
        }
        yearlyByUser[e.UserID][e.Source] += e.Points
        yearlyTotals[e.UserID] += e.Points
    }

    var result []model.UserPointsBreakdown
    for _, u := range users {
        result = append(result, model.UserPointsBreakdown{
            UserID:          u.ID,
            Login:           u.Login,
            FullName:        u.FullName,
            MonthlyTotal:    monthlyTotals[u.ID],
            YearlyTotal:     yearlyTotals[u.ID],
            MonthlyBySource: monthlyByUser[u.ID],
            YearlyBySource:  yearlyByUser[u.ID],
        })
    }

    return result, nil
}


