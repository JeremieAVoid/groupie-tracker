document.addEventListener('mousemove', (e) => {
    const yeux = document.querySelectorAll('.œil');

    let moveX = 0;
    let moveY = 0;
    let premier = true;
    yeux.forEach(oeil => {
        if (premier){
            premier = false;
        const rect = oeil.getBoundingClientRect();

        // 1. Centre de l'œil
        const oeilX = rect.left + rect.width / 2;
        const oeilY = rect.top + rect.height / 2;

        // 2. Calcul des vecteurs de distance (dx, dy)
        const dx = e.clientX - oeilX;
        const dy = e.clientY - oeilY;

        // 3. Calcul de l'angle et de la distance réelle de la souris
        const angle = Math.atan2(dy, dx);
        const distanceSouris = Math.hypot(dx, dy);

        // 4. Logique améliorée : 
        // On définit une portée (visionRange) comme dans le code React.
        // On limite le mouvement pour ne pas dépasser 5px (le rayon max de l'œil).
        const visionRange = 150;
        const distancePupille = Math.min(distanceSouris / visionRange * 5, 5);

        // 5. Calcul des coordonnées finales
        moveX = Math.cos(angle) * distancePupille;
        moveY = Math.sin(angle) * distancePupille;
        }
        oeil.style.setProperty('--x', `${moveX}px`);
        oeil.style.setProperty('--y', `${moveY}px`);
    });
});