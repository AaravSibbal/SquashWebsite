# Squash Website

## Author
Aarav Sibbal

## Reason

This is a website that is made along side SquashEloRatingSystem which is a discord bot for Carleton University Squash Club. It's purpose is to show players their ranking, access to all matches, a graph showing progress for last 20 games. 

## Choices made

1. I made the choice to have this website view only, as to not require any authentication. This is possible because all my authentication is handled by Discord itself. Data can only be added through our discord server.

2. Having both a website and a Discord Bot: I made this decision because I wanted the users of the bot to like the bot is intuitive that required limiting commands for th bot. Also a website makes it so than people can see their data without feeling like they are adding to the discord chat. The seperation of concerns just make sense. 

3. Graph only including last 20 games: I am using go-echarts to a make charts on the server, I wanted to limit the number of entries that I had to render in the chart because even chess.com and github were having trouble rendering a chart fast enough and I strongly believe people only want to use things that are fast. By this choice people can get a good idea of how they are progressing with the help of visuals and can also have access to it quick.

4. Not using tmpl: From a glance it would appear that I could have used tmpl with this project but I have not moved away from tmpl because I feel like it is hard to debug and a normal html, js and server stack just makes it intuitive to debug. 

## How to run

1. Make sure you have a postgres DB and run the commands from `db.sql` on it. 

2. create an `.env` file and have all connection variables including address and port. 

3. run command ```go run main.go```

## TODO

- Create the initiating event for lazy loading i.e. have the fetch call when scrolled to the end of the table
- Create tests