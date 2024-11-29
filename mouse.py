import pyautogui
import time

def move_mouse(duration=10, interval=1):
    """
    Déplace la souris de manière périodique.
    
    :param duration: Durée totale du déplacement en secondes.
    :param interval: Intervalle entre chaque mouvement en secondes.
    """
    end_time = time.time() + duration
    while time.time() < end_time:
        # Obtenir la position actuelle de la souris
        x, y = pyautogui.position()
        
        # Déplacer la souris légèrement
        pyautogui.moveTo(x + 10, y + 10)
        time.sleep(interval)
        
        # Revenir à la position initiale
        pyautogui.moveTo(x, y)
        time.sleep(interval)

if __name__ == "__main__":
    move_mouse(duration=20, interval=2)  # Bouge la souris pendant 20 secondes
