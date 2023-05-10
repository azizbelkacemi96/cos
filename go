import codecs

# Ouvrir le fichier binaire en mode lecture
with open('fichier_binaire', 'rb') as f_in:
    # Ouvrir le fichier UTF-8 en mode écriture
    with codecs.open('fichier_utf8', 'w', 'utf-8') as f_out:
        # Lire les données binaires du fichier binaire
        data = f_in.read()
        # Convertir les données binaires en unicode
        unicode_data = data.decode('utf-8')
        # Écrire les données unicode dans le fichier UTF-8
        f_out.write(unicode_data)
