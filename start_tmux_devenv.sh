#!/bin/bash

SESSION_NAME="StatsForager"

if ! tmux has-session -t $SESSION_NAME 2>/dev/null; then

	# database window
	tmux new-session -d -s $SESSION_NAME -n "statsforagerdata"
	tmux split-window -h -t $SESSION_NAME:0
	tmux send-keys -t $SESSION_NAME:0.0 'cd ./statsforagerdata' C-m
	tmux send-keys -t $SESSION_NAME:0.0 './dev_runDocker.sh' C-m
	tmux send-keys -t $SESSION_NAME:0.0 'cat psql.md' C-m
	tmux send-keys -t $SESSION_NAME:0.1 'cd ./statsforagerdata' C-m
	tmux send-keys -t $SESSION_NAME:0.1 'psql stats -h localhost -U postgres'

	# development window
	tmux new-window -t $SESSION_NAME -n "statsforagerweb"
	tmux send-keys -t $SESSION_NAME:1 'cd ./statsforagerweb' C-m
	tmux send-keys -t $SESSION_NAME:1 'nvim .' C-m
	tmux split-window -h -t $SESSION_NAME:2
	tmux send-keys -t $SESSION_NAME:1.1 'make run-native' C-m

	# client window
	tmux new-window -t $SESSION_NAME -n "statsforagerclient"
	tmux send-keys -t $SESSION_NAME:2 -n 'cd ./statsforagerclient' C-m
	tmux send-keys -t $SESSION_NAME:2 -n 'npm run dev' C-m
fi

# Attach to the session and select the 'database' window
tmux attach-session -t $SESSION_NAME
