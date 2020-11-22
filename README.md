# game of life

an artist's rendition of Conway's [game of life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

<!---
![game demo](https://media.giphy.com/media/iw2GFMAnlCzaR2Hpby/giphy.gif)
--->

![demo](img/demo.gif)


## Technology

[Pixel](https://github.com/faiface/pixel) - A hand-crafted 2D game library in Go


## Usage

The Game of Life is a zero-person game, and critically depends on 
its initial state.

Starting the game, a Setup screen greets you, allowing you to draw any 
initial state to start the game. Left-click on any square to plot it.

Once you have finished setting up, press the SPACE key to start the game.


## Development

clone the repository locally

    go get -u github.com/ronnathaniel/game-of-life
    
run the game

    cd $HOME/go/src/github.com/ronnathaniel/game-of-life
    go run .
    
build the application 

    go build
    
build and save package

    go install
    