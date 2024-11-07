const searchInput = document.getElementById('player-search')

// Event listener for the search input
searchInput.addEventListener("input", function (e) {
  const searchTerm = e.target.value.toLowerCase();
  const filteredPlayers = players.filter(player =>
    player.name.toLowerCase().includes(searchTerm)
  );
  
  displayResults(filteredPlayers);
});

// Function to display the filtered results
/**
 * 
 * @param {Player[]} players 
 */
function displayResults(players) {
  rankingTableBody.innerHTML = ""; // Clear previous results

  players.forEach(player => {
    row = player.toHTMLTableRow()
    rankingTableBody.appendChild(playerElement);
  });
}
