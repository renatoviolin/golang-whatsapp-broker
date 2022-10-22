package postgres

import (
	"broker/dto"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AgentRepository struct {
	dbPool *pgxpool.Pool
}

func NewAgentRepository(conn string) (*AgentRepository, error) {
	dbPool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = dbPool.Ping(ctx); err != nil {
		return nil, err
	}

	return &AgentRepository{dbPool: dbPool}, nil
}

func (m *AgentRepository) FindById(id int) (agent dto.Agent, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select id, name, read_topic, write_topic, error_topic from agents where id = $1 and active=1`
	rows := m.dbPool.QueryRow(ctx, query, id)
	err = rows.Scan(&agent.ID, &agent.Name, &agent.ReadTopic, &agent.WriteTopic, &agent.ErrorTopic)
	if err != nil || agent.ID == 0 {
		return agent, errors.New("agent not found")
	}
	return agent, nil
}

func (m *AgentRepository) All() (agents []*dto.Agent, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select id, name, read_topic, write_topic, error_topic from agents where active=1`
	rows, err := m.dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var agent dto.Agent
		if err := rows.Scan(&agent.ID, &agent.Name, &agent.ReadTopic, &agent.WriteTopic, &agent.ErrorTopic); err != nil {
			return nil, err
		}
		agents = append(agents, &agent)
	}
	if len(agents) == 0 {
		return agents, errors.New("agent not found")
	}

	return agents, nil
}
