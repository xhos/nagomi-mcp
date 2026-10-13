# nagomi-mcp

An [MCP](https://modelcontextprotocol.io) server that gives AI agents read access to one user's [nagomi](https://github.com/xhos/nagomi) data.

> [!IMPORTANT]
> There is no auth. Run it locally or behind a VPN.

## config

See [.env.example](.env.example).

## clients

It serves SSE at `<NAGOMI_MCP_BASE_URL>/sse`. Clients that only support stdio, like Claude Desktop, need `mcp-proxy`:

```json
{
  "mcpServers": {
    "nagomi": {
      "command": "uvx",
      "args": ["mcp-proxy", "https://mcp.example.com/sse"]
    }
  }
}
```

`npx -y mcp-proxy` or `nix run nixpkgs#mcp-proxy --` work too.

## dev

```bash
run
tst    # tstv for verbose
fmt
regen
bump-protos
```
