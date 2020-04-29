# Linux Word Reminder

### Dependence
- `go get github.com/claudiu/gocron`
- `go get github.com/bregydoc/gtranslate`

### Build
- `go build -o word-notify`

### Manual Install
- `sudo cp -r word-notify/ /opt`
- `sudo cp /opt/word-notify/word-notify.service /etc/systemd/system`
- `sudo systemctl start word-notify.service`
- `sudo systemctl restart word-notify.service`
