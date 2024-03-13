document.querySelectorAll('.item').forEach(item => {
  item.addEventListener('click', () => {
    // Rétracte tous les autres éléments actifs
    document.querySelectorAll('.item.active').forEach(activeItem => {
      if (activeItem !== item) {
        activeItem.classList.remove('active');
        activeItem.querySelector('.detail-content').style.display = 'none';
      }
    });

    // Active ou désactive l'élément cliqué
    const detailContent = item.querySelector('.detail-content');
    if (detailContent.style.display === 'block') {
      detailContent.style.display = 'none';
      item.classList.remove('active');
    } else {
      detailContent.style.display = 'block';
      item.classList.add('active');
    }
  });
});

    