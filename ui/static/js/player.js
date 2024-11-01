class Player {
    constructor(id, ranking, name, eloRating, wins, losses, draws, totalMatches){
        this.id = id
        this.ranking = ranking
        this.name = name
        this.eloRating = eloRating
        this.wins = wins
        this.losses = losses
        this.draws = draws
        this.totalMatches = totalMatches
    }

    // getters
    getID(){ return this.id; }
    getRanking() { return this.ranking }
    getName() { return this.name; }
    getEloRating() { return this.eloRating }
    getWins() { return this.wins }
    getLosses() { return this.losses }
    getDraws() { return this.draws }
    getTotalMatches() { return this.totalMatches }

    // html
    toHTMLRow(){
        let row = document.createElement('tr')

        let ranking = document.createElement("td")
        ranking.textContent = this.getRanking()

        let name = document.createElement('td')
        name.textContent = this.getName()

        let eloRating = document.createElement('td')
        eloRating.textContent = this.getEloRating()

        let wins = document.createElement('td')
        wins.textContent = this.getWins()

        let losses = document.createElement('td')
        losses.textContent = this.getLosses()

        let totalMatches = document.createElement('td')
        totalMatches.textContent = this.getTotalMatches()

        row.append(ranking, name, eloRating, wins, losses, totalMatches)

        return row
    }
}