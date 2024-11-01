class Match {
    constructor(id, playerA, playerB, playerARating, playerBRating, playerWon, when){
        // all these are strings exccept when will figure that out later
        this.id = id
        this.playerA = playerA
        this.playerB = playerB
        this.playerARating = playerARating
        this.playerBRating = playerBRating
        this.playerWon = playerWon
        this.when = when
    }

    // getters
    getID() { return this.id }
    getPlayerA() { return this.playerA }
    getPlayerB() { return this.playerB }
    getPlayerARating() { return this.playerARating }
    getPlayerBRating() { return this.playerBRating }
    getPlayerWon() { return this.playerWon }
    getWhen() { return this.when }

    // html
    toHTMLRow(){
        let row = document.createElement("tr")
        
        let playerA = document.createElement("td")
        playerA.textContent = this.getPlayerA()
        
        let playerB = document.createElement("td")
        playerB.textContent = this.getPlayerB()

        let playerARating = document.createElement("td")
        playerARating.textContent = this.getPlayerARating()

        let playerBRating = document.createElement("td")
        playerBRating.textContent = this.getPlayerBRating()

        let playerWon = document.createElement("td")
        playerWon.textContent = this.getPlayerWon()

        let when = document.createElement("td")
        when.textContent = this.getWhen()

        row.append(playerA, playerARating, playerB, playerBRating, playerWon, when)

        return row
    }


}