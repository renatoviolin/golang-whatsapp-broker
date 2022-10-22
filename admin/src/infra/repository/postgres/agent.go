package postgres

import (
	"context"
	"errors"
	"time"
	"whatsapp-admin/dto"

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

func (h *AgentRepository) Create(agent dto.Agent) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `	INSERT INTO agents
				(name, read_topic, write_topic, error_topic, active)
				VALUES ($1, $2, $3, $4, $5) RETURNING "id"`
	if err = h.dbPool.QueryRow(ctx, query, agent.Name, agent.ReadTopic, agent.WriteTopic, agent.ErrorTopic, 1).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (h *AgentRepository) Update(agent dto.Agent) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `	UPDATE agents
				SET name=$1, read_topic=$2, write_topic=$3, error_topic=$4
				WHERE id = $5`
	res, err := h.dbPool.Exec(ctx, query, agent.Name, agent.ReadTopic, agent.WriteTopic, agent.ErrorTopic, agent.ID)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (h *AgentRepository) Delete(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `UPDATE agents SET active=0 WHERE id = $1 and active=1`
	res, err := h.dbPool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("no rows affected")
	}

	return nil
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

func (m *AgentRepository) FindByName(name string) (agents []*dto.Agent, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `select id, name, read_topic, write_topic, error_topic from agents where name like $1 and active=1`
	rows, err := m.dbPool.Query(ctx, query, "%"+name+"%")
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
func (m *AgentRepository) All2() (agents []dto.Agent, err error) {
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
		agents = append(agents, agent)
	}
	if len(agents) == 0 {
		return agents, errors.New("agent not found")
	}

	return agents, nil
}
