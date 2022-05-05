dnd-cli
-------------------------------------------------------------------------
A command line tool to assist dungeon masters in managing character data from [dndbeyond.com](dndbeyond.com)

Dndbeyond still does not have a public API so this tool serves as a work around by grabbing the raw JSON of
your party's characters. 

<h3>Usage</h3>

Start the listening for requests by running `dnd-cli start`

Pull character data by passing it character ID(s), eg: `dnd-cli add 69377870 5141398`

Get a full list of all character data by running `dnd-cli get`

<h3>TODO</h3>

Implement a real DB rather than in memory or a file

Add more test coverage

Add optional flags for `get` like `-name "Ezmerelda"` to return data for specific characters

Add subcommands under `get` like `ac` to retrieve just "Armor Class" or other commonly used key:value 
atrributes of a character(health, initiative, ability scores, etc). These commands will return the values
for each character unless using a `-name` flag.

Add authentication

Duplicate check when adding a character

Update a character rather than add a duplicate entry

"Roll" initiative by calculating a random number 1-20 and adding initiative bonus per character
