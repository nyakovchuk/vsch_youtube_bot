package command

type Command struct {
	Name        string
	Route       string
	Description string
}

type Commands []Command

var commandList = Commands{
	{"start", "/start", "About the Bot"},
}

func GetCommands() CommandManager {
	return commandList
}

func (c Commands) Get() Commands {
	return commandList
}

func (c Commands) GetByName(name string) Command {
	for _, command := range c.Get() {
		if command.Name == name {
			return command
		}
	}
	return Command{}
}

func (c Commands) GetByRoute(route string) Command {
	for _, command := range c.Get() {
		if command.Route == route {
			return command

		}
	}
	return Command{}
}
