module github.com/TicketsBot/translationgenerator

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/TicketsBot/database v0.0.0-20200710191148-306959ba00d2
	github.com/gin-gonic/gin v1.6.3
)

replace github.com/TicketsBot/database => ../database
