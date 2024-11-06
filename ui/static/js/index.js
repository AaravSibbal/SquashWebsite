/**
 * @type {Player[]}
 */
let players = []
/**
 * @type {HTMLTableSectionElement}
 */
const rankingTableBody = document.getElementById("ranking-table-body")

class Player {
    constructor(ranking, name, eloRating, wins, losses, totalMatches){
        this.ranking = ranking
        this.name = name
        this.eloRating = eloRating
        this.wins = wins
        this.losses = losses
        this.totalMatches = totalMatches
    }

    // getters
    /**
     * 
     * @returns {string}
     */
    getRanking(){ return this.ranking }
    /**
     * 
     * @returns {string}
     */
    getName(){ return this.name }
    /**
     * 
     * @returns {string}
     */
    getEloRating(){ return this.eloRating }
    /**
     * 
     * @returns {string}
     */
    getWins(){ return this.wins }
    /**
     * 
     * @returns {string}
     */
    getLosses(){ return this.losses }
    /**
     * 
     * @returns {string}
     */
    getTotalMatches(){ return this.totalMatches }
    
    /**
     * creates a html table row element using player attributes
     * @returns {HTMLTableRowElement}
     */
    toHTMLTableRow(){
        let row = document.createElement('tr')

        let ranking = document.createElement('td')
        let name = document.createElement('td')
        let eloRanking = document.createElement('td')
        let wins = document.createElement('td')
        let losses = document.createElement('td')
        let totalMatches = document.createElement('td')
       
        ranking.innerText = this.getRanking()

        let link = document.createElement('a')
        link.href = `/player/${this.getName}`
        link.innerText = this.getName()
        name.appendChild(link)

        eloRanking.innerText = this.getEloRating()
        wins.innerText = this.getWins()
        losses.innerText = this.getLosses()
        totalMatches.innerText = this.totalMatches()

        row.append(ranking, name, eloRanking, wins, losses, totalMatches)

        return row
    }

}

async function getRanking(){
    fetch("/players/ranking", {
        method: "GET",
    }).then((response)=>{
        if (!response.ok) {
            throw new Error("response was not okay")
        }

        return response.json()
    }).then(jsonObj=>{
        console.log(jsonObj)
    }).catch(error => {
        console.error("there was a problem with the fetch operating\n\n"+error)
    })
}

/**
 * takes in the json object and fills html table with the data recieved
 * also adds the players to an array for searching purposes
 * @param {*} obj 
 */
function createPlayerArrFromJson(obj){
    let len = obj.length

    for(i=0; i<len; i++){
        let playerObj = obj[i]
        let player = new Player(playerObj.ranking, playerObj.name, 
            playerObj.eloRating, playerObj.wins, playerObj.losses, 
            playerObj.totalMatches)
        players.push(player)
        let playerHTMLRow = player.toHTMLTableRow()
        rankingTableBody.appendChild(playerHTMLRow)
    }
}

document.addEventListener('DOMContentLoaded', ()=>{
    getRanking()
})