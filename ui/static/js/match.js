class Match {
    
    /**
     * 
     * @param {string} playerA 
     * @param {string} playerB 
     * @param {string} playerARating 
     * @param {string} playerBRating 
     * @param {string} playerWon 
     * @param {string} when 
     */
    constructor(playerA, playerB, playerARating, playerBRating, playerWon, when){
        // all these are strings exccept when will figure that out later
        this.playerA = playerA
        this.playerB = playerB
        this.playerARating = playerARating
        this.playerBRating = playerBRating
        this.playerWon = playerWon
        this.when = when
    }

    // getters

    /**
     * 
     * @returns {string}
     */
    getPlayerA() { return this.playerA }
    /**
     * 
     * @returns {string}
     */
    getPlayerB() { return this.playerB }
    /**
     * 
     * @returns {string} 
     */
    getPlayerARating() { return this.playerARating }
    /**
     * 
     * @returns {string}
     */
    getPlayerBRating() { return this.playerBRating }
    /**
     * 
     * @returns {string}
     */
    getPlayerWon() { return this.playerWon }
    /**
     * @returns {string}
     */
    getWhen() { return this.when }

    // html
    /**
     * returns HTMLTableRowElement which can be added to the table
     * @returns {HTMLTableRowElement}
     */
    toHTMLRow(){
        let row = document.createElement("tr")
        
        let playerA = document.createElement("td")
        let playerALink = document.createElement('a')
        playerALink.href = `/player/${this.playerA}`
        playerALink.innerText = this.getPlayerA()
        playerA.appendChild(playerALink)
        
        let playerB = document.createElement("td")
        let playerBLink = document.createElement('a')
        playerBLink.href = `/player/${this.playerB}` 
        playerBLink.innerText = this.playerB
        playerB.appendChild(playerBLink)


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