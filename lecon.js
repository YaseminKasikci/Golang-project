// La version majeure 0 est une version spéciale indiquant qu'une bibliothèque peut introduire des modifications importantes.
//  nous pouvons supprimer une fonction de la bibliothèque ou la modifier pour accepter un nouvel argument.
//  v1.23..45 à v2.0.0.
// ajoute des fonctionnalités rétrocompatibles : v1.23.45à v1.24.0
// version :  v1.23.45 à v1.23.46.
//  15(modification leger)


//! handler 
//////////////////////////
// http.Handler - interface with the ServerHTTP method;
// http.HandlerFunc - a function type that accepts same args as ServerHTTP
// method.also implements http.ServerHTTP

// http.Handle("/", http.Handler)
// http.HandleFunc("/", http.HandlerFunc)



// fmt.Println("Starting the server on :3000...")

// http.HandleFunc("/", pathHandler) /
// http.HandleFunc("/contact", pathHandler)
//ou
// var router Router
//ou 
// var router http.HandlerFunc
// router = pathHandler
//ou 
// var router http.HandlerFunc = pathHandler

//ou convertir
// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))





////// HTTP 

// Il existe deux types principaux dans le net/httppackage que nous examinons actuellement :

// http.Handler- une interface avec une méthode ServeHTTP
// http.HandlerFunc- une fonction qui accepte les mêmes arguments qu'une méthode ServeHTTP
// Il existe également deux fonctions dans le httppackage qui se ressemblent et qui sont utilisées pour enregistrer leurs types respectivement similaires pour un serveur Web :

// http.Handle- une fonction qui accepte un motif et un http.Handlercomme arguments.
// http.HandleFunc- une fonction qui accepte un modèle et une fonction qui ressemble à un http.HandlerFunccomme arguments.


// Handleest utilisé pour enregistrer tout ce qui implémente l' http.Handlerinterface.
// HandleFuncest utilisé pour enregistrer des fonctions qui ressemblent à un http.HandlerFunc.

// http.HandleFunc("/", homeHandler)
// http.Handle("/contact", http.HandlerFunc(contactHandler))

// | Cas                                                    | Lequel utiliser ?     | Pourquoi                                           |
// | ------------------------------------------------------ | --------------------- | -------------------------------------------------- |
// | Tu veux utiliser une **fonction simple**               | `HandleFunc`          | Le plus rapide et pratique                         |
// | Tu veux utiliser une **struct personnalisée**          | `Handle`              | Tu peux gérer plus de logique, stocker des données |
// | Tu veux passer une **fonction comme un objet handler** | `HandlerFunc`         | Nécessaire quand tu dois donner un `http.Handler`  |
// | Tu crées un **middleware ou serveur custom**           | `Handler` (interface) | Base pour tout handler                             |



// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, r.URL.Path)
// 	// fmt.Fprintln(w, r.URL.RawPath)

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)

// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "PAGE NOT FOUND AUSSI", http.StatusNotFound)
// 		// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		// http.NotFound(w, r)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "page not FOUND")
// 		//  http.NotFoundHandler().ServeHTTP(w, r)
// 	}

// }