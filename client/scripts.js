
const connString = "ws://127.0.0.1:8080/ws";

document.getElementById("ws-log-msg").innerHTML = `
            <p class="entry">
                <span class="timestamp"><b>${new Date().toLocaleString()}</b></span>
                <span class="message">Welcome to our general chat</span>
            </p>
            </br></br></br>`;

let socket = new WebSocket(connString);
console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");
    socket.send("Hi! My name is Web Client!");
};

socket.onmessage = (msg) => {
    document.getElementById('ws-log-msg').innerHTML += `
            <p class="entry">
                <span class="timestamp"><b>${new Date().toLocaleString()}</b></span>
                <span class="message">${msg.data}</span>
            </p>`;
    console.log("New message received:", msg);
};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
};

function wsSendMessage() {
    let msg = document.getElementById('inputBox').value;
    socket.send(msg);
    document.getElementById('inputBox').value = "";
}