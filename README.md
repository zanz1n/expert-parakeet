# expert-parakeet

A Discord bot made with the only purpose of annoying someone. It was originally
designed to move people between calls until this person rage, but i will add
more stuff in the future.

# How to host

## 1- Create the bot and add it to the server

See the
[official discord documentation](https://discord.com/developers/docs/getting-started#step-1-creating-an-app)

## 2- Enable privileged intents for the bot

Go to the
[Discord developer portal](https://discord.com/developers/applications)

![Image 1](images/1.png)

Mark these 3 fields and save

![Image 2](images/2.png)

## 3- Get the guild id

![Image 3](images/3.png)

## 4- Download the latest release

![Image 4](images/4.png) ![Image 5](images/5.png)

## 5- Open the CMD on the download folder

![Image 6](images/6.png)

## 6- Run the bot

There are three ways to run the bot

### Using environment variables

Create a .env file in the executable directory containing

```bash
# Replace <something> with the real values
DISCORD_TOKEN="<bot token>"
GUILD_ID="<guild id>"
```

And run the bot using

```bash
# The name of the downloaded executable
./windows-amd64.exe
```

### Using command line arguments

```bash
# Replace <something> with the real values
./windows-amd64.exe --token=<bot_token_here> --guild=<guild_id_here>
```

### Prompting the information to stdin

Run the bot

```bash
# The name of the downloaded executable
./windows-amd64.exe
```

You will be prompted with the required information for the bot run. Just copy
them and paste in the CMD. **Note that when you type the token it will not be
shown by default.**

![Image 7](images/7.png) ![Image 8](images/8.png)

### Ignored user
**Setting this option to some user id makes this user invulnerable to the bot.**

You can set it adding the --ignored-user command argument

```bash
# Replace <something> with the real values
./windows-amd64.exe --ignored-user=<user_id_here>
```

Or you can set it with environment variables too!

```bash
# Replace <something> with the real values
IGNORED_USER="<user_id_here>"
```
