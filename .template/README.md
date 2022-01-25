# golang-cli-template

This is a template for creating golang cli application. 

It uses github actions to run goreleaser and tests

Edit all files and replace the following variables with the appropriate values

```
--BINARY-- # Would likely be golang-cli-template for this repo, if this was actually a cli
--REPOLINK-- # https://github.com/Jmainguy/golang-cli-template
--DESCRIPTION-- # Github Description, A template for creating golang cli applications
--REPOOWNER-- # Jmainguy
--REPONAME-- # golang-cli-template
--MAINTAINERNAME-- # Jonathan Seth Mainguy
--MAINTAINEREMAIL-- # jon@soh.re
--USAGE-- # Usage of binary, usually multi lines, suggest placing inside .usage text file
```

Add a secret to the new repo called GORELEASER_TOKEN with permissions to homebrew-tap repo.
