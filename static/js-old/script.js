function changeMessage() {
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;
    // Lista de mensajes alternativos
    const alternateMessages = [
    "¡Hola de nuevo!",
    "¡Bienvenido de vuelta!",
    "¡Qué tengas un buen día!"
    ];
    // Generar un mensaje aleatorio que no sea igual al mensaje actual
    let newMessage;
    do {
    newMessage = alternateMessages[Math.floor(Math.random() *
    alternateMessages.length)];
    } while (newMessage === currentMessage);
    // Actualizar el contenido del elemento de mensaje
    messageElement.textContent = newMessage;
    }