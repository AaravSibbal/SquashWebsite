/**
 * @type {Match[]}
 */
let macthes = []
const playerDiv = document.getElementById('player-stat')
const playerGraph = document.getElementById("player-graph")
const playerMatchesBody = document.getElementById('player-matches-body')
const playerName = getPlayerNameFromUrl()
let record = 0

document.addEventListener('DOMContentLoaded', ()=>{
    getPlayer()
    // getGraph()
    getMatches()
})

async function getPlayer(){
    fetch(`/player/${playerName}/stat`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }

        return response.json()
    }).then(jsonObj=>{
        console.log(jsonObj)
        playerJsonToHTML(jsonObj)
    }).catch(error=>{
        console.error(error)
    })
}

async function getGraph(){
    fetch(`/player/${playerName}/graph`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }

        return response.text()
    }).then(htmlContent=>{
        playerGraph.innerHTML = htmlContent
        console.log(htmlContent)
    }).catch(error=>{
        console.error(error)
    })
}

/**
 * gets player matches in html format
 */
async function getMatches(){

    fetch(`/player/${playerName}/matches?record=${record}`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }
        return response.json()
    }).then((jsonObj)=>{
        record++
        console.log(jsonObj)
        jsonToHTMLMatches(jsonObj)
    }).catch(error=>{
        console.error("there was a problem with the request", error)
        record--
    })


}

/**
 * return the name of the player from the url which is in the format /player/:name
 * @returns {string}
 */
function getPlayerNameFromUrl(){
    /**
     * the path is always going to be /player/:name
     * I just need the name part which is why I just split 
     * and get the name
     */

    const path = window.location.pathname
    let pathArr = path.split("/")

    console.log(pathArr)

    return pathArr[2]
}

function jsonToHTMLMatches(jsonObj){
    let len = jsonObj.length

    for(let i=0; i<len; i++){
        let matchObj = jsonObj[i]
        let match = new Match (matchObj.playerA, matchObj.playerB, 
            matchObj.playerARating, matchObj.playerBRating, 
            matchObj.playerWon, matchObj.when)
        macthes.push(match)
        let matchHTMLRow = match.toHTMLRow()
        playerMatchesBody.appendChild(matchHTMLRow)
    }
}

function playerJsonToHTML(jsonObj){
    const nameElem = document.createElement('p')
    const statLineElem = document.createElement('p')
    nameElem.className = "player-name"
    statLineElem.className = "player-statline"  

    const name = jsonObj.name
    const wins = jsonObj.wins
    const losses = jsonObj.losses
    const totalMatches = jsonObj.totalMatches
    const eloRating = jsonObj.eloRating

    const stat = `Elo Rating: ${eloRating}\tWins: ${wins}\tLosses: ${losses}\tTotal Matches: ${totalMatches}`
    nameElem.innerText = name
    statLineElem.innerText = stat

    playerDiv.append(nameElem, statLineElem)

}