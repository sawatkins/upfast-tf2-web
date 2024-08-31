// Configuration
const API_ENDPOINT = '/api/server-ips';
const POLL_INTERVAL = 60000; // 1 minute

// Function to fetch server IPs
async function fetchServerIPs() {
    try {
        const response = await fetch(API_ENDPOINT);
        return await response.json();
    } catch (error) {
        console.error('Error fetching server IPs:', error);
        return [];
    }
}

// Function to fetch server info
async function fetchServerInfo(ip) {
    try {
        const response = await fetch(`/api/server-info?ip=${ip}`);
        return await response.json();
    } catch (error) {
        console.error(`Error fetching server info for ${ip}:`, error);
        return null;
    }
}

// Function to update table row
function updateTableRow(serverInfo) {
    const tableBody = document.querySelector('#server-table table tbody');
    let row = tableBody.querySelector(`tr[data-ip="${serverInfo.public_ip}"]`);
    
    if (!row) {
        row = document.createElement('tr');
        row.setAttribute('data-ip', serverInfo.public_ip);
        tableBody.appendChild(row);
    }

    // console.log("serverInfo", serverInfo);

    // Always set the region to "us-west" for now until more regions are supported
    const region = "us-west";

    row.innerHTML = `
        <td class="region ${region}">
            <img src="${region}.svg" alt="${region} flag" class="flag-icon">
            <span>${region}</span>
        </td>
        <td>Online</td>
        <td>${serverInfo.map}</td>
        <td>${serverInfo.players}/${serverInfo.max_players}</td>
        <td><a href="steam://connect/${serverInfo.public_ip}:27015">Connect</a></td>
    `;
}

// Main polling function
async function pollServers() {
    const ips = await fetchServerIPs();
    
    const headerRow = document.querySelector('#header-row');
    const defaultRow = document.querySelector('#default-row');
    const table = document.querySelector('#server-table table');

    if (ips.length > 0) {
        defaultRow.style.display = 'none';

        ips.forEach(ip => {
            fetchServerInfo(ip).then(serverInfo => {
                updateTableRow(serverInfo);
            });
        });
    } else {
        defaultRow.style.display = 'table-row';
        // Remove all rows except the header and default row
        Array.from(table.rows).forEach(row => {
            if (row !== headerRow && row !== defaultRow) {
                row.remove();
            }
        });
    }

}

// Start polling
setInterval(pollServers, POLL_INTERVAL);
pollServers(); // Initial poll

