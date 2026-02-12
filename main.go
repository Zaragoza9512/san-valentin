package main

// ============================================
// PASO 1: IMPORTAR LIBRERÍAS
// ============================================
import (
	"fmt"           // Para imprimir en consola
	"html/template" // Para manejar HTML con variables
	"log"           // Para mostrar errores
	"net/http"      // Para crear el servidor web
	"os"            
)

// ============================================
// PASO 2: CONFIGURACIÓN PERSONALIZABLE
// ============================================
// 👇 CAMBIA ESTOS VALORES POR LOS TUYOS
const (
	nombreDeTuNovia = "Yus 💝"
	tuNombre        = "Luis"
)

// ============================================
// PASO 3: PLANTILLA HTML COMPLETA
// ============================================
// Todo el HTML, CSS y JavaScript va aquí
const pageTemplate = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>💝 Hola mi amoooooor...</title>
    
    <!-- ========== ESTILOS CSS ========== -->
    <style>
        /* Resetear estilos por defecto del navegador */
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        /* Estilo del fondo de toda la página */
        body {
            font-family: 'Arial', sans-serif;
            /* Degradado morado */
            background: linear-gradient(135deg, #ff9a9e 0%, #fad0c4 50%, #a18cd1 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        
        /* Caja blanca central */
        .container {
            background: white;
            border-radius: 30px;
            padding: 50px 40px;
            max-width: 500px;
            text-align: center;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            animation: fadeIn 1s ease-out;
        }
        
        /* Animación de entrada suave */
        @keyframes fadeIn {
            from {
                opacity: 0;
                transform: translateY(-20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        /* Corazones animados */
        .hearts {
            font-size: 50px;
            margin-bottom: 20px;
            animation: pulse 2s infinite;
        }
        
        /* Animación de latido */
        @keyframes pulse {
            0%, 100% { transform: scale(1); }
            50% { transform: scale(1.1); }
        }
        
        /* Título principal */
        h1 {
            color: #764ba2;
            margin-bottom: 15px;
            font-size: 28px;
        }
        
        /* Texto del mensaje */
        .message {
            color: #555;
            font-size: 18px;
            line-height: 1.6;
            margin-bottom: 30px;
        }
        
        /* La pregunta importante */
        .question {
            font-size: 24px;
            font-weight: bold;
            color: #667eea;
            margin: 30px 0;
        }
        
        /* Contenedor de botones */
        .buttons {
            display: flex;
            gap: 15px;
            justify-content: center;
            margin-top: 30px;
        }
        
        /* Estilo base de botones */
        button {
            padding: 15px 40px;
            font-size: 18px;
            border: none;
            border-radius: 50px;
            cursor: pointer;
            transition: all 0.3s;
            font-weight: bold;
        }
        
        /* Botón SÍ - degradado rosa */
        .btn-si {
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
            color: white;
			transition: all 0.5s ease;
        }
        
        /* Efecto hover del botón SÍ */
        .btn-si:hover {
            transform: scale(1.1);
            box-shadow: 0 5px 20px rgba(245, 87, 108, 0.4);
        }
        
        /* Botón NO - gris y más pequeño */
        .btn-no {
            background: #e0e0e0;
            color: #666;
            font-size: 14px;
            padding: 10px 20px;
			transition: all 0.3s ease;
        }
        
        /* El botón NO se hace más pequeño al pasar el mouse */
        .btn-no:hover {
            transform: scale(0.9);
        }
        
        /* Mensaje de éxito (oculto al inicio) */
        .success {
            display: none;
            animation: fadeIn 1s ease-out;
        }
        
        .success h2 {
            color: #f5576c;
            font-size: 32px;
            margin: 20px 0;
        }
        
        /* Emojis de celebración */
        .fireworks {
            font-size: 40px;
            margin: 20px 0;
        }
        
        /* Firma al final */
        .signature {
            margin-top: 30px;
            color: #888;
            font-style: italic;
        }
    </style>
</head>
<body>
    <div class="container">
        <!-- ========== PREGUNTA INICIAL ========== -->
        <div id="pregunta">
            <div class="hearts">💕💖💕</div>
            
            <!-- {{.Nombre}} será reemplazado por Go -->
            <h1>Hola {{.Nombre}}!</h1>
            
            <p class="message">
                No necesito una fecha en el calendario para recordarte cuánto te amo, pero sí quiero aprovecharla para consentirte como te mereces
            </p>
            
            <p class="question">
                ¿Quieres ser mi San Valentín? 💝
            </p>
            
            <div class="buttons">
                <!-- Botón SÍ: llama a responderSi() -->
                <button class="btn-si" id="btnSi" onclick="responderSi()">¡Sí! 💕</button>
                
                <!-- Botón NO: se mueve cuando pasas el mouse -->
                <button class="btn-no" id="btnNo" onclick="cambiarTextoNo()">No</button>
            </div>
            
            <!-- {{.DeName}} será reemplazado por Go -->
            <p class="signature">Con todo mi amor, {{.DeName}} 💌</p>
        </div>
        
        <!-- ========== MENSAJE DE ÉXITO ========== -->
        <!-- Oculto al inicio, se muestra cuando dice SÍ -->
        <div id="exito" class="success">
            <div class="fireworks">🎉✨🎊✨🎉</div>
            <h2>¡Sííííí! 🎉</h2>
            <p class="message">
                ¡Sabía que dirías que sí!💕
				Ya quiero que sea sábado para que disfrutemos de nuestra compañía en el picnic y la ida al cine.
				Eres mi persona favorita en el mundo, ¡te amo! 🎬✨💖
            </p>
            <div class="hearts">❤️💖❤️</div>
        </div>
    </div>
    
    <!-- ========== JAVASCRIPT (LÓGICA) ========== -->
    <script>
		let contadorNo = 0;
    
    	const frasesNo = [
        	"No",
        	"¿Estás segura? 🤔",
        	"Piénsalo bien... 💭",
        	"¿De verdad? 🥺",
        	"Dame una oportunidad 💕",
        	"No seas así... 😢",
        	"¿Por favor? 🙏",
        	"Te va a encantar 💖",
        	"Última oportunidad... ⏰",
        	"Ya solo di que sí 😊"
    	];
        // Función que se ejecuta cuando presionan SÍ
        function responderSi() {
            // Ocultar la pregunta
            document.getElementById('pregunta').style.display = 'none';
            // Mostrar mensaje de éxito
            document.getElementById('exito').style.display = 'block';
            
            // Crear confetti (corazones cayendo)
            for(let i = 0; i < 50; i++) {
                crearConfetti();
            }
        }
         // Nueva función para cambiar texto del botón NO
    	function cambiarTextoNo() {
        	const btnNo = document.getElementById('btnNo');
        	const btnSi = document.getElementById('btnSi');
        
        	contadorNo++;
        
        	// Cambiar el texto del botón NO
        	if (contadorNo < frasesNo.length) {
            	btnNo.textContent = frasesNo[contadorNo];
        	}
        
        	// Agrandar el botón SÍ progresivamente
        	const nuevoTamano = 1 + (contadorNo * 0.15);
        	btnSi.style.transform = 'scale(' + nuevoTamano + ')';
        
        	// Achicar el botón NO progresivamente
        	const tamanoNo = 1 - (contadorNo * 0.08);
        	if (tamanoNo > 0.3) {
            	btnNo.style.transform = 'scale(' + tamanoNo + ')';
        	} else {
            	btnNo.style.transform = 'scale(0.3)';
        	}
        
        	// Cuando llegue al final, hacer el botón NO casi invisible
        	if (contadorNo >= frasesNo.length - 1) {
            	btnNo.style.opacity = '0.3';
            	btnNo.style.cursor = 'not-allowed';
        	}
    	}
        // Función que mueve el botón NO cuando pasan el mouse
        
        // Función que crea un corazón cayendo
        function crearConfetti() {
            const confetti = document.createElement('div');
            // Emojis aleatorios de corazón
            confetti.textContent = ['❤️', '💕', '💖', '💗', '💝'][Math.floor(Math.random() * 5)];
            confetti.style.position = 'fixed';
            confetti.style.left = Math.random() * 100 + '%';
            confetti.style.top = '-50px';
            confetti.style.fontSize = '30px';
            confetti.style.animation = 'caer 3s linear';
            confetti.style.pointerEvents = 'none';
            
            document.body.appendChild(confetti);
            
            // Eliminar después de 3 segundos
            setTimeout(() => confetti.remove(), 3000);
        }
        
        // Crear animación de caída
        const style = document.createElement('style');
        style.textContent = ` + "`" + `
            @keyframes caer {
                to {
                    transform: translateY(100vh) rotate(360deg);
                    opacity: 0;
                }
            }
        ` + "`" + `;
        document.head.appendChild(style);
    </script>
</body>
</html>
`

// ============================================
// PASO 4: ESTRUCTURA DE DATOS
// ============================================
// Esta estructura define qué datos vamos a pasar al HTML
type PageData struct {
	Nombre string // El nombre que aparecerá en {{.Nombre}}
	DeName string // Tu nombre que aparecerá en {{.DeName}}
}

// ============================================
// PASO 5: FUNCIÓN PRINCIPAL
// ============================================
func main() {
	// 1️⃣ PARSEAR EL TEMPLATE
	// Convertir el string HTML en un template usable
	tmpl, err := template.New("valentine").Parse(pageTemplate)
	if err != nil {
		log.Fatal("❌ Error parseando template:", err)
	}

	// 2️⃣ CREAR EL HANDLER (MANEJADOR DE PETICIONES)
	// Esta función se ejecuta cada vez que alguien visita tu sitio
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Crear los datos que vamos a insertar en el HTML
		data := PageData{
			Nombre: nombreDeTuNovia,
			DeName: tuNombre,
		}

		// Indicar que vamos a enviar HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Ejecutar el template: reemplazar {{.Nombre}} y {{.DeName}}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error generando página", http.StatusInternalServerError)
			return
		}
	})

	// 3️⃣ INICIAR EL SERVIDOR
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   💝 SERVIDOR DE SAN VALENTÍN 💝      ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("\n🚀 Servidor corriendo en: http://localhost:%s\n", port)
	fmt.Println("💻 Abre esa URL en tu navegador")
	fmt.Println("🛑 Presiona Ctrl+C para detener el servidor\n")

	// Iniciar el servidor (esto bloquea el programa hasta que lo detengas)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("❌ Error iniciando servidor:", err)
	}
}