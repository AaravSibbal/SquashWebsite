const playerDiv = document.getElementById('player-stat')
const playerGraph = document.getElementById("player-graph")
const playerMatchesBody = document.getElementById('player-matches-body')
const playerName = getPlayerNameFromUrl()
let record = 1

document.addEventListener('DOMContentLoaded', ()=>{
    getPlayer()
    getGraph()
    getMatches()
})

async function getPlayer(){
    fetch(`${playerName}/stat`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }

        return response.text()
    }).then(htmlContent=>{
        playerDiv.innerHTML = htmlContent
    }).catch(error=>{
        console.error(error)
    })
}

async function getGraph(){
    fetch(`/${playerName}/graph`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }

        return response.text()
    }).then(htmlContent=>{
        playerGraph.innerHTML = htmlContent
    }).catch(error=>{
        console.error(error)
    })
}

/**
 * gets player matches in html format
 */
async function getMatches(){

    fetch(`/${playerName}/matches?record=${record}`, {
        method: "GET"
    }).then((response)=>{
        if(!response.ok){
            throw new Error("response was not okay")
        }

        return response.text()
    }).then((htmlContent)=>{
        playerMatchesBody.innerHTML += htmlContent
        record++
    }).catch(error=>{
        console.error("there was a problem with the request")
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