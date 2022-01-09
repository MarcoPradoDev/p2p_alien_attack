# ENUNCIADO

Se le pide implementar una versión distribuida descentralizada (P2P) del juego Capturando Alienígenas. Estos alienígenas aparecen en posiciones aleatorias en la tierra con una longitud y latitud; para contrarrestar este peligro se crea una red de seguridad conformada por nodos (cada uno tiene una longitud y latitud fija en la tierra) que representan naves aliadas interconectadas y un centinela es el encargado de enviar a cualquier de estas naves un mensaje con la longitud y latitud de cada alienígena que aparece en la tierra para que luego la nave aliada transmita al resto este mensaje. Con estos datos (longitud y latitud de alienígena) cada navel aliada calcula la distancia que tiene al alienígena de modo que la nave que este mas cerca capturara al alienígena y mostrara un mensaje con la identification de la nave. Debera validar que, si dos o mas aliados logran obtener la misma distancia, solo uno de ellos realizara la captura. Se pide implementar lo siguiente:

- Crear las estructuras alienígena y aliados con atributos según su criterio, pero ambos deben incluir obligatoriamente información de latitud y longitud.
- Toda comunicación sera usando mensajes serializados en formato JSON.
- Cada aliado al conectarse, deberá solicitar su dirección al usuario y la dirección de un aliado cercano que ya sea parte de la red. En caso sea el primer jugador no solicita dicha información.
- Implementar la función que permita a un nuevo aliado solicitar su registro en la red.
- Implementar la función que permita a un aliado que recibe el registro de un nuevo aliado, notificar a los demás de la llegada de dicho miembro.
- Implementar una función que permita a un aliado en la red, aceptar la petición de registro de un nuevo aliado.
- Implementar una función que permita a un aliado ser notificado de la llegada de un nuevo miembro y agregarlo a su lista para comunicaciones posteriores.
- Implementar componente que comunique la aparición de alienígenas en posiciones aleatorias cada cierto tiempo.
- Implementar la función que permita recibir los mensajes de la aparición de los alienígenas.
- Implementar la función para realizar la captura de alienígenas por parte de un aliado.
- Asegurar que todas las funcionalidades implementadas estén libres de condiciones de carrera (principio de competencia) y deadlocks.
