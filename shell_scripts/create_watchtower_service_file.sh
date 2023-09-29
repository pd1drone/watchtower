#!/bin/bash

# Create the file watchtower.service with the specified content
cat << EOF > /etc/systemd/system/watchtower.service
[Unit]
Description=bims backend http server
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/watchtower/
ExecStart=/root/watchtower/cmd/watchtower
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the watchtower.service file so it starts on boot
sudo systemctl enable watchtower.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the watchtower.service
sudo systemctl start watchtower.service
