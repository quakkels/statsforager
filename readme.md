# StatsForager

- [Mission](#Mission)
- [Work](#Work)
- [Development](#Development)
- [License](License)

## Mission

1. Make a website analytics service that respects site owner's data, user privacy, and doesn't depend on Big Tech.
2. Promote an independent internet by offering alternatives to Big Tech services.

## Development

Dependencies:

- Go 1.22
- Postgres
- NPM

Development Environment:

- Neovim
- Docker
- Bash
- Tmux

The script [start_tmux_devenv.sh](start_tmux_devenv.sh) will fire up a Tmux session with three windows.

1. "Data" window: runs Postgress via Docker in one pane, and connects to it via `psql` in another
3. "Web" window: runs the built Go project alongside an open Neovim editor
2. "Client" window: runs a simulated a third party consumer's website using the StatsForager service

## Work

[Todo](todo.md)

## License

MIT license
