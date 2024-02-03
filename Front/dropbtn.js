// Obtenez le bouton qui ouvre le menu déroulant
var dropbtn = document.querySelector('.dropbtn');

// Écoutez le clic sur le bouton
dropbtn.onclick = function() {
  // Trouvez le contenu déroulant en utilisant parentNode et querySelector
  var dropdownContent = this.parentNode.querySelector('.dropdown-content');
  
  // Toggle the "show" class qui contrôle l'affichage du menu
  dropdownContent.classList.toggle("show");
}
