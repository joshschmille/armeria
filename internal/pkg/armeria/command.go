package armeria

import (
	"fmt"
	"strings"
)

type Command struct {
	Name         string
	AltNames     []string
	Help         string
	Alias        string
	Permissions  *CommandPermissions
	AllowedRoles []int
	Arguments    []*CommandArgument
	Subcommands  []*Command
	Handler      func(r *CommandContext)
}

type CommandArgument struct {
	Name             string
	IncludeRemaining bool
	Optional         bool
}

type CommandPermissions struct {
	RequireNoCharacter bool
	RequireCharacter   bool
	RequirePermission  string
}

type CommandContext struct {
	Command         *Command
	Player          *Player
	PlayerInitiated bool
	Character       *Character
	Args            map[string]string
}

// CheckPermissions returns whether or not a player can use the command
func (cmd *Command) CheckPermissions(p *Player) bool {
	if cmd.Permissions == nil {
		return true
	}

	if cmd.Permissions.RequireNoCharacter {
		if p.Character() != nil {
			return false
		}
	}

	if cmd.Permissions.RequireCharacter {
		if p.Character() == nil {
			return false
		}
	}

	if len(cmd.Permissions.RequirePermission) > 0 {
		if !p.Character().HasPermission(cmd.Permissions.RequirePermission) {
			return false
		}
	}
	return true
}

// ShowSubcommandHelp returns the list of sub-commands that the player has access to as a string
func (cmd *Command) ShowSubcommandHelp(p *Player, commandsEntered []string) string {
	if len(cmd.Subcommands) == 0 {
		return "There are no sub-commands available."
	}

	output := []string{
		"[b]Help:[/b]",
		"  " + cmd.Help,
		fmt.Sprintf("  [b]Syntax:[/b] /%s &lt;sub-command&gt;\n", strings.Join(commandsEntered, " ")),
		"[b]Sub-commands:[/b]",
	}

	var allowedSubCommands []*Command
	var longestCommandSize int
	for _, scmd := range cmd.Subcommands {
		if cmd.CheckPermissions(p) {
			if len(scmd.Name) > longestCommandSize {
				longestCommandSize = len(scmd.Name)
			}
			allowedSubCommands = append(allowedSubCommands, scmd)
		}
	}

	for _, scmd := range allowedSubCommands {
		output = append(output, fmt.Sprintf("  %-10v %s", scmd.Name, scmd.Help))
	}

	return strings.Join(output, "\n")
}

func (cmd *Command) ShowArgumentHelp(p *Player, commandsEntered []string) string {
	if len(cmd.Arguments) == 0 {
		return "There are no command arguments."
	}

	var argumentStrings []string
	for _, arg := range cmd.Arguments {
		if arg.Optional {
			argumentStrings = append(argumentStrings, fmt.Sprintf("[%s]", arg.Name))
		} else {
			argumentStrings = append(argumentStrings, fmt.Sprintf("&lt;%s&gt;", arg.Name))
		}

	}

	output := []string{
		"[b]Help:[/b]",
		"  " + cmd.Help,
		fmt.Sprintf(
			"  [b]Syntax:[/b] /%s %s\n",
			strings.Join(commandsEntered, " "),
			strings.Join(argumentStrings, " "),
		),
	}

	return strings.Join(output, "\n")
}
