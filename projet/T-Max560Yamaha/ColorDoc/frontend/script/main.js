const urlAPI = "http://localhost:3000";

const btn = document.getElementById('btn');

function getSneakers() {
    fetch(urlAPI + '/sneakers')
        .then(response => {
           return response.json();
        })
        .then(data => {
            console.log(data);
        })
        .catch(error => {
            console.log(error);
        })
}

btn.addEventListener('click', getSneakers);