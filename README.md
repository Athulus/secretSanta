# secretSanta
an app that pairs up people for a secret Santa gift exchange

This program takes in a JSON file with a list of names and email addresses. It will randomly pair each person on the list with another person on the list. Then it will send an email to each person letting them know who they need to get a gift for. It makes sure that no one is paired with themselves and that no ones is chosen as a recepient more than once.

## usage
you need to prepare two files for input 

`list.json` will give the names of all of the participants in the secret santa. it looks like:
```
{
    "name1":"name1@email.com",
    "name2":"name2@email.com"
}
```
you will also need information for the smtp serve you want to connect to in a file `smtp.json`. This should look like:
```
{
    "Uname":"username",
    "Pass":"password",
    "Server":"host.address",
    "Port":"587"
}
```
You can use any smtp server that you  have acess to but if you need one, I was able to send email through gmails smtp interface with this documentation from google https://support.google.com/accounts/answer/185833

once you have these files in the directory of the go file you just need to run `go run secretSanta.go`

