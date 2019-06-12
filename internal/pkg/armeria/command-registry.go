package armeria

func RegisterGameCommands() {
	commands := []*Command{
		{
			Name: "login",
			Help: "Logs your character into the game world.",
			Permissions: &CommandPermissions{
				RequireNoCharacter: true,
			},
			Arguments: []*CommandArgument{
				{
					Name: "character",
				},
				{
					Name: "password",
				},
			},
			Handler: handleLoginCommand,
		},
		{
			Name: "look",
			Help: "Looks at something.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Handler: handleLookCommand,
		},
		{
			Name: "say",
			Help: "Says something to everyone in the same room as you.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Arguments: []*CommandArgument{
				{
					Name:             "text",
					IncludeRemaining: true,
				},
			},
			Handler: handleSayCommand,
		},
		{
			Name: "move",
			Help: "Moves your character into a connecting room.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Arguments: []*CommandArgument{
				{
					Name: "direction",
				},
			},
			Handler: handleMoveCommand,
		},
		{Name: "north", Alias: "move north"},
		{Name: "south", Alias: "move south"},
		{Name: "east", Alias: "move east"},
		{Name: "west", Alias: "move west"},
		{Name: "up", Alias: "move up"},
		{Name: "down", Alias: "move down"},
		{
			Name: "room",
			Help: "Allows you to manage rooms.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Subcommands: []*Command{
				{
					Name:    "edit",
					Help:    "Opens the editor panel for the current room.",
					Handler: handleRoomEditCommand,
				},
				{
					Name: "set",
					Help: "Allows you to set a room attribute. Leave value empty to revert to default.",
					Arguments: []*CommandArgument{
						{
							Name: "property",
						},
						{
							Name:             "value",
							IncludeRemaining: true,
							Optional:         true,
						},
					},
					Handler: handleRoomSetCommand,
				},
				{
					Name: "create",
					Help: "Creates a new room in a particular direction.",
					Arguments: []*CommandArgument{
						{
							Name: "direction",
						},
					},
					Handler: handleRoomCreateCommand,
				},
				{
					Name: "destroy",
					Help: "Destroys a room in a particular direction.",
					Arguments: []*CommandArgument{
						{
							Name: "direction",
						},
					},
					Handler: handleRoomDestroyCommand,
				},
			},
		},
		{
			Name: "character",
			Help: "Allows you to manage characters.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Subcommands: []*Command{
				{
					Name: "list",
					Help: "Lists the characters in the game, optionally using a filter.",
					Arguments: []*CommandArgument{
						{
							Name:     "filter",
							Optional: true,
						},
					},
					Handler: handleCharacterListCommand,
				},
				{
					Name: "set",
					Help: "Allows you to set a character attribute. Leave value empty to revert to default.",
					Arguments: []*CommandArgument{
						{
							Name: "character",
						},
						{
							Name: "property",
						},
						{
							Name:             "value",
							IncludeRemaining: true,
							Optional:         true,
						},
					},
					Handler: handleCharacterSetCommand,
				},
				{
					Name: "edit",
					Help: "Opens the editor panel for a character.",
					Arguments: []*CommandArgument{
						{
							Name:     "character",
							Optional: true,
						},
					},
					Handler: handleCharacterEditCommand,
				},
			},
		},
		{
			Name: "save",
			Help: "Writes the in-memory game data to disk.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Handler: handleSaveCommand,
		},
		{
			Name: "reload",
			Help: "Updates, builds and reloads the server, client, or both.",
			Arguments: []*CommandArgument{
				{
					Name: "component",
				},
			},
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Handler: handleReloadCommand,
		},
		{
			Name: "refresh",
			Help: "Asks the server to re-render the data on the client.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Handler: handleRefreshCommand,
		},
		{
			Name:     "whisper",
			AltNames: []string{"w"},
			Help:     "Sends a private message to another online character.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Arguments: []*CommandArgument{
				{
					Name: "target",
				},
				{
					Name:             "message",
					IncludeRemaining: true,
				},
			},
			Handler: handleWhisperCommand,
		},
		{
			Name: "who",
			Help: "Displays the characters currently playing.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Handler: handleWhoCommand,
		},
		{
			Name: "mob",
			Help: "Allows you to manage mobiles (npcs/monsters).",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Subcommands: []*Command{
				{
					Name: "list",
					Help: "Lists the mobs in the game, optionally using a filter.",
					Arguments: []*CommandArgument{
						{
							Name:     "filter",
							Optional: true,
						},
					},
					Handler: handleMobListCommand,
				},
				{
					Name: "create",
					Help: "Creates a new mob in the game.",
					Arguments: []*CommandArgument{
						{
							Name: "name",
						},
					},
					Handler: handleMobCreateCommand,
				},
				{
					Name: "edit",
					Help: "Opens the editor panel for a mob.",
					Arguments: []*CommandArgument{
						{
							Name: "mob",
						},
					},
					Handler: handleMobEditCommand,
				},
				{
					Name: "set",
					Help: "Allows you to set a mob attribute. Leave value empty to revert to default.",
					Arguments: []*CommandArgument{
						{
							Name: "mob",
						},
						{
							Name: "property",
						},
						{
							Name:             "value",
							IncludeRemaining: true,
							Optional:         true,
						},
					},
					Handler: handleMobSetCommand,
				},
				{
					Name: "spawn",
					Help: "Spawns a mob at this location.",
					Arguments: []*CommandArgument{
						{
							Name: "mob",
						},
					},
					Handler: handleMobSpawnCommand,
				},
			},
		},
		{
			Name: "item",
			Help: "Allows you to manage items.",
			Permissions: &CommandPermissions{
				RequireCharacter: true,
			},
			Subcommands: []*Command{
				{
					Name: "list",
					Help: "Lists the items in the game, optionally using a filter.",
					Arguments: []*CommandArgument{
						{
							Name:     "filter",
							Optional: true,
						},
					},
					Handler: handleItemListCommand,
				},
				{
					Name: "create",
					Help: "Creates a new item in the game.",
					Arguments: []*CommandArgument{
						{
							Name: "name",
						},
					},
					Handler: handleItemCreateCommand,
				},
			},
		},
	}

	for _, cmd := range commands {
		Armeria.commandManager.RegisterCommand(cmd)
	}
}
