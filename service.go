package main

import (
	"context"

	common "buf.build/gen/go/nonsense/nonsense-api/protocolbuffers/go/agent/common"
	resources "buf.build/gen/go/nonsense/nonsense-api/protocolbuffers/go/agent/resources"
	services "buf.build/gen/go/nonsense/nonsense-api/protocolbuffers/go/agent/services"
	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (s *AgentServer) CreateAgent(
	ctx context.Context,
	req *connect.Request[services.CreateAgentRequest],
) (*connect.Response[services.CreateAgentResponse], error) {
	agent := &resources.Agent{
		Id:          uuid.New().String(),
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
	}

	return connect.NewResponse(&services.CreateAgentResponse{
		Agent: agent,
	}), nil
}

func (s *AgentServer) GetAgent(
	ctx context.Context,
	req *connect.Request[services.GetAgentRequest],
) (*connect.Response[services.GetAgentResponse], error) {
	agent := &resources.Agent{
		Id:          req.Msg.Id,
		Name:        "Sample Agent",
		Description: "Sample Description",
	}

	return connect.NewResponse(&services.GetAgentResponse{
		Agent: agent,
	}), nil
}

func (s *AgentServer) ListAgents(
	ctx context.Context,
	req *connect.Request[services.ListAgentsRequest],
) (*connect.Response[services.ListAgentsResponse], error) {
	agents := []*resources.Agent{
		{
			Id:          "1",
			Name:        "Agent 1",
			Description: "Description 1",
		},
		{
			Id:          "2",
			Name:        "Agent 2",
			Description: "Description 2",
		},
	}

	return connect.NewResponse(&services.ListAgentsResponse{
		Agents:        agents,
		NextPageToken: "",
	}), nil
}

func (s *AgentServer) UpdateAgent(
	ctx context.Context,
	req *connect.Request[services.UpdateAgentRequest],
) (*connect.Response[services.UpdateAgentResponse], error) {
	agent := &resources.Agent{
		Id:          req.Msg.Id,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
	}

	return connect.NewResponse(&services.UpdateAgentResponse{
		Agent: agent,
	}), nil
}

func (s *AgentServer) DeleteAgent(
	ctx context.Context,
	req *connect.Request[services.DeleteAgentRequest],
) (*connect.Response[common.Empty], error) {
	return connect.NewResponse(&common.Empty{}), nil
}
