package tools

import (
	"github.com/charmbracelet/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/xhos/nagomi-mcp/internal/gen/nagomi/v1/nagomiv1connect"
)

// Handler holds the Connect RPC clients and registers all MCP tools.
type Handler struct {
	userID    string
	accounts  nagomiv1connect.AccountServiceClient
	txns      nagomiv1connect.TransactionServiceClient
	cats      nagomiv1connect.CategoryServiceClient
	dashboard nagomiv1connect.DashboardServiceClient
	log       *log.Logger
}

func New(
	userID string,
	accounts nagomiv1connect.AccountServiceClient,
	txns nagomiv1connect.TransactionServiceClient,
	cats nagomiv1connect.CategoryServiceClient,
	dashboard nagomiv1connect.DashboardServiceClient,
	log *log.Logger,
) *Handler {
	return &Handler{
		userID:    userID,
		accounts:  accounts,
		txns:      txns,
		cats:      cats,
		dashboard: dashboard,
		log:       log,
	}
}

func (h *Handler) Register(s *server.MCPServer) {
	h.registerAccounts(s)
	h.registerCategories(s)
	h.registerTransactions(s)
	h.registerDashboard(s)
}

func (h *Handler) result(msg proto.Message) (*mcp.CallToolResult, error) {
	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (h *Handler) errResult(err error) (*mcp.CallToolResult, error) {
	h.log.Error("tool call failed", "err", err)
	return mcp.NewToolResultError(err.Error()), nil
}
