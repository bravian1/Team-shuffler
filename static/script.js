function reloadPage() {
    window.location.reload(); // Reload the page
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
        document.getElementById("shuffle-button").disabled = true;
        var button=document.getElementById("shuffle-button");
            button.style.display="none";

});
function createTableRow(match) {
    const row = document.createElement('tr');
    const homeCell = document.createElement('td');
    homeCell.textContent = match.home;
    const awayCell = document.createElement('td');
    awayCell.textContent = match.away;
    row.appendChild(homeCell);
    row.appendChild(awayCell);
    return row;
}

function createTable(fixtures, week) {
    const table = document.createElement("table");
    table.classList.add("table");
    const weekDateDiv = document.createElement("div");
    weekDateDiv.classList.add("table-week-date");
    weekDateDiv.textContent = "Game Week " + week;
    table.appendChild(weekDateDiv);
    const thr = document.createElement("tr");
    const thHome = document.createElement("th");
    thHome.textContent = "Home";
    const thAway = document.createElement("th");
    thAway.textContent = "Away";
    thr.appendChild(thHome);
    thr.appendChild(thAway);
    table.appendChild(thr);

    fixtures.forEach(fixture => {
        const row = createTableRow(fixture);
        table.appendChild(row);
    });

    return table;
}

document.getElementById('show-fixtures-button').addEventListener('click', () => {
    fetch('/fixtures')
        .then(response => response.json())
        .then(fixtures => {
            const fixtureDiv = document.getElementById('fixtures');
            fixtureDiv.innerHTML = ''; // Clear previous content

            // Group fixtures by week
            const fixturesByWeek = fixtures.reduce((acc, fixture) => {
                if (!acc[fixture.week]) {
                    acc[fixture.week] = [];
                }
                acc[fixture.week].push(fixture);
                return acc;
            }, {});

            // Create tables for each week
            Object.entries(fixturesByWeek).forEach(([week, weekFixtures]) => {
                const table = createTable(weekFixtures, week);
                fixtureDiv.appendChild(table);
            });

            // Add CSS class for spacing
            fixtureDiv.classList.add('fixture-spacing');
        })
        .catch(error => console.error('Error fetching fixtures:', error));
});