function reloadPage() {
    window.location.reload(); // Reload the page
}

function hideFixtures() {
    document.getElementById('fixtures-container').classList.add('hidden');
}

// Function to handle form submission
const handleFormSubmit = (event) => {
    event.preventDefault(); // Prevent the default form submission

    // Extract form data
    const form = event.target;
    const formData = new FormData(form);

    // Send a POST request to the server with form data
    fetch(form.action, {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(player => {
        // Create a new row for the player
        const playersTableBody = document.querySelector('#players-table tbody');
        const row = document.createElement('tr');
        const roleCell = document.createElement('td');
        roleCell.textContent = player.role;
        const nameCell = document.createElement('td');
        nameCell.textContent = player.name;
        row.appendChild(roleCell);
        row.appendChild(nameCell);

        // Insert the new row at the end of the table
        playersTableBody.appendChild(row);
    })
    .catch(error => console.error('Error adding new player:', error));

    // Clear the form fields after submission
    form.reset();
};

// Attach the handleFormSubmit function to the form submission event
document.getElementById('register-form').addEventListener('submit', handleFormSubmit);

// Retrieve player data from the server on page load
fetch('/players')
    .then(response => response.json())
    .then(players => {
        const playersTableBody = document.querySelector('#players-table tbody');
        players.forEach(player => {
            const row = document.createElement('tr');
            const roleCell = document.createElement('td');
            roleCell.textContent = player.role;
            const nameCell = document.createElement('td');
            nameCell.textContent = player.name;
            row.appendChild(roleCell);
            row.appendChild(nameCell);
            playersTableBody.appendChild(row);
        });
    })
    .catch(error => console.error('Error fetching player data:', error));

// Handle "Shuffle Teams" button click
const shuffledTeamsContainer = document.querySelector('#shuffled-teams-container');
document.getElementById('shuffle-button').addEventListener('click', () => {
    fetch('/shuffle')
        .then(response => response.json())
        .then(teams => {
            const shuffledTeamsTableBody = document.querySelector('#shuffled-teams-table tbody');
            shuffledTeamsTableBody.innerHTML = ''; // Clear previous data

            teams.forEach(team => {
                const row = document.createElement('tr');
                const teamCell = document.createElement('td');
                teamCell.textContent = team.name;
                const strikerCell = document.createElement('td');
                strikerCell.textContent = team.striker;
                const defenderCell = document.createElement('td');
                defenderCell.textContent = team.defender;
                row.appendChild(teamCell);
                row.appendChild(strikerCell);
                row.appendChild(defenderCell);
                shuffledTeamsTableBody.appendChild(row);
            });

            shuffledTeamsContainer.classList.remove('hidden');
        })
        .catch(error => console.error('Error fetching shuffled teams:', error));
});

document.getElementById('show-fixtures-button').addEventListener('click', () => {
fetch('/fixtures')
.then(response => response.json())
.then(fixtures => {
    const fixturesTableBody = document.querySelector('#fixtures-table tbody');
    fixturesTableBody.innerHTML = ''; // Clear previous data

    fixtures.forEach((fixture, index) => {
        const row = document.createElement('tr');
        const weekCell = document.createElement('td');
        weekCell.textContent = `Week ${index + 1}`;
        const gamesCell = document.createElement('td');
        gamesCell.textContent = fixture;
        row.appendChild(weekCell);
        row.appendChild(gamesCell);
        fixturesTableBody.appendChild(row);
    });

    document.getElementById('fixtures-container').classList.remove('hidden');
})
.catch(error => console.error('Error fetching fixtures:', error));
});
