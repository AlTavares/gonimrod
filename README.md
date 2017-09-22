
# Gonimrod

Command line app to send messages using the Nimrod API https://www.nimrod-messenger.io/

## Installation

`go get -u github.com/AlTavares/gonimrod`

## Usage

`gonimrod -k [api_key] [message]`

or set the NIMROD_KEY environment variable and use it like this:

`gonimrod [message]`

you can also pipe the output of a program to gonimrod

`cat file.txt | gonimrod`

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## Credits

https://www.nimrod-messenger.io/
