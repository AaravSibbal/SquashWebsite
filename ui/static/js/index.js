const rankingTableBody = document.getElementById("ranking-table-body")

async function getRanking(){
    fetch("/players/ranking", {
        method: "GET",
    }).then((response)=>{
        if (!response.ok) {
            throw new Error("response was not okay")
        }

        return response.text()
    }).then(htmlContent=>{
        rankingTableBody.innerHTML = htmlContent 
    }).catch(error => {
        console.error("there was a problem with the fetch operating\n\n"+error)
    })
}

document.addEventListener('DOMContentLoaded', ()=>{
    getRanking()
})