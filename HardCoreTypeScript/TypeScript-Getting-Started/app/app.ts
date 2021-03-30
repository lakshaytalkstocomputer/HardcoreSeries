function startGame(){
 // starting a new game
    let messagesElement = document.getElementById("messages");
    messagesElement!.innerText = ""
}

document.getElementById("startGame")!.addEventListener("click", startGame);
