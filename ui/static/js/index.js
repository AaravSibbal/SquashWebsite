async function getRanking(){
    try {
        const response = await fetch("/players/ranking", {
            method: "GET",
        })

        response.json().then((data)=>{
            console.log(data)
        }).catch((error)=>{
            console.log("there was an error reading the data\n\n", error)
        })
    } catch (error) {
        console.log("there was an error getting the data")
    }
}