const isLoggedIn = true;

function updateHeader() {
    const authSection = document.getElementById('authSection');
    const userSection = document.getElementById('userSection');
    const username = document.getElementById('username');

    if (isLoggedIn) {
        authSection.style.display = 'none';
        userSection.style.display = 'flex';

        username.textContent = "{{Utilisateurs.Username}}";
    } else {
        authSection.style.display = 'flex';
        userSection.style.display = 'none';
    }
}


updateHeader();